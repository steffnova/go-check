package generator

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

type StructField struct {
	Name      string
	Generator Arbitrary
}

func Struct(fieldGenerators ...StructField) Arbitrary {
	return func(target reflect.Type) (Type, error) {
		generatorMap := make(map[string]Type, len(fieldGenerators))
		structFields := make([]reflect.StructField, len(fieldGenerators))
		for index, field := range fieldGenerators {
			generator, err := field.Generator(nil)
			if err != nil {
				return Type{}, fmt.Errorf("failed to create generator for field [%s]. %s", field.Name, err)
			}
			if len(field.Name) == 0 {
				return Type{}, fmt.Errorf("")
			}
			if _, exists := generatorMap[field.Name]; exists {
				return Type{}, fmt.Errorf("field with name [%s] defined mutiple times", field.Name)
			}
			generatorMap[field.Name] = generator
			structFields[index] = reflect.StructField{
				Name: field.Name,
				Type: generator.Type,
			}
		}

		return Type{
			Type: reflect.StructOf(structFields),
			Generate: func(rand *rand.Rand) arbitrary.Type {
				fields := make([]arbitrary.StructField, len(fieldGenerators))
				for index, structField := range fieldGenerators {
					fields[index] = arbitrary.StructField{
						Name: structField.Name,
						Type: generatorMap[structField.Name].Generate(rand),
					}
				}

				return arbitrary.Struct{Fields: fields}
			},
		}, nil
	}
}
