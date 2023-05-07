package shrinker

// import (
// 	"fmt"
// 	"reflect"

// 	"github.com/steffnova/go-check/arbitrary"
// )

// type Shrinker func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error)

// func (shrinker Shrinker) Map(mapper interface{}) Shrinker {
// 	mapperVal := reflect.ValueOf(mapper)
// 	switch {
// 	case mapperVal.Kind() != reflect.Func:
// 		return Fail(fmt.Errorf("mapper must be a function"))
// 	case mapperVal.Type().NumIn() != 1:
// 		return Fail(fmt.Errorf("mapper must have 1 input value"))
// 	case mapperVal.Type().NumOut() != 1:
// 		return Fail(fmt.Errorf("mapper must have 1 output value"))
// 	case shrinker == nil:
// 		return nil
// 	default:
// 		return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
// 			if mapperVal.Type().In(0) != arb.Precursors[0].Value.Type() {
// 				return arbitrary.Arbitrary{}, nil, fmt.Errorf("mapper input type must match shrink type")
// 			}

// 			shrink, shrinker, err := shrinker(arb.Precursors[0], propertyFailed)
// 			if err != nil {
// 				return arbitrary.Arbitrary{}, nil, err
// 			}

// 			return arbitrary.Arbitrary{
// 				Value:      mapperVal.Call([]reflect.Value{shrink.Value})[0],
// 				Precursors: []arbitrary.Arbitrary{shrink},
// 			}, shrinker.Map(mapper), nil
// 		}
// 	}
// }

// func (shrinker Shrinker) Filter(predicate interface{}) Shrinker {
// 	val := reflect.ValueOf(predicate)
// 	switch {
// 	case val.Kind() != reflect.Func:
// 		return Fail(fmt.Errorf("predicate must be a function"))
// 	case val.Type().NumIn() != 1:
// 		return Fail(fmt.Errorf("predicate must have one input value"))
// 	case val.Type().NumOut() != 1:
// 		return Fail(fmt.Errorf("predicate must have one output value"))
// 	case val.Type().Out(0).Kind() != reflect.Bool:
// 		return Fail(fmt.Errorf("predicate must have bool as a output value"))
// 	case shrinker == nil:
// 		return nil
// 	default:
// 		return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
// 			shrink, nextShrinker, err := shrinker(arb, propertyFailed)
// 			switch {
// 			case err != nil:
// 				return arbitrary.Arbitrary{}, nil, err
// 			case val.Call([]reflect.Value{shrink.Value})[0].Bool():
// 				return shrink, nextShrinker.Filter(predicate), nil
// 			case nextShrinker == nil:
// 				return shrink, nil, nil
// 			default:
// 				return nextShrinker.Filter(predicate)(shrink, false)
// 			}
// 		}
// 	}

// }

// type binder func(arbitrary.Arbitrary) (arbitrary.Arbitrary, Shrinker, error)

// // Bind returns a shrinker that uses the shrunk value to generate shrink returned by
// // binder. Binder is not guaranteed to be deterministic, as it returns new result value
// // based on root shrinker's shrink and it should be considered non-deterministic. Two
// // shrinkers needs to be passed alongside binder, next and lastFailing. Next shrinker
// // is the shrinker from the previous iteration of shrinking is shrinkering where lastFail
// // that caused last property falsification. Because of "non-deterministic" property of
// // binder, Bind is best paired with Retry combinator that can improve shrinking efficiency.
// func (shrinker Shrinker) Bind(binder binder, last arbitrary.Arbitrary, next, lastFailing Shrinker) Shrinker {
// 	if binder == nil {
// 		return Fail(fmt.Errorf("binder is nil"))
// 	}
// 	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
// 		if propertyFailed {
// 			lastFailing = next
// 			last = arb
// 		}

