package generator

import (
	"errors"
	"math"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
)

func TestWeighted(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"NoWeights": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Weighted(nil, Int()),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidConfig) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidConfig)
			}
		},
		"Noarbitrary.Generators": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Weighted([]uint64{5}),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidConfig) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidConfig)
			}
		},
		"InvalidConfiguration": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Weighted([]uint64{5, 4}, Int()),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidConfig) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidConfig)
			}
		},
		"InvalidWieight": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Weighted([]uint64{0}, Int()),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidConfig) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidConfig)
			}
		},
		"WeightOverflow": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Weighted([]uint64{10, math.MaxUint64}, Int(), Int()),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidConfig) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidConfig)
			}
		},
		"InvalidTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(uint) {},
				Weighted([]uint64{1}, Int()),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidTarget)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
