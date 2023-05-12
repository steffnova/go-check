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
func Slice(elementGenerator arbitrary.Generator, limits ...constraints.Length) arbitrary.Generator {
	constraint := constraints.LengthDefault()
	if len(limits) != 0 {
		constraint = limits[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		switch {
		case target.Kind() != reflect.Slice:
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "Slice")
		case constraint.Min > constraint.Max:
			return arbitrary.Arbitrary{}, fmt.Errorf("%w. Minimal length value %d can't be greater than max length value %d", arbitrary.ErrorInvalidConstraints, constraint.Min, constraint.Max)
		case constraint.Max > uint64(math.MaxInt64):
			return arbitrary.Arbitrary{}, fmt.Errorf("%w. Max length %d can't be greater than %d", arbitrary.ErrorInvalidConstraints, constraint.Max, uint64(math.MaxInt64))
		}

		biasedConstraints := constraints.Uint64(constraint)
		size := r.Uint64(biasedConstraints)

		value := reflect.MakeSlice(target, int(size), int(size))
		elements := make([]arbitrary.Arbitrary, int(size))

		for index := range elements {
			var err error
			elements[index], err = elementGenerator(target.Elem(), bias, r)
			if err != nil {
				return arbitrary.Arbitrary{}, fmt.Errorf("failed to use slice element generator. %w", err)
			}
			value.Index(index).Set(elements[index].Value)
		}

		arb := arbitrary.Arbitrary{
			Value:    value,
			Elements: elements,
		}

		arb.Shrinker = shrinker.Slice(arb, constraints.Length(biasedConstraints))

		return arb, nil
	}

}
