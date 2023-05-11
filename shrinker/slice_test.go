package shrinker

import (
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestSlice(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"InvalidOriginal": func(t *testing.T) {
			shrinker := Slice(arbitrary.Arbitrary{}, constraints.LengthDefault())
			if _, err := shrinker(arbitrary.Arbitrary{}, true); err == nil {
				t.Fatalf("Expected error when original arbitrary is not slice")
			}
		},
		"InvalidElements": func(t *testing.T) {
			slice := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(slice)}
			shrinker := Slice(arb, constraints.LengthDefault())
			if _, err := shrinker(arb, true); err == nil {
				t.Fatalf("Expected error when number of arbitrary elements doesn't match the slice size")
			}
		},
		"Shrink": func(t *testing.T) {
			slice := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
			elements := make([]arbitrary.Arbitrary, len(slice))

			for index, element := range slice {
				elements[index] = arbitrary.Arbitrary{
					Value:    reflect.ValueOf(element),
					Shrinker: Uint64(constraints.Uint64Default()),
				}
			}

			arb := arbitrary.Arbitrary{
				Value:    reflect.ValueOf(slice),
				Elements: elements,
			}
			arb.Shrinker = Slice(arb, constraints.Length{Min: 0, Max: 10})

			property := func(in []uint64) bool {
				for _, element := range in {
					if element == 5 {
						return true
					}
				}
				return false
			}

			propertyFailed := true

			for arb.Shrinker != nil {
				var err error
				arb, err = arb.Shrinker(arb, propertyFailed)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
				propertyFailed = property(arb.Value.Interface().([]uint64))

			}

			result := arb.Value.Interface().([]uint64)
			if len(result) != 1 || result[0] != 5 {
				t.Fatalf("Invalid value: %v", result)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
