package property

import (
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

type inputShrinker func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputShrinker, error)

func (shrinker inputShrinker) Filter(predicate any) inputShrinker {
	if shrinker == nil {
		return nil
	}
	return func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputShrinker, error) {
		val := reflect.ValueOf(predicate)
		shrinks, shrinker, err := shrinker(arbs, propertyFailed)
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

func shrinkers(arb arbitrary.Arbitrary) inputShrinker {
	if arb.Shrinker == nil {
		return nil
	}
	return func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputShrinker, error) {
		arb.Elements = arbs
		shrink, err := arb.Shrinker(arb, propertyFailed)
		if err != nil {
			return nil, nil, err
		}

		return shrink.Elements, shrinkers(shrink), nil
	}
}
