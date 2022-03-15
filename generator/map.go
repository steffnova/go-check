package generator

import (
	"fmt"
	"math"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Map returns generator for map types. Generators for map's key and value are
// specified by "key" and "value" parameters, respectively. Range of map size values
// is defined by "limits" parameter. If "limits" parameter is not specified default
// [0, 100] range is used instead. Error is returned if generator's target is not a
// map type, key generator returns an error, value generator returns an error or
// limits.Min > limits.Max
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
		if target.Kind() != reflect.Map {
			return nil, fmt.Errorf("can't use Map generator for %s type", target)
		}
		if constraint.Min > constraint.Max {
			return nil, fmt.Errorf("minimal length value %d can't be greater than max length value %d", constraint.Min, constraint.Max)
		}
		if constraint.Max > uint64(math.MaxInt64) {
			return nil, fmt.Errorf("max length %d can't be greater than %d", constraint.Max, uint64(math.MaxInt64))
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
			size := r.Uint64(constraints.Uint64{
				Min: uint64(constraint.Min),
				Max: uint64(constraint.Max),
			})

			arb := arbitrary.Arbitrary{
				Value:    reflect.MakeMap(target),
				Elements: make(arbitrary.Arbitraries, size),
			}

			shrinkers := make([]shrinker.Shrinker, size)

			filter := arbitrary.FilterPredicate(target, func(in reflect.Value) bool {
				return in.Len() >= int(constraint.Min)
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
				shrinkers[index] = shrinker.Chain(
					shrinker.CollectionElement(keyShrinker, valueShrinker),
					shrinker.CollectionElements(keyShrinker, valueShrinker),
				)

				arb.Value.SetMapIndex(key.Value, value.Value)
			}
			return arb, shrinker.Map(shrinker.CollectionSize(arb.Elements, shrinkers, 0, constraint)).Filter(arb, filter)
		}, nil
	}
}
