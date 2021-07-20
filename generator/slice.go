package generator

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func Slice(arb Arbitrary, constraint constraints.Length) Arbitrary {
	return func(target reflect.Type) (Type, error) {
		if target.Kind() != reflect.Slice {
			return Type{}, fmt.Errorf("targets kind must be Slice. Got: %s", target.Kind())
		}

		generator, err := arb(target.Elem())
		if err != nil {
			return Type{}, fmt.Errorf("failed to create generator for slice elements: %s", err)
		}

		intGenerator, err := Int(constraints.Int(constraint))(reflect.TypeOf(int(0)))
		if err != nil {
			return Type{}, fmt.Errorf("failed to create slice length generator: %s", err)
		}

		return Type{
			Type: reflect.SliceOf(generator.Type),
			Generate: func(rand *rand.Rand) arbitrary.Type {
				elements := make([]arbitrary.Type, intGenerator.Generate(rand).Value().Int())
				for index := range elements {
					arbType := generator.Generate(rand)
					elements[index] = arbType
				}

				return arbitrary.Slice{
					Constraint:  constraint,
					ElementType: generator.Type,
					Elements:    elements,
				}
			},
		}, nil
	}

}
