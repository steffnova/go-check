package arbitrary

import (
	"fmt"
	"reflect"
)

type Shrinker func(arb Arbitrary, propertyFailed bool) (Arbitrary, error)

func (shrinker Shrinker) Map(mapper interface{}) Shrinker {
	return func(arb Arbitrary, propertyFailed bool) (Arbitrary, error) {
		mapperVal := reflect.ValueOf(mapper)
		switch {
		case mapperVal.Kind() != reflect.Func:
			return Arbitrary{}, fmt.Errorf("mapper must be a function")
		case mapperVal.Type().NumIn() != 1:
			return Arbitrary{}, fmt.Errorf("mapper must have 1 input value")
		case mapperVal.Type().NumOut() != 1:
			return Arbitrary{}, fmt.Errorf("mapper must have 1 output value")
		case shrinker == nil:
			return arb, nil
		}

		if mapperVal.Type().In(0) != arb.Precursors[0].Value.Type() {
			return Arbitrary{}, fmt.Errorf("mapper input type must match shrink type")
		}

		shrink, err := shrinker(arb.Precursors[0], propertyFailed)
		if err != nil {
			return Arbitrary{}, err
		}

		return Arbitrary{
			Value:      mapperVal.Call([]reflect.Value{shrink.Value})[0],
			Precursors: []Arbitrary{shrink},
			Shrinker:   shrink.Shrinker.Map(mapper),
		}, nil
	}
}

func (shrinker Shrinker) Filter(predicate interface{}) Shrinker {
	return func(arb Arbitrary, propertyFailed bool) (Arbitrary, error) {
		val := reflect.ValueOf(predicate)
		switch {
		case val.Kind() != reflect.Func:
			return Arbitrary{}, fmt.Errorf("predicate must be a function")
		case val.Type().NumIn() != 1:
			return Arbitrary{}, fmt.Errorf("predicate must have one input value")
		case val.Type().NumOut() != 1:
			return Arbitrary{}, fmt.Errorf("predicate must have one output value")
		case val.Type().Out(0).Kind() != reflect.Bool:
			return Arbitrary{}, fmt.Errorf("predicate must have bool as a output value")
		case shrinker == nil:
			return arb, nil
		default:
			shrink, err := shrinker(arb, propertyFailed)
			switch {
			case err != nil:
				return Arbitrary{}, err
			case val.Call([]reflect.Value{shrink.Value})[0].Bool():
				shrink.Shrinker = shrink.Shrinker.Filter(predicate)
				return shrink, nil
			case shrink.Shrinker == nil:
				return shrink, nil
			default:
				return shrink.Shrinker.Filter(predicate)(shrink, false)
			}
		}
	}

}

type binder func(Arbitrary) (Arbitrary, error)

// Bind returns a shrinker that uses the shrunk value to generate shrink returned by
// binder. Binder is not guaranteed to be deterministic, as it returns new result value
// based on root shrinker's shrink and it should be considered non-deterministic. Two
// shrinkers needs to be passed alongside binder, next and lastFailing. Next shrinker
// is the shrinker from the previous iteration of shrinking is shrinkering where lastFail
// that caused last property falsification. Because of "non-deterministic" property of
// binder, Bind is best paired with Retry combinator that can improve shrinking efficiency.
func (shrinker Shrinker) Bind(binder binder, last Arbitrary) Shrinker {
	return func(arb Arbitrary, propertyFailed bool) (Arbitrary, error) {
		if binder == nil {
			return Arbitrary{}, fmt.Errorf("binder is nil")
		}
		if propertyFailed {
			last = arb
		}

		if shrinker == nil {
			after := func(in Arbitrary) Arbitrary {
				in.Precursors = append(in.Precursors, last.Precursors[len(last.Precursors)-1])
				return in
			}
			before := func(in Arbitrary) Arbitrary {
				in.Precursors = in.Precursors[:len(arb.Precursors)-1]
				return in
			}
			last.Shrinker = last.Shrinker.TransformBefore(before).TransformAfter(after)
			return last, nil
		}

		sourceShrink, err := shrinker(arb.Precursors[len(arb.Precursors)-1], propertyFailed)
		if err != nil {
			return Arbitrary{}, err
		}

		boundShrink, err := binder(sourceShrink)
		if err != nil {
			return Arbitrary{}, err
		}
		boundShrink.Shrinker = sourceShrink.Shrinker.Bind(binder, last)
		return boundShrink, nil
	}
}

type transform func(Arbitrary) Arbitrary

func (shrinker Shrinker) TransformAfter(transformer transform) Shrinker {
	if shrinker == nil {
		return nil
	}
	return func(arb Arbitrary, propertyFailed bool) (Arbitrary, error) {
		if transformer == nil {
			return Arbitrary{}, fmt.Errorf("transformer is nil")
		}
		shrink, err := shrinker(arb, propertyFailed)
		if err != nil {
			return Arbitrary{}, err
		}
		shrink.Shrinker = arb.Shrinker.TransformAfter(transformer)
		return transformer(shrink), nil
	}
}

func (shrinker Shrinker) TransformBefore(transformer transform) Shrinker {
	if shrinker == nil {
		return nil
	}
	return func(arb Arbitrary, propertyFailed bool) (Arbitrary, error) {
		if transformer == nil {
			return Arbitrary{}, fmt.Errorf("transformer is nil")
		}
		shrink, err := shrinker(transformer(arb), propertyFailed)
		if err != nil {
			return Arbitrary{}, err
		}
		shrink.Shrinker = shrink.Shrinker.TransformBefore(transformer)
		return shrink, nil
	}
}

func (shrinker Shrinker) Or(next Shrinker) Shrinker {
	if shrinker == nil {
		return next
	}
	return func(arb Arbitrary, propertyFailed bool) (Arbitrary, error) {
		if !propertyFailed {
			return next(arb, !propertyFailed)
		}
		return shrinker(arb, propertyFailed)
	}
}

// Retry returns a shrinker that returns retryValue, and shrinker receiver until either
// reminingRetries equals 0 or propertyFailed is true. Retry is useful for shrinkers
// that do not shrink deterministically like shrinkers returned by Bind. On deterministic
// shrinkers this has no effect and will only increase total time of shrinking process.
func (shrinker Shrinker) Retry(maxRetries, remainingRetries uint, retryValue Arbitrary) Shrinker {
	if shrinker == nil {
		return nil
	}

	return func(arb Arbitrary, propertyFailed bool) (Arbitrary, error) {
		if propertyFailed || remainingRetries == 0 {
			shrink, err := shrinker(arb, propertyFailed)
			if err != nil {
				return Arbitrary{}, err
			}
			shrink.Shrinker = shrink.Shrinker.Retry(maxRetries, maxRetries, shrink)
			return shrink, nil

		}
		return Arbitrary{
			Value:      retryValue.Value,
			Precursors: retryValue.Precursors,
			Elements:   retryValue.Elements,
			Shrinker:   shrinker.Retry(maxRetries, remainingRetries-1, retryValue),
		}, nil
	}
}

func (shrinker Shrinker) Validate(validation func(Arbitrary) error) Shrinker {
	if shrinker == nil {
		return nil
	}
	return func(arb Arbitrary, propertyFailed bool) (Arbitrary, error) {
		if validation == nil {
			return Arbitrary{}, fmt.Errorf("validation is nil")
		}
		if err := validation(arb); err != nil {
			return Arbitrary{}, err
		}
		arb, err := shrinker(arb, propertyFailed)
		if err != nil {
			return Arbitrary{}, err
		}

		arb.Shrinker = arb.Shrinker.Validate(validation)
		return arb, nil
	}
}
