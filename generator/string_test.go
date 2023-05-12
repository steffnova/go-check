package generator

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestString(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			limits := constraints.String{
				Length: constraints.Length{
					Min: 0,
					Max: 20,
				},
				Rune: constraints.RuneDefault(),
			}
			Stream(0, 100, Streamer(
				func(s string) {
					if len([]rune(s)) < int(limits.Length.Min) || len([]rune(s)) > int(limits.Length.Max) {
						t.Fatalf("String length: %d is not withing limits: %+v", len(s), limits)
					}
				},
				String(limits),
			))
		},
		"InvalidLengthRange": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(s string) {},
				String(constraints.String{Length: constraints.Length{Min: 50, Max: 0}}),
			))
			if !errors.Is(err, arbitrary.ErrorInvalidConstraints) {
				t.Fatalf("Expected error: %s", arbitrary.ErrorInvalidConstraints)
			}
		},
		"InvalidRuneRange": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(s string) {},
				String(constraints.String{
					Length: constraints.Length{Min: 1, Max: 10},
					Rune:   constraints.Rune{MinCodePoint: 1000, MaxCodePoint: 0}},
				),
			))
			if !errors.Is(err, arbitrary.ErrorInvalidConstraints) {
				t.Fatalf("Expected error: %s, got: %s", arbitrary.ErrorInvalidConstraints, err)
			}
		},
		// "InvalidType": func(t *testing.T) {
		// 	err := Stream(0, 1, Streamer(
		// 		func(n uint64) {},
		// 		String(),
		// 	))
		// 	if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
		// 		t.Fatalf("Expected error because constraints are invalid: %s", err)
		// 	}
		// },
		"UnderlyingType": func(t *testing.T) {
			type newType string
			err := Stream(0, 100, Streamer(
				func(n newType) {},
				String(),
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
