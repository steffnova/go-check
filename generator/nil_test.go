package generator

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
)

func TestNil(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"InvalidTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Nil(),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidTarget)
			}
		},
		"Nil": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(ch chan int, n *int, m map[int]int, slice []int, in interface{}, fn func()) {
					switch {
					case ch != nil:
						t.Fatalf("Expected channel to be nil")
					case n != nil:
						t.Fatalf("Expected pointer to be nil")
					case m != nil:
						t.Fatalf("Expected map to be nil")
					case slice != nil:
						t.Fatalf("Expected slice to be nil")
					case in != nil:
						t.Fatalf("Expected interface to be nil")
					case fn != nil:
						t.Fatalf("Expected function to be nil")
					}
				},
				Nil(),
				Nil(),
				Nil(),
				Nil(),
				Nil(),
				Nil(),
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
