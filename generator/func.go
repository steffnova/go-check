package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Func is Arbitrary that creates function Generator. Generator returns pure
// functions (for same input, same output will be returned). outputs parameter
// is variadic parameter that specifies Arbitrary that will be used to create
// Generator for each output parameter. Arbitrary returned by Func will fail to
// create Generator if: target's reflect.Kind is not a function, lenght of outputs
// variadic parameter doesn't match the number of target's function outputs, any
// of the output Arbitraries fails to create it's generator.
func Func(outputs ...Generator) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		if target.Kind() != reflect.Func {
			return nil, fmt.Errorf("funcPtr must be a pointer to function")
		}
		if len(outputs) != target.NumOut() {
			return nil, fmt.Errorf("invalid number of output parameters")
		}
		generators := make([]Generate, len(outputs))
		randoms := make([]Random, len(outputs))
		for index, arb := range outputs {
			random := r.Split()
			generator, err := arb(target.Out(index), bias, random)
			if err != nil {
				return nil, fmt.Errorf("failed to create generator for output[%d]. %s", index, err)
			}
			generators[index] = generator
			randoms[index] = random
		}
		randomInt64 := r.Int64(constraints.Int64Default())
		return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
			fn := reflect.MakeFunc(target, func(inputs []reflect.Value) []reflect.Value {
				// In order to create 2 different pure functions that have the
				// same signature but generate different ouput, random value is
				// added to the hashed input parameters. This ensure that each
				// function has differently seeded Random.
				seed := int64(arbitrary.HashToInt64(inputs...)) + randomInt64

				outputs := make(arbitrary.Arbitraries, target.NumOut())
				for index, generate := range generators {
					randoms[index].Seed(seed)
					outputs[index], _ = generate()
				}

				return outputs.Values()
			})

			return arbitrary.Arbitrary{
				Value: fn,
			}, nil

		}, nil
	}
}
