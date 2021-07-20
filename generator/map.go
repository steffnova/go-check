package generator

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func Map(key, value Arbitrary, constraint constraints.Int) Arbitrary {
	return func(target reflect.Type) (Type, error) {
		keyGenerator, keyErr := key(target.Key())
		valueGenerator, valueErr := value(target.Elem())
		sizeGenerator, sizeErr := Int(constraint)(reflect.TypeOf(int(0)))
		switch {
		case keyErr != nil:
			return Type{}, fmt.Errorf("failed to create key generator. %s", keyErr)
		case valueErr != nil:
			return Type{}, fmt.Errorf("failed to create value generator. %s", valueErr)
		case sizeErr != nil:
			return Type{}, fmt.Errorf("failed to generate size generator. %s", sizeErr)
		case !keyGenerator.Type.Comparable():
			return Type{}, fmt.Errorf("key generator type [%s] is not comparable", keyGenerator.Type.Name())
		}

		return Type{
			Type: reflect.MapOf(keyGenerator.Type, valueGenerator.Type),
			Generate: func(rand *rand.Rand) arbitrary.Type {
				size := sizeGenerator.Generate(rand).Value().Int()
				pairs := make([]arbitrary.KeyValue, size, size)
				for index := range pairs {
					pairs[index] = arbitrary.KeyValue{
						Key:   keyGenerator.Generate(rand),
						Value: valueGenerator.Generate(rand),
					}
				}

				return arbitrary.Map{
					Constraint: constraint,
					Key:        keyGenerator.Type,
					Val:        valueGenerator.Type,
					Pairs:      pairs,
				}
			},
		}, nil
	}
}
