package property

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

func TestInputsShrinker(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"NilShrinker": func(t *testing.T) {
			if shrinkers(arbitrary.Arbitrary{}) != nil {
				t.Fatalf("Expected shrinker to be nil")
			}
		},
		"ShrinkingError": func(t *testing.T) {
			shrinkerErr := fmt.Errorf("shrinker error")
			arb := arbitrary.Arbitrary{
				Shrinker: arbitrary.Shrinker(nil).Fail(shrinkerErr),
			}
			shrinker := shrinkers(arb)

			if _, _, err := shrinker(nil, true); !errors.Is(err, shrinkerErr) {
				t.Fatalf("Expected error: %s. Got: %s", shrinkerErr, err)
			}
		},
		"InvalidNumberOfShrinks": func(t *testing.T) {
			arb := arbitrary.Arbitrary{
				Shrinker: arbitrary.Shrinker(func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
					return arbitrary.Arbitrary{}, nil
				}),
			}

			shrinker := shrinkers(arb)
			if _, _, err := shrinker([]arbitrary.Arbitrary{{}, {}}, true); err == nil {
				t.Fatalf("Expected error for invalid number of elements returned")
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}

}

func TestInputsShrinkerLog(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"LogOnNilShrinker": func(t *testing.T) {
			shrinker := shrinkers(arbitrary.Arbitrary{}).Log(0)
			if shrinker != nil {
				t.Errorf("Using Log on nil shrinker should return nil")
			}
		},
		"ShrinkingError": func(t *testing.T) {
			shrinkingError := fmt.Errorf("shrinking error")
			shrinker := inputShrinker(nil).Fail(shrinkingError).Log(0)

			_, _, err := shrinker(arbitrary.Arbitraries{{}}, true)
			if !errors.Is(err, shrinkingError) {
				t.Fatalf("Expected error: %s", shrinkingError)
			}
		},
		"ShrunkValue": func(t *testing.T) {
			shrinks := arbitrary.Arbitraries{
				{Value: reflect.ValueOf(uint64(100))},
				{Value: reflect.ValueOf(uint64(200))},
			}

			shrinker := inputShrinker(func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputShrinker, error) {
				return shrinks, nil, nil
			}).Log(0)

			arbs, _, err := shrinker(arbitrary.Arbitraries{{}}, true)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			for index := range arbs {
				if !reflect.DeepEqual(arbs[index].Value.Interface(), shrinks[index].Value.Interface()) {
					t.Fatalf("shrink at index %d doesn't match expected value", index)
				}
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}

func TestInputsShrinkerFilter(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"FilterOnNilShrinker": func(t *testing.T) {
			shrinker := shrinkers(arbitrary.Arbitrary{}).
				Filter(func(int) bool {
					return false
				})
			if shrinker != nil {
				t.Errorf("Using filter on nil shrinker should return nil")
			}
		},
		"PredicateIsNotAFunction": func(t *testing.T) {
			shrinker := inputShrinker(func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputShrinker, error) {
				return nil, nil, nil
			}).Filter(5)

			_, _, err := shrinker(arbitrary.Arbitraries{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to predicate not being a function")
			}
		},
		"PredicateNumberOfOutputsInvalid": func(t *testing.T) {
			shrinker := inputShrinker(func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputShrinker, error) {
				return nil, nil, nil
			}).Filter(func(int) {
			})

			_, _, err := shrinker(arbitrary.Arbitraries{{}}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to predicate not having exactly one output")
			}
		},
		"PredicateOutputTypeNotBool": func(t *testing.T) {
			shrinker := inputShrinker(func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputShrinker, error) {
				return nil, nil, nil
			}).Filter(func(int) int {
				return 0
			})

			_, _, err := shrinker(arbitrary.Arbitraries{{}}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to predicate not being a Bool")
			}
		},
		"PredicateNumberOfInputsInvalid": func(t *testing.T) {
			shrinker := inputShrinker(func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputShrinker, error) {
				return nil, nil, nil
			}).Filter(func(int, int, int) bool {
				return false
			})

			_, _, err := shrinker(arbitrary.Arbitraries{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to number of arbs passed to shrinker not matching a number of predicate parameters")
			}
		},
		"ShrinkingError": func(t *testing.T) {
			shrinkingError := fmt.Errorf("shrinking error")
			shrinker := inputShrinker(nil).Fail(shrinkingError).
				Filter(func(int) bool {
					return true
				})

			_, _, err := shrinker(arbitrary.Arbitraries{{}}, true)
			if !errors.Is(err, shrinkingError) {
				t.Fatalf("Expected error: %s", shrinkingError)
			}
		},
		"ShrunkValue": func(t *testing.T) {
			arb := arbitrary.Arbitrary{
				Elements: arbitrary.Arbitraries{
					{
						Value:    reflect.ValueOf(uint64(100)),
						Shrinker: shrinker.Uint64(constraints.Uint64Default()),
					},
					{
						Value:    reflect.ValueOf(uint64(200)),
						Shrinker: shrinker.Uint64(constraints.Uint64Default()),
					},
				},
			}
			arb.Shrinker = shrinker.CollectionElements(arb)
			shrinker := shrinkers(arb).Filter(func(x, y uint64) bool {
				return x >= 50 && y >= 150

			})

			shrinks := arb.Elements
			var err error
			for shrinker != nil {
				shrinks, shrinker, err = shrinker(shrinks, true)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}

			v1 := shrinks[0].Value.Uint()
			v2 := shrinks[1].Value.Uint()
			if v1 != 50 || v2 != 150 {
				t.Fatalf("Expected shrinks to have values that satisfy filter constraints. Got: %d, %d", v1, v2)
			}

		},
		"ExhaustedShrinker": func(t *testing.T) {
			shrinks := arbitrary.Arbitraries{
				{Value: reflect.ValueOf(uint64(100))},
				{Value: reflect.ValueOf(uint64(200))},
			}

			shrinker := inputShrinker(func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputShrinker, error) {
				return arbs, nil, nil
			}).Filter(func(x, y uint64) bool {
				return false
			})

			var err error
			for shrinker != nil {
				shrinks, shrinker, err = shrinker(shrinks, true)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
