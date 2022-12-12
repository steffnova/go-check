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
	return func(target reflect.Type, bias constraints.Bias, r Random) (arbitrary.Arbitrary, shrinker.Shrinker, error) {
		constraint := constraints.LengthDefault()
		if len(limits) != 0 {
			constraint = limits[0]
		}
		if target.Kind() != reflect.Slice {
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("can't use Slice generator for %s type", target)
		}
		if constraint.Min > constraint.Max {
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("minimal length value %d can't be greater than max length value %d", constraint.Min, constraint.Max)
		}
		if constraint.Max > uint64(math.MaxInt64) {
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("max length %d can't be greater than %d", constraint.Max, uint64(math.MaxInt64))
		}

		biasedConstraints := constraints.Uint64(constraint).Baised(bias)
		size := r.Uint64(biasedConstraints)

		shrinkers := make([]shrinker.Shrinker, size)

		arb := arbitrary.Arbitrary{
			Value:    reflect.MakeSlice(target, int(size), int(size)),
			Elements: make([]arbitrary.Arbitrary, int(size)),
		}

		baseSpeed := 1
		for index := range shrinkers {
			elem, shrinker, err := element(target.Elem(), bias.Speed(baseSpeed), r)
			if err != nil {
				return arbitrary.Arbitrary{}, nil, fmt.Errorf("failed to create generator for slice elements: %s", err)
			}
			arb.Elements[index], shrinkers[index], baseSpeed = elem, shrinker, baseSpeed*2
			arb.Value.Index(index).Set(arb.Elements[index].Value)
		}

		return arb, shrinker.Slice(shrinker.CollectionSize(arb.Elements, shrinkers, 0, constraints.Length(biasedConstraints))), nil
	}

}
