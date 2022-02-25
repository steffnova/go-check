package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
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
//
// Note: Generator will always try to create a map within size limits. This
// means that during key generation it will take into account collision with
// existing map key's. Pool of values from which keys are generated must have
// number of unique values equal to map's maximum length value defined by
// limits parameter, otherwise map generation will be stuck in endless loop.
func Map(key, value Generator, limits ...constraints.Length) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		constraint := constraints.LengthDefault()
		if len(limits) != 0 {
			constraint = limits[0]
		}

		generateKey, err := key(target.Key(), bias, r)
		if err != nil {
			return nil, fmt.Errorf("failed to create map's Key generator. %s", err)
		}

		generateValue, err := value(target.Elem(), bias, r)
		if err != nil {
			return nil, fmt.Errorf("failed to create map's Value generator. %s", err)
		}

		return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
			size := r.Int64(constraints.Int64{
				Min: int64(constraint.Min),
				Max: int64(constraint.Max),
			})

			arb := arbitrary.Arbitrary{
				Value:    reflect.MakeMap(target),
				Elements: make(arbitrary.Arbitraries, size),
			}

			shrinkers := make([]shrinker.Shrinker, int(size))

			filter := arbitrary.FilterPredicate(target, func(in reflect.Value) bool {
				return in.Len() >= constraint.Min
			})

			for index := 0; index < int(size); index++ {
				key, keyShrinker := generateKey()
				value, valueShrinker := generateValue()

				for arb.Value.MapIndex(key.Value).IsValid() {
					key, keyShrinker = generateKey()
				}

				arb.Elements[index] = arbitrary.Arbitrary{
					Elements: arbitrary.Arbitraries{key, value},
				}
				shrinkers[index] = shrinker.CollectionElement(keyShrinker, valueShrinker) // shrinker.Chain(shrinker.OneByOne(keyShrinker, valueShrinker), shrinker.All(keyShrinker, valueShrinker))

				arb.Value.SetMapIndex(key.Value, value.Value)
			}
			return arb, shrinker.Map(shrinker.CollectionSize(arb.Elements, shrinkers, 0, constraint)).Filter(arb, filter)
		}, nil
	}
}
