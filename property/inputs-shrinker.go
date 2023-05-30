package property

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

type inputShrinker func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputShrinker, error)

func (shrinker inputShrinker) Fail(err error) inputShrinker {
	return func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputShrinker, error) {
		return nil, nil, err
	}
}

func (shrinker inputShrinker) Filter(predicate any) inputShrinker {
	val := reflect.ValueOf(predicate)
	switch {
	case val.Kind() != reflect.Func:
		return shrinker.Fail(fmt.Errorf("predicate must be a function"))
	case val.Type().NumOut() != 1:
		return shrinker.Fail(fmt.Errorf("predicate must have one output value"))
	case val.Type().Out(0).Kind() != reflect.Bool:
		return shrinker.Fail(fmt.Errorf("predicate must have bool as a output value"))
	case shrinker == nil:
		return nil
	}
	return func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputShrinker, error) {
		val := reflect.ValueOf(predicate)
		shrinks, shrinker, err := shrinker(arbs, propertyFailed)
		switch {
		case val.Type().NumIn() != len(arbs):
			return nil, nil, fmt.Errorf("predicate must have one input value")
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
		if len(shrink.Elements) != len(arbs) {
			return nil, nil, fmt.Errorf("number of shrinked inputs %d must match number of original inputs %d", len(shrink.Elements), len(arb.Elements))
		}

		return shrink.Elements, shrinkers(shrink), nil
	}
}
