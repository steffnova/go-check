package generator

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Struct returns generator for struct types. Generators for struct fields can be
// passed through "fields" parameter. If generator for a field is not provided, Any()
// generator is used for that field. Error is returned if generator's target is not
// struct, generator for a field that struct doesn't contain is specified or any of
// the field generators returns an error.
func Struct(fields ...map[string]Generator) Generator {
	fieldGenerators := map[string]Generator{}
	if len(fields) != 0 {
		fieldGenerators = fields[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r Random) (arbitrary.Arbitrary, shrinker.Shrinker, error) {
		if target.Kind() != reflect.Struct {
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("can't use Struct generator for %s type", target)
		}

		for fieldName := range fieldGenerators {
			if _, exists := target.FieldByName(fieldName); !exists {
				return arbitrary.Arbitrary{}, nil, fmt.Errorf("%s doesn't have a field: %s", target.String(), fieldName)
			}
		}

		generators := make([]Generator, target.NumField())
		for index := range generators {
			field := target.Field(index)
			if generator, exists := fieldGenerators[field.Name]; !exists {
				generators[index] = Any()
			} else {
				generators[index] = generator
			}
		}

		arb := arbitrary.Arbitrary{
			Value:    reflect.New(target).Elem(),
			Elements: make(arbitrary.Arbitraries, target.NumField()),
		}

		shrinkers := make([]shrinker.Shrinker, target.NumField())
		baseSpeed := 1
		for index, generator := range generators {
			element, shrinker, err := generator(target.Field(index).Type, bias.Speed(baseSpeed), r)
			if err != nil {
				return arbitrary.Arbitrary{}, nil, fmt.Errorf("failed to create generator for field: %s. %s", target.Field(index).Name, err)
			}
			arb.Elements[index], shrinkers[index] = element, shrinker
			reflect.NewAt(
				arb.Value.Field(index).Type(),
				unsafe.Pointer(arb.Value.Field(index).UnsafeAddr()),
			).Elem().Set(arb.Elements[index].Value)
			baseSpeed *= 2
		}

		return arb, shrinker.Struct(shrinker.Chain(shrinker.CollectionElement(shrinkers...), shrinker.CollectionElements(shrinkers...))), nil
	}
}
