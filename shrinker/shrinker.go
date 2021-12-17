package shrinker

import (
	"fmt"
	"reflect"
)

type Shrink struct {
	Value    reflect.Value
	Shrinker Shrinker
}

// Shrinker returns shrinked value and next shrinker. Returned values
// are usually affected by propertyFailed parameter. It helps shrinker
// to decide how to properly shrink the value. If returned Shrinker
// is nil, it indicates that value can no longer be shrinked
type Shrinker func(propertyFailed bool) (reflect.Value, Shrinker, error)

// Map maps the shrinked value to a new one using the mapper.
// Mapper must be a function with one input and one output parameter where
// input parameter must match shrinked type, otherwise panic occurs.
func (shrinker Shrinker) Map(mapper interface{}) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		shrink, shrinker, err := shrinker(propertyFailed)

		switch mapperVal := reflect.ValueOf(mapper); {
		case mapperVal.Kind() != reflect.Func:
			return reflect.Value{}, nil, fmt.Errorf("mapper must be a function")
		case mapperVal.Type().NumIn() != 1:
			return reflect.Value{}, nil, fmt.Errorf("mapper must have 1 input value")
		case mapperVal.Type().NumOut() != 1:
			return reflect.Value{}, nil, fmt.Errorf("mapper must have 1 output value")
		case err != nil:
			return reflect.Value{}, nil, err
		case mapperVal.Type().In(0) != shrink.Type():
			return reflect.Value{}, nil, fmt.Errorf("mapper input type must match shrink type")
		case shrinker != nil:
			shrinker = shrinker.Map(mapper)
			fallthrough
		default:
			shrink = mapperVal.Call([]reflect.Value{shrink})[0]
			return shrink, shrinker, nil
		}
	}
}

// Convert converts shrinked value to a type specified by target parameter.
// Error is returned if shrinking fails or shrinked value can't be converted
// to a target type.
func (shrinker Shrinker) Convert(target reflect.Type) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		switch shrink, shrinker, err := shrinker(propertyFailed); {
		case err != nil:
			return reflect.Value{}, nil, err
		case !shrink.Type().ConvertibleTo(target):
			return reflect.Value{}, nil, fmt.Errorf("shrink of type: %s can't be converted to target: %s", shrink.Type(), target)
		case shrinker == nil:
			return shrink.Convert(target), nil, nil
		default:
			return shrink.Convert(target), shrinker.Convert(target), nil
		}
	}
}

// Filter filters shrink values returned by shrinker using a predicate. Predicate must be
// a function that receives one value and returns boolean. Predicate's input must match
// value that is being shrank. if predicate returns true shrank value is returned from
// shrinker. In case of failure a shrinker (receiver) is being called as if the property
// is failing until either predicate returns true or shrinker is exhausted. Default value
// is returned in case shrinker is exhausted and predicate wasn't satisfied.
func (shrinker Shrinker) Filter(defaultValue reflect.Value, predicate interface{}) Shrinker {
	if shrinker == nil {
		return nil
	}
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		shrink, nextShrinker, err := shrinker(propertyFailed)

		switch val := reflect.ValueOf(predicate); {
		case err != nil:
			return reflect.Value{}, nil, err
		case val.Kind() != reflect.Func:
			return reflect.Value{}, nil, fmt.Errorf("predicate must be a function")
		case val.Type().NumIn() != 1:
			return reflect.Value{}, nil, fmt.Errorf("predicate must have one input value")
		case val.Type().NumOut() != 1:
			return reflect.Value{}, nil, fmt.Errorf("predicate must have one output value")
		case val.Type().Out(0).Kind() != reflect.Bool:
			return reflect.Value{}, nil, fmt.Errorf("predicate must have bool as a output value")
		case val.Call([]reflect.Value{shrink})[0].Bool():
			return shrink, nextShrinker.Filter(shrink, predicate), nil
		case nextShrinker == nil:
			return defaultValue, nil, nil
		default:
			return nextShrinker.Filter(defaultValue, predicate)(false)
		}
	}
}

// Or returns a shrinker that uses either shrinker(receiver) or next. Decision is made
// by evaluating the value of propertyFailed parameter. If property failed, shrinker
// (receiver) will be used, otherwise next is used instead.
func (shrinker Shrinker) Or(next Shrinker) Shrinker {
	if shrinker == nil {
		return next
	}

	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		if !propertyFailed {
			return next(!propertyFailed)
		}
		return shrinker(propertyFailed)
	}
}

// Retry returns a shrinker that returns retryValue, and shrinker receiver until either
// reminingRetries equals 0 or propertyFailed is true. Retry is useful for shrinkers
// that do not shrink deterministically like shrinkers returned by Bind. On deterministic
// shrinkers this has no effect and will only increase total time of shrinking process.
func (shrinker Shrinker) Retry(maxRetries, remainingRetries uint, retryValue reflect.Value) Shrinker {
	if shrinker == nil {
		return nil
	}

	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		if propertyFailed || remainingRetries == 0 {
			val, next, err := shrinker(propertyFailed)
			if err != nil {
				return reflect.Value{}, nil, err
			}
			return val, next.Retry(maxRetries, maxRetries, val), nil

		}
		return retryValue, shrinker.Retry(maxRetries, remainingRetries-1, retryValue), nil
	}
}

type binder func(reflect.Value) (reflect.Value, Shrinker, error)

// Bind returns a shrinker that uses the shrunk value to generate shrink returned by
// binder. Binder is not guaranteed to be deterministic, as it returns new result value
// based on root shrinker's shrink and it should be considered non-deterministic. Two
// shrinkers needs to be passed alongside binder, next and lastFailing. Next shrinker
// is the shrinker from the previous iteration of shrinking where lastFailing is shrinker
// that caused last property falsification. Because of "non-deterministic" property of
// binder, Bind is best paired with Retry combinator that can improve shrinking efficiency.
func (shrinker Shrinker) Bind(binder binder, next, lastFailing Shrinker) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		if propertyFailed {
			lastFailing = next
		}

		// if shrinker is exhausted, call the lastShrinker that falsified the
		// property with propertyFailed set to true, to continue shrinking process
		if shrinker == nil {
			return lastFailing(true)
		}

		source, sourceShrinker, err := shrinker(propertyFailed)
		if err != nil {
			return reflect.Value{}, nil, err
		}
		boundValue, boundShrinker, err := binder(source)
		if err != nil {
			return reflect.Value{}, nil, err
		}

		return boundValue, sourceShrinker.Bind(binder, boundShrinker, lastFailing), nil
	}
}
