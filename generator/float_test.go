package generator

import (
	"errors"
	"math"
	"testing"

	"github.com/steffnova/go-check/constraints"
)

func TestFloat64(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			floatRange := constraints.Float64{Min: -50, Max: 50}
			Stream(0, 100, Streamer(
				func(n float64) {
					if n < floatRange.Min || n > floatRange.Max {
						t.Fatalf("Generated value is not withing given range: [%f, %f]", floatRange.Min, floatRange.Max)
					}
				},
				Float64(floatRange),
			))
		},
		"InvalidRange": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n float64) {},
				Float64(constraints.Float64{Min: 50, Max: -50}),
			))
			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"LowerRangeInvalid": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n float64) {},
				Float64(constraints.Float64{Min: math.Inf(-1), Max: 0}),
			))
			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"UpperRangeInvalid": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n float64) {},
				Float64(constraints.Float64{Min: 0, Max: math.Inf(0)}),
			))
			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Float64(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s':", ErrorInvalidTarget)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType float64
			err := Stream(0, 100, Streamer(
				func(n newType) {},
				Float64(),
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

func TestFloat32(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			floatRange := constraints.Float32{Min: -50, Max: 50}
			Stream(0, 100, Streamer(
				func(n float32) {
					if n < floatRange.Min || n > floatRange.Max {
						t.Fatalf("Generated value is not withing given range: [%f, %f]", floatRange.Min, floatRange.Max)
					}
				},
				Float32(floatRange),
			))
		},
		"InvalidRange": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n float32) {},
				Float32(constraints.Float32{Min: 50, Max: -50}),
			))
			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"LowerRangeInvalid": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n float32) {},
				Float32(constraints.Float32{Min: float32(math.Inf(-1)), Max: 0}),
			))
			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"UpperRangeInvalid": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n float32) {},
				Float32(constraints.Float32{Min: 0, Max: float32(math.Inf(0))}),
			))
			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Float32(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s':", ErrorInvalidTarget)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType float32
			err := Stream(0, 100, Streamer(
				func(n newType) {},
				Float32(),
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
