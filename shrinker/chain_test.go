package shrinker

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestChain(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"Chaining0Shrinkers": func(t *testing.T) {
			if Chain() != nil {
				t.Fatalf("Chaining 0 shrinkers should return nil shrinker")
			}
		},
		"ShrinkingError": func(t *testing.T) {
			randomError := fmt.Errorf("random error")
			shrinker := Chain(Fail(randomError), Uint64(constraints.Uint64{}))

			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if !errors.Is(err, randomError) {
				t.Errorf("Expected error: %s, got: %s", randomError, err)
			}
		},
		"ShrinkingEnds": func(t *testing.T) {
			shrinker1 := Uint64(constraints.Uint64{Min: 50, Max: 100})
			shrinker2 := Uint64(constraints.Uint64{Min: 10, Max: 60})

			arb := arbitrary.Arbitrary{
				Value:    reflect.ValueOf(uint64(80)),
				Shrinker: Chain(shrinker1, shrinker2),
			}

			for arb.Shrinker != nil {
				var err error
				arb, err = arb.Shrinker(arb, true)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}

			// First shrinker, shrinks original value to it's lower bound (50)
			// Second shrinker, should further shrink value to it's lower bound (10)
			if arb.Value.Uint() != 10 {
				t.Fatalf("Expected final shrink to be 10. Got: %d", arb.Value.Uint())
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
