package check

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

type inputsGenerator func([]reflect.Type, constraints.Bias, arbitrary.Random) (arbitrary.Arbitraries, inputsShrinker, error)

func (i inputsGenerator) Filter(predicate any) inputsGenerator {
	return func(targets []reflect.Type, b constraints.Bias, r arbitrary.Random) (arbitrary.Arbitraries, inputsShrinker, error) {
		for {
			arbs, shrinker, err := i(targets, b, r)
			if err != nil {
				return nil, nil, err
			}
			out := reflect.ValueOf(predicate).Call(arbitrary.Arbitraries(arbs).Values())
			if out[0].Bool() {
				return arbs, shrinker.Filter(predicate), nil
			}
		}
	}
}

func (i inputsGenerator) Log() inputsGenerator {
	return func(t []reflect.Type, b constraints.Bias, r arbitrary.Random) (arbitrary.Arbitraries, inputsShrinker, error) {
		arbs, shrinker, err := i(t, b, r)
		if err != nil {
			return nil, nil, err
		}
		for _, arb := range arbs.Values() {
			fmt.Println(arb.Interface())
		}

		return arbs, shrinker, nil
	}
}

func (i inputsGenerator) NoShrink() inputsGenerator {
	return func(t []reflect.Type, b constraints.Bias, r arbitrary.Random) (arbitrary.Arbitraries, inputsShrinker, error) {
		arbs, shrinker, err := i(t, b, r)
		if err != nil {
			return nil, nil, err
		}
		for index := range arbs {
			arbs[index].Shrinker = nil
		}
		return arbs, shrinker, nil
	}
}

func Inputs(gens ...arbitrary.Generator) inputsGenerator {
	return func(targets []reflect.Type, b constraints.Bias, r arbitrary.Random) (arbitrary.Arbitraries, inputsShrinker, error) {
		arbs := make(arbitrary.Arbitraries, len(targets))
		for index := range targets {
			var err error
			arbs[index], err = gens[index](targets[index], b, r)
			if err != nil {
				return nil, nil, err
			}
		}

		arb := arbitrary.Arbitrary{Elements: arbs}
		arb.Shrinker = shrinker.CollectionElements(arb)

		return arbs, newInputsShrinker(arb), nil
	}
}
