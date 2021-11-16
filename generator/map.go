package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Map is arbitrary that creates map Generator. key and value parameters
// are Arbitraries used to crete Generators for map's key and value. Map's
// size range is specifed by limits paramter (minimal and maximal value are
// included). Even though limits is a variadic argument only the first value
// is used for defining constraints. Arbitrary will fail to create map Generator
// if target's reflect.Kind is not Map, fails to create map's key and value
// Generator or if creation of map's size fails.
func Map(key, value Arbitrary, limits ...constraints.Length) Arbitrary {
	return func(target reflect.Type, r Random) (Generator, error) {
		constraint := constraints.LengthDefault()
		if len(limits) != 0 {
			constraint = limits[0]
		}

		generateKey, err := key(target.Key(), r)
		if err != nil {
			return nil, fmt.Errorf("failed to create map's Key generator. %s", err)
		}

		generateValue, err := value(target.Elem(), r)
		if err != nil {
			return nil, fmt.Errorf("failed to create map's Value generator. %s", err)
		}

		return func() (reflect.Value, shrinker.Shrinker) {
			size := r.Int64(int64(constraint.Min), int64(constraint.Max))

			mapElements := []shrinker.MapShrink{}
			val := reflect.MakeMap(target)
			for index := 0; index < int(size); index++ {
				key, keyShrinker := generateKey()
				value, valueShrinker := generateValue()

				if val.MapIndex(key).IsValid() {
					continue
				}

				mapElements = append(mapElements, shrinker.MapShrink{
					Key: shrinker.Shrink{
						Value:    key,
						Shrinker: keyShrinker,
					},
					Value: shrinker.Shrink{
						Value:    value,
						Shrinker: valueShrinker,
					},
				})
				val.SetMapIndex(key, value)
			}

			return val, shrinker.Map(val, mapElements, constraint)
		}, nil
	}
}
