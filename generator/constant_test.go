package generator

import (
	"errors"
	"testing"
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
					if in != ErrorInvalidTarget {
						t.Fatalf("invalid generated constant value: %d", in)
					}
				},
				Constant(ErrorInvalidTarget),
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
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s':", ErrorInvalidTarget)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
