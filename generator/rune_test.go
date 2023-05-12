package generator

import (
	"errors"
	"math"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestRune(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(string) {},
				Rune(),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidTarget)
			}
		},
		"InvalidConstraints": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(rune) {},
				Rune(constraints.Rune{MinCodePoint: 100, MaxCodePoint: 0}),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidConstraints)
			}
		},
		"InvalidMaxCodePoint": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(rune) {},
				Rune(constraints.Rune{MaxCodePoint: math.MaxInt32}),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidConstraints)
			}
		},
		"InvalidMinCodePoint": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(rune) {},
				Rune(constraints.Rune{MinCodePoint: math.MinInt32}),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidConstraints)
			}
		},
		"WithinConstraints": func(t *testing.T) {
			// Constraints are from uppercase A-Z
			constraint := constraints.Rune{MinCodePoint: 65, MaxCodePoint: 90}
			err := Stream(0, 100, Streamer(
				func(c rune) {
					if int32(c) < constraint.MinCodePoint || int32(c) > constraint.MaxCodePoint {
						t.Fatalf("Rune %c is not within constraints: [%c-%c]", c, constraint.MinCodePoint, constraint.MaxCodePoint)
					}
				},
				Rune(constraint),
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
