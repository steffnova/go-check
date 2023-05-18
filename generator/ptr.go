package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Ptr returns an [arbitrary.Generator] that creates pointer types. The element parameter defines the
// type to which the pointer points. The limits parameter, even though it is variadic, evaluates only
// the first instance of [constraints.Ptr]. If limits are omitted, [constraints.PtrDefault] is used
// instead. [constraints.Ptr.NilFrequency] influences the occurrence of nil pointers:
//   - NilFrequency of 0 ensures no nil pointers will be generated
//   - NilFrequency > 0 uses formula 1/NilFrequency.
//
// An error is returned if the generator's target parameter is not a pointer or if the element generator
// returns an error.
func Ptr(element arbitrary.Generator, limits ...constraints.Ptr) arbitrary.Generator {
	limit := constraints.PtrDefault()
	if len(limits) != 0 {
		limit = limits[0]
	}
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		if target.Kind() != reflect.Ptr {
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "Ptr")
		}

		isNil := limit.NilFrequency != 0 && r.Uint64(constraints.Uint64{Min: 1, Max: limit.NilFrequency})%limit.NilFrequency == 0

		if isNil {
			return arbitrary.Arbitrary{
				Value: reflect.Zero(target),
			}, nil
		}

		element, err := element(target.Elem(), bias, r)
		if err != nil {
			return arbitrary.Arbitrary{}, fmt.Errorf("failed to generated value pointer: %s points to. %w", target, err)
		}

		value := reflect.New(target.Elem())
		value.Elem().Set(element.Value)

		arb := arbitrary.Arbitrary{
			Value:    value,
			Elements: arbitrary.Arbitraries{element},
		}
		arb.Shrinker = shrinker.Ptr(arb, limit)

		return arb, nil
	}

}
