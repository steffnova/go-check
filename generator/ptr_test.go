package generator

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestPtrTo(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"InvalidTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Ptr(Int()),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidTarget)
			}
		},
		"InvalidElementTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(*uint) {},
				Ptr(Int()),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidTarget)
			}
		},
		"NonNilPtr": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(n *int) {
					if n == nil {
						t.Fatalf("Nil pointer value generated")
					}
				},
				Ptr(Int(), constraints.Ptr{NilFrequency: 0}),
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
