package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Struct is Arbitrary that creates struct Generator. Each of the struct's fields has Arbitrary
// assigned implicitly or explictly. Arbitrary for struct fields can be provided explicitly by
// adding it to fieldArbitraries, otherwise implicit Any() Arbitrary is assigned. Error is returned
// if target's reflect.Kind is not Struct, or creation of Generator for any of the fields fails.
func Struct(fieldArbitraries ...map[string]Generator) Generator {
	fieldGenerators := map[string]Generator{}
	if len(fieldArbitraries) != 0 {
		fieldGenerators = fieldArbitraries[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		if target.Kind() != reflect.Struct {
			return nil, fmt.Errorf("target must be a struct")
		}
		generators := make([]Generate, target.NumField())
		for index := range generators {
			field := target.Field(index)
			generator, exists := fieldGenerators[field.Name]
			if !exists {
				generator = Any()
			}
			generate, err := generator(field.Type, bias, r)
			if err != nil {
				return nil, fmt.Errorf("failed to create generator for field: %s. %s", field.Name, err)
			}
			generators[index] = generate
		}

		return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
			arb := arbitrary.Arbitrary{
				Value:    reflect.New(target).Elem(),
				Elements: make(arbitrary.Arbitraries, target.NumField()),
			}

			shrinkers := make([]shrinker.Shrinker, target.NumField())
			for index, generator := range generators {
				arb.Elements[index], shrinkers[index] = generator()
				arb.Value.Field(index).Set(arb.Elements[index].Value)
			}
			return arb, shrinker.Struct(shrinker.Chain(shrinker.CollectionElement(shrinkers...), shrinker.CollectionElements(shrinkers...)))
		}, nil
	}
}
