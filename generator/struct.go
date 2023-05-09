package generator

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Struct returns generator for struct types. arbitrary.Generators for struct fields can be
// passed through "fields" parameter. If generator for a field is not provided, Any()
// generator is used for that field. Error is returned if generator's target is not
// struct, generator for a field that struct doesn't contain is specified or any of
// the field generators returns an error.
func Struct(fields ...map[string]arbitrary.Generator) arbitrary.Generator {
	fieldGenerators := map[string]arbitrary.Generator{}
	if len(fields) != 0 {
		fieldGenerators = fields[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		if target.Kind() != reflect.Struct {
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "Struct")
		}

		for fieldName := range fieldGenerators {
			if _, exists := target.FieldByName(fieldName); !exists {
				return arbitrary.Arbitrary{}, fmt.Errorf("%w. %s doesn't have a field: %s", arbitrary.ErrorInvalidConfig, target, fieldName)
			}
		}

		value := reflect.New(target).Elem()
		elements := make(arbitrary.Arbitraries, target.NumField())

		for index := range elements {
			field := target.Field(index)
			generator, exists := fieldGenerators[field.Name]
			if !exists {
				generator = Any()
			}
			arb, err := generator(field.Type, bias, r)
			if err != nil {
				return arbitrary.Arbitrary{}, fmt.Errorf("Failed to use generator for field: %s. %w", field.Name, err)
			}
			elements[index] = arb
		}

		for index := range elements {
			reflect.NewAt(
				value.Field(index).Type(),
				unsafe.Pointer(value.Field(index).UnsafeAddr()),
			).Elem().Set(elements[index].Value)
		}

		arb := arbitrary.Arbitrary{
			Value:    value,
			Elements: elements,
		}
		arb.Shrinker = shrinker.Struct(arb)

		return arb, nil
	}
}
