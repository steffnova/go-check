package shrinker

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestCollectionAllElements(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"InputArbitraryIsNotChanged": func(t *testing.T) {
			arb := arbitrary.Arbitrary{
				Elements: arbitrary.Arbitraries{
					{
						Value:    reflect.ValueOf(uint64(200)),
						Shrinker: Uint64(constraints.Uint64Default()),
					},
					{
						Value:    reflect.ValueOf(uint64(200)),
						Shrinker: nil,
					},
				},
				Shrinker: CollectionAllElements(),
			}

			original := arb
			_, err := arb.Shrinker(arb, true)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			for index := range arb.Elements {
				if !reflect.DeepEqual(arb.Elements[index].Value, original.Elements[index].Value) {
					t.Fatalf("Original arbitrary element at index: %d has changed", index)
				}
			}
		},
		"ShrinkingError": func(t *testing.T) {
			randomError := fmt.Errorf("random error")
			arb := arbitrary.Arbitrary{
				Value: reflect.ValueOf("test"),
				Elements: arbitrary.Arbitraries{
					{
						Value:    reflect.ValueOf(uint64(200)),
						Shrinker: Fail(randomError),
					},
					{
						Value:    reflect.ValueOf(uint64(200)),
						Shrinker: nil,
					},
				},
				Shrinker: CollectionAllElements(),
			}

			_, err := arb.Shrinker(arb, true)
			if !errors.Is(err, randomError) {
				t.Fatalf("Expected error: %s, got: %s", randomError, err)
			}
		},
		"ShrinkAllElements": func(t *testing.T) {
			arb := arbitrary.Arbitrary{
				Elements: arbitrary.Arbitraries{
					{
						Value:    reflect.ValueOf(uint64(200)),
						Shrinker: Uint64(constraints.Uint64Default()),
					},
					{
						Value:    reflect.ValueOf(uint64(200)),
						Shrinker: Uint64(constraints.Uint64Default()),
					},
				},
				Shrinker: CollectionAllElements(),
			}

			shrink, err := arb.Shrinker(arb, true)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			for index := range arb.Elements {
				if arb.Elements[index].Value.Uint() <= shrink.Elements[index].Value.Uint() {
					t.Fatalf("Arbitrary element at index: %d hasn't changed", index)
				}
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}

}
