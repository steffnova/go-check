package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Func returns generator for pure functions. Generated function is defined by
// it's output values, and generator for each output value needs to be provided
// through "outputs" parameter. Error is returned if generator's target is not a
// function, len(outputs) doesn't match number of function output values, or
// generator for any of output values returns an error.
func Func(outputs ...Generator) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (arbitrary.Arbitrary, shrinker.Shrinker, error) {
		if target.Kind() != reflect.Func {
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("can't use Func generator for %s type", target)
		}
		if len(outputs) != target.NumOut() {
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("invalid number of output parameters")
		}
		// generators := make([]Generate, len(outputs))
		randoms := make([]Random, len(outputs))
		for index := range outputs {
			// generator, err := arb(target.Out(index), bias, random)
			// if err != nil {
			// 	return nil, fmt.Errorf("failed to create generator for output[%d]. %s", index, err)
			// }
			// generators[index] = generator
			randoms[index] = r.Split()
		}
		randomInt64 := r.Uint64(constraints.Uint64Default())

		return arbitrary.Arbitrary{
			Value: reflect.MakeFunc(target, func(inputs []reflect.Value) []reflect.Value {
				// In order to create 2 different pure functions that have the
				// same signature but generate different ouput, random value is
				// added to the hashed input parameters. This ensure that each
				// function has differently seeded Random.
				seed := int64(arbitrary.HashToInt64(inputs...)) + int64(randomInt64)

				outputValues := make(arbitrary.Arbitraries, target.NumOut())
				for index, generate := range outputs {
					randoms[index].Seed(seed)
					outputValues[index], _, _ = generate(target.Out(index), bias, randoms[index])
				}

				return outputValues.Values()
			}),
		}, nil, nil
	}
}
