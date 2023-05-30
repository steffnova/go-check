package property

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// InputsGenerator generates inputs for [Property] that are passed to the property's predicate (see [Define]).
// Unlike [arbitrary.Generator], which generates a single instance of [arbitrary.Arbitrary], this generator
// produces the same number of arbitraries as the number of passed targets. The bias and random parameters have
// the same role as in [arbitrary.Generator].
type InputsGenerator func(targets []reflect.Type, bias constraints.Bias, random arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error)

// Filter returns a generator that generates values only if the predicate is satisfied. The predicate
// is a function that takes one input (of any type) and returns a bool. The input parameters of the predicate must
// match the input parameters of the property, and the output parameter must be a bool. An error is returned if the
// predicate is invalid or if the generation of any values fails.
//
// NOTE: The returned generator will retry generation until the predicate is satisfied, which
// can affect the speed of the generator.
func (generator InputsGenerator) Filter(predicate any) InputsGenerator {
	return func(targets []reflect.Type, b constraints.Bias, r arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
		predicateType := reflect.TypeOf(predicate)
		if predicateType.NumIn() != len(targets) {
			return nil, nil, inputMissmatchError(targets, predicateType)
		}

		for index := range targets {
			if predicateType.In(index) != targets[index] {
				return nil, nil, inputMissmatchError(targets, predicateType)
			}
		}

		for {
			arbs, shrinker, err := generator(targets, b, r)
			if err != nil {
				return nil, nil, err
			}
			predicateVal := reflect.ValueOf(predicate)
			out := predicateVal.Call(arbitrary.Arbitraries(arbs).Values())
			if out[0].Bool() {
				return arbs, shrinker.Filter(predicate), nil
			}
		}
	}
}

// Log returns a generator that prints generated arbitrary values and their types to standard output.
func (generator InputsGenerator) Log() InputsGenerator {
	return func(t []reflect.Type, b constraints.Bias, r arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
		arbs, shrinker, err := generator(t, b, r)
		if err != nil {
			return nil, nil, err
		}
		for _, val := range arbs.Values() {
			fmt.Printf("<%s> %#v\n", val.Type().String(), val.Interface())
		}

		return arbs, shrinker, nil
	}
}

// NoShrink returns a generator that generates arbitraries without shrinking capabilites (without shrinker).
func (generator InputsGenerator) NoShrink() InputsGenerator {
	return func(t []reflect.Type, b constraints.Bias, r arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
		arbs, shrinker, err := generator(t, b, r)
		if err != nil {
			return nil, nil, err
		}
		for index := range arbs {
			arbs[index].Shrinker = nil
		}
		return arbs, shrinker, nil
	}
}

// Inputs returns [InputsGenerator] used by [Property] for generating property inputs (see [Define]). The variadic
// generators parameter are generators used for each input of property's predicate. Number of generators must match
// the number of targets. Error is returned if:
//   - Number of generators doesn't match number of targets passed to [Inputs Generator]
//   - Any of generators retuns an error while generating arbitrary
func Inputs(generators ...arbitrary.Generator) InputsGenerator {
	return func(targets []reflect.Type, b constraints.Bias, r arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
		if len(targets) != len(generators) {
			return nil, nil, fmt.Errorf("%w: Number of generators (%d) must match number of targets (%d)", ErrorInputs, len(generators), len(targets))
		}
		arbs := make(arbitrary.Arbitraries, len(targets))
		for index := range targets {
			var err error
			arbs[index], err = generators[index](targets[index], b, r)
			if err != nil {
				return nil, nil, fmt.Errorf("generator with index %d failed to generate arbitrary for target: %s. %w", index, targets[index], ErrorInputs)
			}
		}

		arb := arbitrary.Arbitrary{Elements: arbs}
		arb.Shrinker = shrinker.CollectionElements(arb)

		return arbs, shrinkers(arb), nil
	}
}
