package shrinker

import (
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestStruct(t *testing.T) {
	type Point struct {
		X uint64
		Y uint64
	}

	testCases := map[string]func(t *testing.T){
		"OriginalNotAStruct": func(t *testing.T) {
			shrinker := Struct(arbitrary.Arbitrary{})
			if _, err := shrinker(arbitrary.Arbitrary{}, true); err == nil {
				t.Fatalf("Expected error when original arbitrary has no elements")
			}
		},
		"OriginalInsufficientElements": func(t *testing.T) {
			point := Point{X: 10, Y: 10}
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(point)}
			shrinker := Struct(arb)
			if _, err := shrinker(arb, true); err == nil {
				t.Fatalf("Expected error when passed arbitrary is not a array")
			}
		},
		"ShrinkingFinishes": func(t *testing.T) {
			point := Point{X: 10, Y: 10}
			elements := make([]arbitrary.Arbitrary, 2)
			elements[0] = arbitrary.Arbitrary{
				Value:    reflect.ValueOf(point.X),
				Shrinker: Uint64(constraints.Uint64Default()),
			}
			elements[1] = arbitrary.Arbitrary{
				Value:    reflect.ValueOf(point.Y),
				Shrinker: Uint64(constraints.Uint64Default()),
			}

			arb := arbitrary.Arbitrary{
				Value:    reflect.ValueOf(point),
				Elements: elements,
			}
			arb.Shrinker = Struct(arb)

			for arb.Shrinker != nil {
				var err error
				arb, err = arb.Shrinker(arb, true)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}

			expected := Point{X: 0, Y: 0}
			result := arb.Value.Interface().(Point)
			if expected != result {
				t.Fatalf("Shrinking failed")
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}

}
