package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Struct returns generator for struct types. Generators for struct fields can be
// passed through "fields" parameter. If generator for a field is not provided, Any()
// generator is used for that field. Struct generator can only be used for structs that
// has all fields exported. Error is returned if generator's target is not struct,
// struct has unexported fields, or any of the field generators returns an error.
func Struct(fields ...map[string]Generator) Generator {
	fieldGenerators := map[string]Generator{}
	if len(fields) != 0 {
		fieldGenerators = fields[0]
	}

	return func(target reflect.Type, r Random) (Generate, error) {
		if target.Kind() != reflect.Struct {
			return nil, fmt.Errorf("can't use Struct generator for %s type", target)
		}
		generators := make([]Generate, target.NumField())
		for index := range generators {
			field := target.Field(index)
			if field.PkgPath != "" {
				return nil, fmt.Errorf("can't generate struct with unexported fields")
			}
			generator, exists := fieldGenerators[field.Name]
			if !exists {
				generator = Any()
			}
			generate, err := generator(field.Type, r)
			if err != nil {
				return nil, fmt.Errorf("failed to create generator for field: %s. %s", field.Name, err)
			}
			generators[index] = generate
		}

		return func(bias constraints.Bias) (arbitrary.Arbitrary, shrinker.Shrinker) {
			arb := arbitrary.Arbitrary{
				Value:    reflect.New(target).Elem(),
				Elements: make(arbitrary.Arbitraries, target.NumField()),
			}

			shrinkers := make([]shrinker.Shrinker, target.NumField())
			for index, generator := range generators {
				arb.Elements[index], shrinkers[index] = generator(bias)
				arb.Value.Field(index).Set(arb.Elements[index].Value)
			}
			return arb, shrinker.Struct(shrinker.Chain(shrinker.CollectionElement(shrinkers...), shrinker.CollectionElements(shrinkers...)))
		}, nil
	}
}
