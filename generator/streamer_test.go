package generator

import (
	"errors"
	"testing"
)

func TestStreamer(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"InvalidStreamTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				0,
				Int(),
			))

			if !errors.Is(err, ErrorStream) {
				t.Fatalf("Expected error; '%s'", ErrorStream)
			}
		},
		"InvalidStreamTargetInput": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func() {},
				Int(),
			))

			if !errors.Is(err, ErrorStream) {
				t.Fatalf("Expected error; '%s'", ErrorStream)
			}
		},
		"InvalidStreamTargetOutput": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) int { return 0 },
				Int(),
			))

			if !errors.Is(err, ErrorStream) {
				t.Fatalf("Expected error; '%s'", ErrorStream)
			}
		},
		"InvalidTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(uint) {},
				Int(),
			))

			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error; '%s'", ErrorInvalidTarget)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
