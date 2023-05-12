package generator

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
)

func TestFunc(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(n int) {},
				Func(),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: %s", arbitrary.ErrorInvalidTarget)
			}
		},
		"InvalidConfiguration": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(func(int) bool) {},
				Func(),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidConfig) {
				t.Fatalf("Expected error: %s", arbitrary.ErrorInvalidConfig)
			}
		},
		"InvalidFuncOutputTarget": func(t *testing.T) {
			Stream(0, 10, Streamer(
				func(in func(int) bool) {
					defer func() {
						err := recover().(error)
						if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
							t.Fatalf("Expected error: %s", arbitrary.ErrorInvalidTarget)
						}
					}()
					in(1)
				},
				Func(Int()),
			))

		},
		"TwoDifferentFunctions": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(fn1, fn2 func(int) int) {
					if fn1(10) == fn2(10) {
						t.Fatalf("Same output for two functions with same input")
					}
				},
				Func(Int()),
				Func(Int()),
			))

			if err != nil {
				t.Fatalf("Unexpected error; %s", err)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
