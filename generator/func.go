package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

func hash(values []reflect.Value) []int64 {
	hashes := make([]int64, len(values))
	for index := range values {
		hashes[index] = 0 // TODO implement hashing function
	}
	return hashes
}

// Func is Arbitrary that creates function Generator. Generator returns pure
// functions (for same input, same output will be returned). outputs parameter
// is variadic parameter that specifies Arbitrary that will be used to create
// Generator for each output parameter. Arbitrary returned by Func will fail to
// create Generator if: target's reflect.Kind is not a function, lenght of outputs
// variadic parameter doesn't match the number of target's function outputs, any
// of the output Arbitraries fails to create it's generator.
func Func(outputs ...Arbitrary) Arbitrary {
	return func(target reflect.Type, r Random) (Generator, error) {
		if target.Kind() != reflect.Func {
			return nil, fmt.Errorf("funcPtr must be a pointer to function")
		}
		if len(outputs) != target.NumOut() {
			return nil, fmt.Errorf("invalid number of output parameters")
		}
		generators := make([]Generator, len(outputs))
		randoms := make([]Random, len(outputs))
		for index, arb := range outputs {
			random := r.Split()
			generator, err := arb(target.Out(index), random)
			if err != nil {
				return nil, fmt.Errorf("failed to create generator for output[%d]. %s", index, err)
			}
			generators[index] = generator
			randoms[index] = random
		}
		return func() arbitrary.Type {
			return arbitrary.Func{
				Fn: reflect.MakeFunc(target, func(inputs []reflect.Value) []reflect.Value {
					seeds := hash(inputs)

					outputs := make([]reflect.Value, target.NumOut())
					for index, generate := range generators {
						randoms[index].Seed(seeds[index])
						outputs[index] = generate().Value()
					}

					return outputs
				}),
			}
		}, nil
	}
}
