package generator

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

// Ptr returns Arbitrary generator that can be used to create Ptr<Type> generator. Type to wich pointer
// points is defined by generator retrieved from arb parameter.
func Ptr(arb Arbitrary) Arbitrary {
	return func(target reflect.Type) (Type, error) {
		if target.Kind() != reflect.Ptr {
			return Type{}, fmt.Errorf("target's kind must be Ptr. Got: %s", target.Kind())
		}
		generator, err := arb(target.Elem())
		if err != nil {
			return Type{}, fmt.Errorf("failed to create base generator. %s", err)
		}
		boolGenerator, err := Bool()(reflect.TypeOf(false))
		if err != nil {
			return Type{}, fmt.Errorf("failed to generate bool generator. %s", err)
		}

		return Type{
			Type: reflect.PtrTo(generator.Type),
			Generate: func(rand *rand.Rand) arbitrary.Type {
				return arbitrary.Ptr{
					IsNull: boolGenerator.Generate(rand).Value().Bool(),
					Type:   generator.Generate(rand),
				}
			},
		}, nil
	}
}
