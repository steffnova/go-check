package property

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

func TestDefine(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"GeneratorNil": func(t *testing.T) {
			property := Define(
				nil,
				Predicate(func() error { return nil }),
			)

			if _, err := property(nil, constraints.Bias{}); !errors.Is(err, ErrorPropertyConfig) {
				t.Fatalf("Expected error: %s. Got: %s", ErrorPropertyConfig, err)
			}
		},
		"PredicateNil": func(t *testing.T) {
			property := Define(
				Inputs(generator.Int()),
				nil,
			)

			if _, err := property(nil, constraints.Bias{}); !errors.Is(err, ErrorPropertyConfig) {
				t.Fatalf("Expected error: %s. Got: %s", ErrorPropertyConfig, err)
			}
		},
		"GeneratorError": func(t *testing.T) {
			generator := InputsGenerator(func([]reflect.Type, constraints.Bias, arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
				return nil, nil, ErrorInputs
			})

			predicate := Predicate(func(x int) error {
				return nil
			})

			_, err := Define(generator, predicate)(nil, constraints.Bias{})
			if !errors.Is(err, ErrorInputs) {
				t.Fatalf("Expected error: %s. Got: %s", ErrorInputs, err)
			}
		},
		"ShrinkingError": func(t *testing.T) {
			shrinkerErr := fmt.Errorf("shrinking error")
			generator := InputsGenerator(func(targets []reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
				arb, _ := generator.Int()(targets[0], bias, r)
				shrinker := inputShrinker(func(arbs arbitrary.Arbitraries, propertyFailed bool) (arbitrary.Arbitraries, inputShrinker, error) {
					return nil, nil, shrinkerErr
				})
				return arbitrary.Arbitraries{arb}, shrinker, nil
			})

			predicate := Predicate(func(x int) error {
				return fmt.Errorf("property error")
			})

			r := arbitrary.RandomNumber{Rand: rand.New(rand.NewSource(0))}
			_, err := Define(generator, predicate)(r, constraints.Bias{})
			if !errors.Is(err, shrinkerErr) {
				t.Fatalf("Expected error: %s. Got: %s", shrinkerErr, err)
			}
		},
		"PropertyPass": func(t *testing.T) {
			property := Define(
				Inputs(),
				Predicate(func() error {
					return nil
				}))

			r := arbitrary.RandomNumber{Rand: rand.New(rand.NewSource(0))}
			_, err := property(r, constraints.Bias{})

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"PropertyFailed": func(t *testing.T) {
			propertyError := fmt.Errorf("property failed")
			property := Define(
				Inputs(generator.Int()),
				Predicate(func(x int) error {
					return propertyError
				}))

			r := arbitrary.RandomNumber{Rand: rand.New(rand.NewSource(0))}
			details, err := property(r, constraints.Bias{})
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if !errors.Is(details.FailureReason, propertyError) {
				t.Fatalf("Expected failure reason: %s. Got: %s", propertyError, details.FailureReason)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
