package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

type Shrinker func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error)

func (shrinker Shrinker) Map(mapper interface{}) Shrinker {
	if shrinker == nil {
		return nil
	}
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
		shrink, shrinker, err := shrinker(arb.Precursors[0], propertyFailed)

		switch mapperVal := reflect.ValueOf(mapper); {
		case mapperVal.Kind() != reflect.Func:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("mapper must be a function")
		case mapperVal.Type().NumIn() != 1:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("mapper must have 1 input value")
		case mapperVal.Type().NumOut() != 1:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("mapper must have 1 output value")
		case mapperVal.Type().In(0) != arb.Precursors[0].Value.Type():
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("mapper input type must match shrink type")
		case err != nil:
			return arbitrary.Arbitrary{}, nil, err
		default:
			return arbitrary.Arbitrary{
				Value:      mapperVal.Call([]reflect.Value{shrink.Value})[0],
				Precursors: []arbitrary.Arbitrary{shrink},
			}, shrinker.Map(mapper), nil
		}
	}
}

func (shrinker Shrinker) Or(next Shrinker) Shrinker {
	if shrinker == nil {
		return next
	}

	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
		if !propertyFailed {
			return next(arb, !propertyFailed)
		}
		return shrinker(arb, propertyFailed)
	}
}

func (shrinker Shrinker) Filter(defaultValue arbitrary.Arbitrary, predicate interface{}) Shrinker {
	if shrinker == nil {
		return nil
	}
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
		shrink, nextShrinker, err := shrinker(arb, propertyFailed)

		switch val := reflect.ValueOf(predicate); {
		case err != nil:
			return arbitrary.Arbitrary{}, nil, err
		case val.Kind() != reflect.Func:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("predicate must be a function")
		case val.Type().NumIn() != 1:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("predicate must have one input value")
		case val.Type().NumOut() != 1:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("predicate must have one output value")
		case val.Type().Out(0).Kind() != reflect.Bool:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("predicate must have bool as a output value")
		case val.Call([]reflect.Value{shrink.Value})[0].Bool():
			return shrink, nextShrinker.Filter(shrink, predicate), nil
		case nextShrinker == nil:
			return defaultValue, nil, nil
		default:
			return nextShrinker.Filter(defaultValue, predicate)(shrink, false)
		}
	}
}

// Retry returns a shrinker that returns retryValue, and shrinker receiver until either
// reminingRetries equals 0 or propertyFailed is true. Retry is useful for shrinkers
// that do not shrink deterministically like shrinkers returned by Bind. On deterministic
// shrinkers this has no effect and will only increase total time of shrinking process.
func (shrinker Shrinker) Retry(maxRetries, remainingRetries uint, retryValue arbitrary.Arbitrary) Shrinker {
	if shrinker == nil {
		return nil
	}

	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
		if propertyFailed || remainingRetries == 0 {
			val, next, err := shrinker(arb, propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, nil, err
			}
			return val, next.Retry(maxRetries, maxRetries, val), nil

		}
		return retryValue, shrinker.Retry(maxRetries, remainingRetries-1, retryValue), nil
	}
}

type binder func(arbitrary.Arbitrary) (arbitrary.Arbitrary, Shrinker, error)

// Bind returns a shrinker that uses the shrunk value to generate shrink returned by
// binder. Binder is not guaranteed to be deterministic, as it returns new result value
// based on root shrinker's shrink and it should be considered non-deterministic. Two
// shrinkers needs to be passed alongside binder, next and lastFailing. Next shrinker
// is the shrinker from the previous iteration of shrinking where lastFailing is shrinker
// that caused last property falsification. Because of "non-deterministic" property of
// binder, Bind is best paired with Retry combinator that can improve shrinking efficiency.
func (shrinker Shrinker) Bind(binder binder, next, lastFailing Shrinker) Shrinker {
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
		if propertyFailed {
			lastFailing = next
		}

		// if shrinker is exhausted, call the lastShrinker that falsified the
		// property with propertyFailed set to true, to continue shrinking process
		if shrinker == nil {
			return lastFailing(arb, true)
		}

		source, sourceShrinker, err := shrinker(arb.Precursors[0], propertyFailed)
		if err != nil {
			return arbitrary.Arbitrary{}, nil, err
		}
		boundValue, boundShrinker, err := binder(source)
		if err != nil {
			return arbitrary.Arbitrary{}, nil, err
		}
		boundValue.Precursors = arbitrary.Arbitraries{source}

		return boundValue, sourceShrinker.Bind(binder, boundShrinker, lastFailing), nil
	}
}
