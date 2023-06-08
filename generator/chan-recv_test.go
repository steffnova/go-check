package generator

import (
	"errors"
	"math"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestChanRecv(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(int) {},
				ChanRecv(Int()),
			))
			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidTarget)
			}
		},
		"InvalidChannelDirection": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(chan int) {},
				ChanRecv(Int()),
			))
			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidTarget)
			}
		},
		"InvalidConstraints1": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(<-chan int) {},
				ChanRecv(Int(), constraints.Length{Min: 10, Max: 0}),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidConstraints)
			}
		},
		"InvalidConstraints2": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(<-chan int) {},
				ChanRecv(Int(), constraints.Length{Min: 10, Max: uint64(math.MaxInt64) + 1}),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidConstraints)
			}
		},
		"WithinConstraints": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(ch <-chan int) {
					if cap(ch) < 5 || cap(ch) > 20 {
						t.Errorf("Invalid channel capacity")
					}
				},
				ChanRecv(Int(), constraints.Length{Min: 5, Max: 20}),
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
