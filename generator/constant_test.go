package generator

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
)

func TestConstant(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"Nil": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(in []int) {
					if in != nil {
						t.Fatalf("Failed to generate nil slice")
					}
				},
				Constant(nil),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Int": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(in int) {
					if in != 20 {
						t.Fatalf("Invalid generated constant value: %d", in)
					}
				},
				Constant(20),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Interface": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(in error) {
					if in != arbitrary.ErrorInvalidTarget {
						t.Fatalf("invalid generated constant value: %d", in)
					}
				},
				Constant(arbitrary.ErrorInvalidTarget),
			))
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(n int) {},
				Constant(0.5),
			))
			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s':", arbitrary.ErrorInvalidTarget)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}

func TestConstantFrom(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(n int) {},
				ConstantFrom("test", 0.5),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expescted error: '%s':", arbitrary.ErrorInvalidTarget)
			}
		},
		"OneOfValues": func(t *testing.T) {
			values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			err := Stream(0, 100, Streamer(
				func(n int) {
					for _, element := range values {
						if element == n {
							return
						}
					}
					t.Fatalf("n: is not one of values: %v", values)
				},
				ConstantFrom(values[0], values[1:]),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expescted error: '%s':", arbitrary.ErrorInvalidTarget)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
