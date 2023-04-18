package generator

import (
	"errors"
	"math"
	"testing"

	"github.com/steffnova/go-check/constraints"
)

func TestSlice(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"InvalidTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Slice(Int()),
			))

			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidTarget)
			}
		},
		"InvalidConstraints": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func([]int) {},
				Slice(Int(), constraints.Length{Min: 100, Max: 0}),
			))

			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"InvalidMaxLenght": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func([]int) {},
				Slice(Int(), constraints.Length{Max: uint64(math.MaxInt64) + 1}),
			))

			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"InvalidElementTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func([]uint) {},
				Slice(Int()),
			))

			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidTarget)
			}
		},
		"WithinConstraints": func(t *testing.T) {
			constraint := constraints.Length{Min: 5, Max: 20}
			err := Stream(0, 100, Streamer(
				func(nums []int) {
					if len(nums) < int(constraint.Min) || len(nums) > int(constraint.Max) {
						t.Fatalf("Slice length: %d is not within constraints %#v", len(nums), constraint)
					}
				},
				Slice(Int(), constraint),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
