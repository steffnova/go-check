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
	constraint := constraints.LengthDefault()
	if len(limits) != 0 {
		constraint = limits[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		switch {
		case target.Kind() != reflect.Map:
			return nil, NewErrorInvalidTarget(target, "Map")
		case constraint.Min > constraint.Max:
			return nil, fmt.Errorf("%w. Minimal length value %d can't be greater than max length value %d", ErrorInvalidConstraints, constraint.Min, constraint.Max)
		case constraint.Max > uint64(math.MaxInt64):
			return nil, fmt.Errorf("%w. max length %d can't be greater than %d", ErrorInvalidConstraints, constraint.Max, uint64(math.MaxInt64))
		}

		generateKey, err := key(target.Key(), bias, r)
		if err != nil {
			return nil, fmt.Errorf("Failed to use map's Key generator. %w", err)
		}

		generateValue, err := value(target.Elem(), bias, r)
		if err != nil {
			return nil, fmt.Errorf("Failed to use map's Value generator. %w", err)
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

			shrinkers := make([][2]shrinker.Shrinker, size)

			for index := 0; index < int(size); index++ {
				key, keyShrinker := generateKey()
				value, valueShrinker := generateValue()

				for arb.Value.MapIndex(key.Value).IsValid() {
					key, keyShrinker = generateKey()
				}

				arb.Elements[index] = arbitrary.Arbitrary{
					Elements: arbitrary.Arbitraries{key, value},
				}
				shrinkers[index] = [2]shrinker.Shrinker{
					keyShrinker,
					valueShrinker,
				}

				arb.Value.SetMapIndex(key.Value, value.Value)
			}

			return arb, shrinker.Map(arb, shrinkers, constraint)
		}, nil
	}
}
