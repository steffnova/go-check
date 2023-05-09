package shrinker

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestSlice(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"Shrink": func(t *testing.T) {
			arr := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
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
			arb.Shrinker = Slice(arb, constraints.Length{Min: 0, Max: 10})

			property := func(in []uint64) bool {
				find5, find7 := false, false
				for _, element := range in {
					if element == 5 {
						find5 = true
					}
					if element == 7 {
						find7 = true
					}
				}
				return find5 && find7
			}

			propertyFailed := property(arb.Value.Interface().([]uint64))

			for arb.Shrinker != nil {
				fmt.Println(propertyFailed)
				var err error
				arb, err = arb.Shrinker(arb, propertyFailed)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
				fmt.Println(arb.Value)
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