// 		if shrinker == nil {
// 			after := func(in arbitrary.Arbitrary) arbitrary.Arbitrary {
// 				in.Precursors = append(in.Precursors, last.Precursors[len(last.Precursors)-1])
// 				return in
// 			}
// 			before := func(in arbitrary.Arbitrary) arbitrary.Arbitrary {
// 				in.Precursors = in.Precursors[:len(arb.Precursors)-1]
// 				return in
// 			}
// 			return last, lastFailing.transformBefore(before).transformAfter(after), nil
// 		}

// 		source, sourceShrinker, err := shrinker(arb.Precursors[len(arb.Precursors)-1], propertyFailed)
// 		// time.Sleep(1 * time.Second)
// 		if err != nil {
// 			return arbitrary.Arbitrary{}, nil, err
// 		}

// 		boundValue, boundShrinker, err := binder(source)
// 		if err != nil {
// 			return arbitrary.Arbitrary{}, nil, err
// 		}
// 		return boundValue, sourceShrinker.Bind(binder, last, boundShrinker, lastFailing), nil
// 	}
// }

// type transform func(arbitrary.Arbitrary) arbitrary.Arbitrary

// func (shrinker Shrinker) transformAfter(transformer transform) Shrinker {
// 	if transformer == nil {
// 		return Fail(fmt.Errorf("transformer is nil"))
// 	}
// 	if shrinker == nil {
// 		return nil
// 	}
// 	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
// 		arb, shrinker, err := shrinker(arb, propertyFailed)
// 		if err != nil {
// 			return arbitrary.Arbitrary{}, nil, err
// 		}
// 		return transformer(arb), shrinker.transformAfter(transformer), nil
// 	}
// }

// func (shrinker Shrinker) transformBefore(transformer transform) Shrinker {
// 	if transformer == nil {
// 		return Fail(fmt.Errorf("transformer is nil"))
// 	}
// 	if shrinker == nil {
// 		return nil
// 	}
// 	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
// 		arb, shrinker, err := shrinker(transformer(arb), propertyFailed)
// 		if err != nil {
// 			return arbitrary.Arbitrary{}, nil, err
// 		}
// 		return arb, shrinker.transformBefore(transformer), nil
// 	}
// }

// func (shrinker Shrinker) Or(next Shrinker) Shrinker {
// 	if shrinker == nil {
// 		return next
// 	}
// 	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
// 		if !propertyFailed {
// 			return next(arb, !propertyFailed)
// 		}
// 		return shrinker(arb, propertyFailed)
// 	}
// }

// // Retry returns a shrinker that returns retryValue, and shrinker receiver until either
// // reminingRetries equals 0 or propertyFailed is true. Retry is useful for shrinkers
// // that do not shrink deterministically like shrinkers returned by Bind. On deterministic
// // shrinkers this has no effect and will only increase total time of shrinking process.
// func (shrinker Shrinker) Retry(maxRetries, remainingRetries uint, retryValue arbitrary.Arbitrary) Shrinker {
// 	if shrinker == nil {
// 		return nil
// 	}

// 	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
// 		if propertyFailed || remainingRetries == 0 {
// 			val, next, err := shrinker(arb, propertyFailed)
// 			if err != nil {
// 				return arbitrary.Arbitrary{}, nil, err
// 			}
// 			return val, next.Retry(maxRetries, maxRetries, val), nil

// 		}
// 		return retryValue, shrinker.Retry(maxRetries, remainingRetries-1, retryValue), nil
// 	}
// }

// func (shrinker Shrinker) Validate(validation func(arbitrary.Arbitrary) error) Shrinker {
// 	if validation == nil {
// 		return Fail(fmt.Errorf("validation is nil"))
// 	}
// 	if shrinker == nil {
// 		return nil
// 	}
// 	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
// 		if err := validation(arb); err != nil {
// 			return arbitrary.Arbitrary{}, nil, err
// 		}
// 		arb, shrinker, err := shrinker(arb, propertyFailed)
// 		if err != nil {
// 			return arbitrary.Arbitrary{}, nil, err
// 		}

// 		return arb, shrinker.Validate(validation), nil
// 	}
// }
