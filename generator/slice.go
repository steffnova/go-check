package generator

import (
	"fmt"
	"math"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Slice returns generator for slice types. Slice elements are generated with
// generator specified by "element" parameter. Range of slice size values is
// defined by "limits" parameter. If "limits" parameter is not specified default
// [0, 100] range is used instead. Error is returned if generator's target is not
// a slice type, element generator returns an error, or limits.Min > limits.Max
func Slice(element Generator, limits ...constraints.Length) Generator {
	constraint := constraints.LengthDefault()
	if len(limits) != 0 {
		constraint = limits[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		switch {
		case target.Kind() != reflect.Slice:
			return nil, NewErrorInvalidTarget(target, "Slice")
		case constraint.Min > constraint.Max:
			return nil, fmt.Errorf("%w. Minimal length value %d can't be greater than max length value %d", ErrorInvalidConstraints, constraint.Min, constraint.Max)
		case constraint.Max > uint64(math.MaxInt64):
			return nil, fmt.Errorf("%w. Max length %d can't be greater than %d", ErrorInvalidConstraints, constraint.Max, uint64(math.MaxInt64))
		}

		generator, err := element(target.Elem(), bias, r)
		if err != nil {
			return nil, fmt.Errorf("Failed to use slice element generator. %w", err)
		}

		return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
			biasedConstraints := constraints.Uint64(constraint).Baised(bias)
			size := r.Uint64(biasedConstraints)

			shrinkers := make([]shrinker.Shrinker, size)

			arb := arbitrary.Arbitrary{
				Value:    reflect.MakeSlice(target, int(size), int(size)),
				Elements: make([]arbitrary.Arbitrary, int(size)),
			}

			for index := range shrinkers {
				arb.Elements[index], shrinkers[index] = generator()
				arb.Value.Index(index).Set(arb.Elements[index].Value)
			}

			return arb, shrinker.Slice(shrinker.CollectionSize(arb.Elements, shrinkers, 0, constraints.Length(biasedConstraints)))
		}, nil
	}

}
