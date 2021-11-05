package shrinker

import (
	"fmt"
	"reflect"
)

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

// Compose creates composition from two shrinkers. Result is a Shrinker that will
// use both shrinkers to provide shrunk values. Both shrinkers are used until they
// are exhausted. Shrinker (receiver) is used first and next (parameter) is used second
func (shrinker Shrinker) Compose(next Shrinker) Shrinker {
	if shrinker == nil {
		return next
	}
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		val, shrinker, err := shrinker(propertyFailed)
		switch {
		case err != nil:
			return reflect.Value{}, nil, err
		case shrinker == nil:
			return val, next, nil
		default:
			return val, shrinker.Compose(next), nil
		}
	}
}

// WithFallback creates a shrinker that will use shrinker (receiver) if property failed
// otherwise next is used instead.
func (shrinker Shrinker) WithFallback(next Shrinker) Shrinker {
	if shrinker == nil {
		return next
	}

	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		switch {
		case propertyFailed:
			return shrinker(true)
		default:
			return next(true)
		}
	}
}
