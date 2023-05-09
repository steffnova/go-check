package generator

import (
	"fmt"
	"math"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Map returns generator for map types. arbitrary.Generators for map's key and value are
// specified by "key" and "value" parameters, respectively. Range of map size values
// is defined by "limits" parameter. If "limits" parameter is not specified default
// [0, 100] range is used instead. Error is returned if generator's target is not a
// map type, key generator returns an error, value generator returns an error or
// limits.Min > limits.Max
//
// Note: arbitrary.Generator will always try to create a map within size limits. This
// means that during key generation it will take into account collision with
// existing map key's. Pool of values from which keys are generated must have
// number of unique values equal to map's maximum length value defined by
// limits parameter, otherwise map generation will be stuck in endless loop.
func Map(keyGenerator, ValueGenerator arbitrary.Generator, limits ...constraints.Length) arbitrary.Generator {
	constraint := constraints.LengthDefault()
	if len(limits) != 0 {
		constraint = limits[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		switch {
		case target.Kind() != reflect.Map:
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "Map")
		case constraint.Min > constraint.Max:
			return arbitrary.Arbitrary{}, fmt.Errorf("%w. Minimal length value %d can't be greater than max length value %d", arbitrary.ErrorInvalidConstraints, constraint.Min, constraint.Max)
		case constraint.Max > uint64(math.MaxInt64):
			return arbitrary.Arbitrary{}, fmt.Errorf("%w. max length %d can't be greater than %d", arbitrary.ErrorInvalidConstraints, constraint.Max, uint64(math.MaxInt64))
		}

		size := r.Uint64(constraints.Uint64{
			Min: uint64(constraint.Min),
			Max: uint64(constraint.Max),
		})

		value := reflect.MakeMap(target)
		elements := make(arbitrary.Arbitraries, size)

		for index := 0; index < int(size); index++ {
			var keyArb arbitrary.Arbitrary
			var err error
			for {
				keyArb, err = keyGenerator(target.Key(), bias, r)
				if err != nil {
					return arbitrary.Arbitrary{}, fmt.Errorf("Failed to use map's Key generator. %w", err)
				}
				if !value.MapIndex(keyArb.Value).IsValid() {
					break
				}
			}

			valueArb, err := ValueGenerator(target.Elem(), bias, r)
			if err != nil {
				return arbitrary.Arbitrary{}, fmt.Errorf("Failed to use map's Value generator. %w", err)
			}

			elements[index] = arbitrary.Arbitrary{
				Elements: arbitrary.Arbitraries{keyArb, valueArb},
				Shrinker: shrinker.CollectionElements(elements[index]),
			}

			value.SetMapIndex(keyArb.Value, valueArb.Value)
		}

		arb := arbitrary.Arbitrary{
			Value:    value,
			Elements: elements,
		}

		arb.Shrinker = shrinker.Map(arb, constraint)

		return arb, nil
	}
}
