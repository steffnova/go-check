package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Func returns generator for pure functions. arbitrary.Arbitraryd function is defined by
// it's output values, and generator for each output value needs to be provided
// through "outputs" parameter. Error is returned if generator's target is not a
// function, len(outputs) doesn't match number of function output values, or
// generator for any of output values returns an error.
func Func(outputs ...arbitrary.Generator) arbitrary.Generator {
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		if target.Kind() != reflect.Func {
			return arbitrary.Arbitrary{}, NewErrorInvalidTarget(target, "Func")
		}
		if len(outputs) != target.NumOut() {
			return arbitrary.Arbitrary{}, fmt.Errorf("%w. Invalid number of generators (%d) used for generating function outputs, expected %d", ErrorInvalidConfig, len(outputs), target.NumOut())
		}
		// arbitraries := make([]arbitrary.Arbitrary, len(outputs))
		randoms := make([]arbitrary.Random, len(outputs))
		for index := range outputs {
			random := r.Split()
			// arb, err := arb(target.Out(index), bias, random)
			// if err != nil {
			// 	return arbitrary.Arbitrary{}, fmt.Errorf("Can't use generator for generating output at index [%d]. %w", index, err)
			// }
			// arbitraries[index] = arb
			randoms[index] = random
		}
		randomInt64 := r.Uint64(constraints.Uint64Default())
		// return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
		return arbitrary.Arbitrary{
			Value: reflect.MakeFunc(target, func(inputs []reflect.Value) []reflect.Value {
				// In order to create 2 different pure functions that have the
				// same signature but generate different ouput, random value is
				// added to the hashed input parameters. This ensure that each
				// function has differently seeded arbitrary.Random.
				seed := int64(arbitrary.HashToInt64(inputs...)) + int64(randomInt64)
				arbitraries := make(arbitrary.Arbitraries, target.NumOut())
				for index := range arbitraries {
					randoms[index].Seed(seed)
					arb, err := outputs[index](target.Out(index), bias, randoms[index])
					if err != nil {
						panic(err)
					}
					arbitraries[index] = arb
				}

				return arbitraries.Values()
			}),
		}, nil
	}
}
