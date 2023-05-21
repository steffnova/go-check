package check

import (
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

type inputsShrinker func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputsShrinker, error)

func (is inputsShrinker) Filter(predicate any) inputsShrinker {
	if is == nil {
		return nil
	}
	return func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputsShrinker, error) {
		val := reflect.ValueOf(predicate)
		shrinks, shrinker, err := is(arbs, propertyFailed)
		switch {
		case err != nil:
			return nil, nil, err
		case val.Call(shrinks.Values())[0].Bool():
			return shrinks, shrinker.Filter(predicate), nil
		case shrinker == nil:
			return shrinks, nil, nil
		default:
			return shrinker.Filter(predicate)(shrinks, false)
		}
	}
}

func newInputsShrinker(arb arbitrary.Arbitrary) inputsShrinker {
	if arb.Shrinker == nil {
		return nil
	}
	return func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputsShrinker, error) {
		arb.Elements = arbs
		shrink, err := arb.Shrinker(arb, propertyFailed)
		if err != nil {
			return nil, nil, err
		}

		return shrink.Elements, newInputsShrinker(shrink), nil
	}
}
