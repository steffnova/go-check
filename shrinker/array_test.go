package shrinker

import (
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestArray(t *testing.T) {

	testCases := map[string]func(t *testing.T){
		"InvalidOriginal": func(t *testing.T) {
			arr := [10]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(arr)}
			shrinker := Array(arb)
			if _, err := shrinker(arbitrary.Arbitrary{}, true); err == nil {
				t.Fatalf("Expected error when original arbitrary is invalid")
			}
		},
		"InvalidShrinkers": func(t *testing.T) {
			arr := [10]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
			elements := make([]arbitrary.Arbitrary, len(arr))
			for index, arr := range arr {
				elements[index] = arbitrary.Arbitrary{Value: reflect.ValueOf(arr)}
			}
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(arr), Elements: elements}
			shrinker := Array(arb)
			if _, err := shrinker(arb, true); err == nil {
				t.Fatalf("Expected error when original arbitrary is invalid")
			}
		},
		"ErrorWhenArbitraryIsNotArray": func(t *testing.T) {
			arr := [10]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(arr)}
			shrinker := Array(arb)
			if _, err := shrinker(arbitrary.Arbitrary{}, true); err == nil {
				t.Fatalf("Expected error when passed arbitrary is not a array")
			}
		},
		"ShrinkingFinishes": func(t *testing.T) {
			arr := [10]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
			elements := make([]arbitrary.Arbitrary, len(arr))

			for index, arr := range arr {
				elements[index] = arbitrary.Arbitrary{
					Value:    reflect.ValueOf(arr),
					Shrinker: Uint64(constraints.Uint64Default()),
				}
			}

			arb := arbitrary.Arbitrary{
				Value:    reflect.ValueOf(arr),
				Elements: elements,
			}
			arb.Shrinker = Array(arb)

			for arb.Shrinker != nil {
				var err error
				arb, err = arb.Shrinker(arb, true)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}

			if arb.Value.Interface().([10]uint64) != [10]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0} {
				t.Fatalf("Shrinking failed")
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
