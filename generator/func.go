package generator

import (
	"fmt"
	"math/rand"
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

func Func(outputs ...Arbitrary) Arbitrary {
	return func(target reflect.Type) (Type, error) {
		if target.Kind() != reflect.Func {
			return Type{}, fmt.Errorf("funcPtr must be a pointer to function")
		}
		if len(outputs) != target.NumOut() {
			return Type{}, fmt.Errorf("invalid number of output parameters")
		}
		generators := make([]Type, len(outputs))
		for index, arb := range outputs {
			generator, err := arb(target.Out(index))
			if err != nil {
				return Type{}, fmt.Errorf("failed to create generator for output[%d]. %s", index, err)
			}
			if target.Out(index) != generator.Type {
				return Type{}, fmt.Errorf("output type at index [%d] doesn't match it's generator type", index)
			}
			generators[index] = generator
		}
		return Type{
			Type: target,
			Generate: func(_ *rand.Rand) arbitrary.Type {
				return arbitrary.Func{
					Fn: reflect.MakeFunc(target, func(inputs []reflect.Value) []reflect.Value {
						seeds := hash(inputs)

						outputs := make([]reflect.Value, target.NumOut())
						for index, generator := range generators {
							r := rand.New(rand.NewSource(seeds[index]))
							outputs[index] = generator.Generate(r).Value()
						}

						return outputs
					}),
				}
			},
		}, nil
	}
}
