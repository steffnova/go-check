package generator

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Map is arbitrary that creates map Generator. key and value parameters
// are Arbitraries used to crete Generators for map's key and value. Map's
// size range is specifed by limits paramter (minimal and maximal value are
// included). Even though limits is a variadic argument only the first value
// is used for defining constraints. Arbitrary will fail to create map Generator
// if target's reflect.Kind is not Map, fails to create map's key and value
// Generator or if creation of map's size fails.
func Map(key, value Arbitrary, limits ...constraints.Length) Arbitrary {
	return func(target reflect.Type) (Generator, error) {
		constraint := constraints.LengthDefault()
		if len(limits) != 0 {
			constraint = limits[0]
		}

		generateKey, keyErr := key(target.Key())
		generateValue, valueErr := value(target.Elem())
		generateSize, sizeErr := Int(constraints.Int(constraint))(reflect.TypeOf(int(0)))

		switch {
		case keyErr != nil:
			return nil, fmt.Errorf("failed to create map's Key generator. %s", keyErr)
		case valueErr != nil:
			return nil, fmt.Errorf("failed to create map's Value generator. %s", valueErr)
		case sizeErr != nil:
			return nil, fmt.Errorf("failed to create map's Size generator. %s", sizeErr)
		}

		return func(rand *rand.Rand) arbitrary.Type {
			size := generateSize(rand).Value().Int()
			pairs := make([]arbitrary.KeyValue, size, size)
			for index := range pairs {
				pairs[index] = arbitrary.KeyValue{
					Key:   generateKey(rand),
					Value: generateValue(rand),
				}
			}

			return arbitrary.Map{
				Constraint: constraint,
				Key:        target.Key(),
				Val:        target.Elem(),
				Pairs:      pairs,
			}
		}, nil
	}
}
