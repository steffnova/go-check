package shrinker

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestShrinkerMap(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"MapOnNilShrinker": func(t *testing.T) {
			if Shrinker(nil).Map(func(int) int { return 0 }) != nil {
				t.Errorf("Mapping a nil shrinker should return nil")
			}
		},
		"MapperNotAFunction": func(t *testing.T) {
			shrinker := Uint64(constraints.Uint64{}).Map(nil)

			_, shrinker, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to mapper not being a function")
			}
		},
		"MapperInputInvalid": func(t *testing.T) {
			shrinker := Uint64(constraints.Uint64{}).Map(func() {})

			_, shrinker, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to mapper not having exactly one input")
			}
		},
		"MapperOutputInvalid": func(t *testing.T) {
			shrinker := Uint64(constraints.Uint64{}).Map(func(uint64) {})

			_, shrinker, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to mapper not having exactly one output")
			}
		},
		"MapperInputIncompatibleWithShrunkType": func(t *testing.T) {
			shrinker := Uint64(constraints.Uint64{}).
				Map(func(int64) int {
					return 0
				})

			arb := arbitrary.Arbitrary{
				Precursors: arbitrary.Arbitraries{
					{Value: reflect.ValueOf(uint64(0))}},
			}
			_, shrinker, err := shrinker(arb, true)
			if err == nil {
				t.Fatalf("Expected a failure due to mapper not having exactly one input")
			}
		},
		"ShrinkingError": func(t *testing.T) {
			shrinker := Fail(fmt.Errorf("random error")).
				Map(func(uint64) int {
					return 0
				})

			arb := arbitrary.Arbitrary{
				Precursors: arbitrary.Arbitraries{
					{Value: reflect.ValueOf(uint64(0))}},
			}
			_, shrinker, err := shrinker(arb, true)
			if err == nil {
				t.Fatalf("Expected a failure due base generator throwing an error")
			}
		},
		"MappedShrink": func(t *testing.T) {
			mapper := func(in uint64) uint64 {
				return in + 2
			}
			shrinker := Uint64(constraints.Uint64Default()).Map(mapper)

			arb := arbitrary.Arbitrary{
				Precursors: arbitrary.Arbitraries{
					{Value: reflect.ValueOf(uint64(100))}},
			}
			shrink, shrinker, err := shrinker(arb, true)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			// Shrink value is shrunk precursor (100) to a value 50 then
			// mapped to new value using mapper (50 + 2)
			if shrink.Value.Uint() != mapper(shrink.Precursors[0].Value.Uint()) {
				t.Fatalf("Invalid shrink value: %d", shrink.Value.Uint())
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}

func TestShrinkerOr(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"OrOnNilShrinker": func(t *testing.T) {
			someError := fmt.Errorf("random error")
			shrinker := Shrinker(nil).Or(Fail(someError))

			_, shrinker, err := shrinker(arbitrary.Arbitrary{}, true)
			if !errors.Is(err, someError) {
				t.Fatalf("expected error returned by Fail shrinker")
			}
		},
		"PropertyFailed": func(t *testing.T) {
			error1 := fmt.Errorf("error1")
			error2 := fmt.Errorf("error2")
			shrinker := Fail(error1).Or(Fail(error2))

			_, shrinker, err := shrinker(arbitrary.Arbitrary{}, true)
			// Testing if base shrinker is called, doesn't matter that it
			// throws an error.
			if !errors.Is(err, error1) {
				t.Fatalf("expected error: %s", error1)
			}
		},
		"PropertySucceed": func(t *testing.T) {
			error1 := fmt.Errorf("error1")
			error2 := fmt.Errorf("error2")
			shrinker := Fail(error1).Or(Fail(error2))

			_, shrinker, err := shrinker(arbitrary.Arbitrary{}, false)
			// Testing if shrinker passed to Or is called, doesn't matter that it
			// throws an error.
			if !errors.Is(err, error2) {
				t.Fatalf("expected error: %s", error2)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}

func TestShrinkerFilter(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"FilterOnNilShrinker": func(t *testing.T) {
			shrinker := Shrinker(nil).
				Filter(arbitrary.Arbitrary{}, func(int) bool {
					return false
				})
			if shrinker != nil {
				t.Errorf("Using filter on nil shrinker should return nil")
			}
		},
		"PredicateIsNotAFunction": func(t *testing.T) {
			shrinker := Shrinker(nil).Filter(arbitrary.Arbitrary{}, nil)

			_, _, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to predicate not being a function")
			}
		},
		"PredicateNumberOfInputsInvalid": func(t *testing.T) {
			shrinker := Shrinker(nil).
				Filter(arbitrary.Arbitrary{}, func() bool {
					return false
				})

			_, _, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to predicate not having exactly one input")
			}
		},
		"PredicateNumberOfOutputsInvalid": func(t *testing.T) {
			shrinker := Shrinker(nil).
				Filter(arbitrary.Arbitrary{}, func(int) {

				})

			_, _, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to predicate not having exactly one output")
			}
		},
		"PredicateOutputTypeNotBool": func(t *testing.T) {
			shrinker := Shrinker(nil).
				Filter(arbitrary.Arbitrary{}, func(int) int {
					return 0
				})

			_, _, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to predicate not having exactly one output")
			}
		},
		"ShrinkingError": func(t *testing.T) {
			randomError := fmt.Errorf("random error")
			shrinker := Fail(randomError).
				Filter(arbitrary.Arbitrary{}, func(int) bool {
					return true
				})

			_, _, err := shrinker(arbitrary.Arbitrary{}, true)
			if !errors.Is(err, randomError) {
				t.Fatalf("Expected error: %s", randomError)
			}
		},
		"ImpossiblePredicate": func(t *testing.T) {
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(uint64(5))}
			shrinker := Uint64(constraints.Uint64Default()).Filter(
				arb, func(in uint64) bool {
					return in > 115
				},
			)

			var err error
			var shrink = arb
			for shrinker != nil {
				shrink, shrinker, err = shrinker(shrink, true)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}

			if shrink.Value.Uint() != arb.Value.Uint() {
				t.Fatalf("Inital value should be returned if predicate can't be satisfied")
			}
		},
		"ShrunkValue": func(t *testing.T) {
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(uint64(115))}
			predicate := func(in uint64) bool {
				return in%2 != 0
			}
			shrinker := Uint64(constraints.Uint64Default()).Filter(arb, predicate)

			var err error
			for shrinker != nil {
				arb, shrinker, err = shrinker(arb, true)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}

			if !predicate(arb.Value.Uint()) {
				t.Fatalf("Shrunk value doesn't satisfy predicate")
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
