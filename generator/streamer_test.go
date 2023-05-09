package generator

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
)

func TestStreamer(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"InvalidStreamTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				0,
				Int(),
			))

			if !errors.Is(err, arbitrary.ErrorStream) {
				t.Fatalf("Expected error; '%s'", arbitrary.ErrorStream)
			}
		},
		"InvalidStreamTargetInput": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func() {},
				Int(),
			))

			if !errors.Is(err, arbitrary.ErrorStream) {
				t.Fatalf("Expected error; '%s'", arbitrary.ErrorStream)
			}
		},
		"InvalidStreamTargetOutput": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) int { return 0 },
				Int(),
			))

			if !errors.Is(err, arbitrary.ErrorStream) {
				t.Fatalf("Expected error; '%s'", arbitrary.ErrorStream)
			}
		},
		"InvalidTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(uint) {},
				Int(),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error; '%s'", arbitrary.ErrorInvalidTarget)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
