package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

// Struct is Arbitrary that creates struct Generator. Each of the struct's fields has Arbitrary
// assigned implicitly or explictly. Arbitrary for struct fields can be provided explicitly by
// adding it to fieldArbitraries, otherwise implicit Any() Arbitrary is assigned. Error is returned
// if target's reflect.Kind is not Struct, or creation of Generator for any of the fields fails.
func Struct(fieldArbitraries ...map[string]Arbitrary) Arbitrary {
	fieldGenerators := map[string]Arbitrary{}
	if len(fieldArbitraries) != 0 {
		fieldGenerators = fieldArbitraries[0]
	}

	return func(target reflect.Type, r Random) (Generator, error) {
		if target.Kind() != reflect.Struct {
			return nil, fmt.Errorf("target must be a struct")
		}
		generators := make([]Generator, target.NumField())
		for index := range generators {
			field := target.Field(index)
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

		return func() arbitrary.Type {
			fields := make([]arbitrary.StructField, target.NumField())
			for index := range fields {
				fields[index] = arbitrary.StructField{
					Name: target.Field(index).Name,
					Type: generators[index](),
				}
			}
			return arbitrary.Struct{
				Fields: fields,
			}
		}, nil
	}
}
