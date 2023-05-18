package shrinker

import (
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestMap(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"OriginalNotAMap": func(t *testing.T) {
			shrinker := Map(arbitrary.Arbitrary{}, constraints.LengthDefault())
			if _, err := shrinker(arbitrary.Arbitrary{}, true); err == nil {
				t.Fatalf("Expected error when original arbitrary is not map")
			}
		},
		"OriginalInsufficientElements": func(t *testing.T) {
			m := map[uint64]uint64{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(m)}
			shrinker := Map(arb, constraints.LengthDefault())
			if _, err := shrinker(arb, true); err == nil {
				t.Fatalf("Expected error when number of arbitrary elements doesn't match the slice size")
			}
		},
		"ShrinkingFinishes": func(t *testing.T) {
			m := map[uint64]uint64{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
			elements := []arbitrary.Arbitrary{}

			for key, value := range m {
				element := arbitrary.Arbitrary{
					Elements: arbitrary.Arbitraries{
						{Value: reflect.ValueOf(key), Shrinker: Uint64(constraints.Uint64Default())},
						{Value: reflect.ValueOf(value), Shrinker: Uint64(constraints.Uint64Default())},
					},
				}
				element.Shrinker = CollectionElements(element)
				elements = append(elements, element)
			}

			arb := arbitrary.Arbitrary{
				Value:    reflect.ValueOf(m),
				Elements: elements,
			}

			arb.Shrinker = Map(arb, constraints.Length{Min: 0, Max: 10})

			property := func(in map[uint64]uint64) bool {
				for key := range in {
					if key == 5 {
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
				propertyFailed = property(arb.Value.Interface().(map[uint64]uint64))

			}

			// Property holds as long there is a key with value 5 in the map.
			// Original value should be shrunk to map with one key value pair (5:0)
			result := arb.Value.Interface().(map[uint64]uint64)
			if len(result) != 1 {
				t.Fatalf("Shrunk map: %v should have only one element.", result)
			}
			if _, ok := result[5]; !ok {
				t.Fatalf("Shrunk map: %v should have a key: %d: ", result, 5)
			}
			if result[5] != 0 {
				t.Fatal("Shrunk value for key 5 should be 0")
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}

}
