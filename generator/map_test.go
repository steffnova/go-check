package generator

import (
	"errors"
	"math"
	"testing"

	"github.com/steffnova/go-check/constraints"
)

func TestMap(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"InvalidTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Map(Int(), Int()),
			))

			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidTarget)
			}
		},
		"InvalidConstraints": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(map[int]int) {},
				Map(Int(), Int(), constraints.Length{Min: 100, Max: 10}),
			))

			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"InvalidConstraints2": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(map[int]int) {},
				Map(Int(), Int(), constraints.Length{Max: uint64(math.MaxInt64) + 1}),
			))

			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"InvalidKeyTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(map[uint]int) {},
				Map(Int(), Int()),
			))

			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidTarget)
			}
		},
		"InvalidValueTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(map[int]uint) {},
				Map(Int(), Int()),
			))

			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidTarget)
			}
		},
		"MapSizeWithinConstraints": func(t *testing.T) {
			constraint := constraints.Length{Min: 5, Max: 100}
			err := Stream(0, 100, Streamer(
				func(in map[int]int) {
					if len(in) < int(constraint.Min) || len(in) > int(constraint.Max) {
						t.Fatalf("Generated map size %d is not withing it's constraints: %#v", len(in), constraint)
					}
				},
				Map(Int(), Int(), constraint),
			))

			if err != nil {
				t.Fatalf("Unexpected error: '%s'", err)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
