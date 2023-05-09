package arbitrary_test

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

func TestShrinkerMap(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"MapOnNilShrinker": func(t *testing.T) {
			if arbitrary.Shrinker(nil).Map(func(int) int { return 0 }) != nil {
				t.Errorf("Mapping a nil shrinker should return nil")
			}
		},
		"MapperNotAFunction": func(t *testing.T) {
			shrinker := shrinker.Uint64(constraints.Uint64{}).Map(nil)

			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to mapper not being a function")
			}
		},
		"MapperInputInvalid": func(t *testing.T) {
			shrinker := shrinker.Uint64(constraints.Uint64{}).Map(func() {})

			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to mapper not having exactly one input")
			}
		},
		"MapperOutputInvalid": func(t *testing.T) {
			shrinker := shrinker.Uint64(constraints.Uint64{}).Map(func(uint64) {})

			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to mapper not having exactly one output")
			}
		},
		"MapperInputIncompatibleWithShrunkType": func(t *testing.T) {
			shrinker := shrinker.Uint64(constraints.Uint64{}).
				Map(func(int64) int {
					return 0
				})

			arb := arbitrary.Arbitrary{
				Precursors: arbitrary.Arbitraries{
					{Value: reflect.ValueOf(uint64(0))}},
			}
			_, err := shrinker(arb, true)
			if err == nil {
				t.Fatalf("Expected a failure due to mapper not having exactly one input")
			}
		},
		"ShrinkingError": func(t *testing.T) {
			shrinker := shrinker.Fail(fmt.Errorf("random error")).
				Map(func(uint64) int {
					return 0
				})

			arb := arbitrary.Arbitrary{
				Precursors: arbitrary.Arbitraries{
					{Value: reflect.ValueOf(uint64(0))}},
			}
			_, err := shrinker(arb, true)
			if err == nil {
				t.Fatalf("Expected a failure due base generator throwing an error")
			}
		},
		"MappedShrink": func(t *testing.T) {
			mapper := func(in uint64) uint64 {
				return in + 2
			}
			shrinker := shrinker.Uint64(constraints.Uint64Default()).Map(mapper)

			arb := arbitrary.Arbitrary{
				Precursors: arbitrary.Arbitraries{
					{Value: reflect.ValueOf(uint64(100))}},
			}
			shrink, err := shrinker(arb, true)
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

func TestShrinkerFilter(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"FilterOnNilShrinker": func(t *testing.T) {
			shrinker := arbitrary.Shrinker(nil).
				Filter(func(int) bool {
					return false
				})
			if shrinker != nil {
				t.Errorf("Using filter on nil shrinker should return nil")
			}
		},
		"PredicateIsNotAFunction": func(t *testing.T) {
			shrinker := shrinker.Uint64(constraints.Uint64{}).Filter(nil)

			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to predicate not being a function")
			}
		},
		"PredicateNumberOfInputsInvalid": func(t *testing.T) {
			shrinker := shrinker.Uint64(constraints.Uint64{}).
				Filter(func() bool {
					return false
				})

			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to predicate not having exactly one input")
			}
		},
		"PredicateNumberOfOutputsInvalid": func(t *testing.T) {
			shrinker := shrinker.Uint64(constraints.Uint64{}).
				Filter(func(int) {

				})

			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to predicate not having exactly one output")
			}
		},
		"PredicateOutputTypeNotBool": func(t *testing.T) {
			shrinker := shrinker.Uint64(constraints.Uint64{}).
				Filter(func(int) int {
					return 0
				})

			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if err == nil {
				t.Fatalf("Expected a failure due to predicate not having exactly one output")
			}
		},
		"ShrinkingError": func(t *testing.T) {
			randomError := fmt.Errorf("random error")
			shrinker := shrinker.Fail(randomError).
				Filter(func(int) bool {
					return true
				})

			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if !errors.Is(err, randomError) {
				t.Fatalf("Expected error: %s", randomError)
			}
		},
		"ImpossiblePredicate": func(t *testing.T) {
			arb := arbitrary.Arbitrary{
				Value: reflect.ValueOf(uint64(5)),
				Shrinker: shrinker.Uint64(constraints.Uint64Default()).Filter(
					func(in uint64) bool {
						return in > 115
					},
				),
			}

			var err error
			var shrink = arb
			for shrink.Shrinker != nil {
				shrink, err = shrink.Shrinker(shrink, true)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}

			if shrink.Value.Uint() != arb.Value.Uint() {
				t.Fatalf("Inital value should be returned if predicate can't be satisfied")
			}
		},
		"ShrunkValue": func(t *testing.T) {
			predicate := func(in uint64) bool {
				return in%2 != 0
			}

			arb := arbitrary.Arbitrary{
				Value:    reflect.ValueOf(uint64(115)),
				Shrinker: shrinker.Uint64(constraints.Uint64Default()).Filter(predicate),
			}

			var err error
			for arb.Shrinker != nil {
				arb, err = arb.Shrinker(arb, true)
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

func TestShrinkerBind(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"BinderIsNil": func(t *testing.T) {
			arb := arbitrary.Arbitrary{Precursors: arbitrary.Arbitraries{{}}}

			shrinker := arbitrary.Shrinker(nil).Bind(nil, arbitrary.Arbitrary{})
			if _, err := shrinker(arb, true); err == nil {
				t.Fatalf("expected error because binder is nil")
			}
		},
		"PrecursorError": func(t *testing.T) {
			randomError := fmt.Errorf("random error")
			arb := arbitrary.Arbitrary{
				Precursors: arbitrary.Arbitraries{
					{
						Shrinker: arbitrary.Shrinker(nil).Fail(randomError),
					},
				},
			}

			binder := func(arbitrary.Arbitrary) (arbitrary.Arbitrary, error) {
				return arbitrary.Arbitrary{}, randomError
			}

			arb.Shrinker = arbitrary.Shrinker(nil).Bind(binder, arb)
			_, err := arb.Shrinker(arb, false)

			if !errors.Is(err, randomError) {
				t.Fatalf("Expected error: %s, got: %s", randomError, err)
			}
		},
		"BinderError": func(t *testing.T) {
			randomError := fmt.Errorf("random error")

			binder := func(arbitrary.Arbitrary) (arbitrary.Arbitrary, error) {
				return arbitrary.Arbitrary{}, randomError
			}
			arb := arbitrary.Arbitrary{
				Precursors: arbitrary.Arbitraries{
					{
						Shrinker: arbitrary.Shrinker(func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
							return arb, nil
						}),
					},
				},
			}

			arb.Shrinker = arbitrary.Shrinker(nil).Bind(binder, arb)

			_, err := arb.Shrinker(arb, false)

			if !errors.Is(err, randomError) {
				t.Fatalf("Expected error: %s, got: %s", randomError, err)
			}
		},
		"RootShrinkerIsNilPropertyFalsified": func(t *testing.T) {
			arb := arbitrary.Arbitrary{
				Value: reflect.ValueOf(uint64(10)),
				Precursors: arbitrary.Arbitraries{
					{Value: reflect.ValueOf(uint64(5))},
				},
			}
			binder := func(arb arbitrary.Arbitrary) (arbitrary.Arbitrary, error) {
				return arbitrary.Arbitrary{}, nil
			}

			arb.Shrinker = arbitrary.Shrinker(nil).Bind(binder, arb)

			shrink, err := arb.Shrinker(arb, true)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if arb.Value.Uint() != shrink.Value.Uint() {
				t.Fatalf("Shrink should be arb if property failed")
			}
		},
		"RootShrinkerIsNilPropertyHolds": func(t *testing.T) {
			arb1 := arbitrary.Arbitrary{Precursors: arbitrary.Arbitraries{{Value: reflect.ValueOf(uint64(5))}}}
			arb2 := arbitrary.Arbitrary{Precursors: arbitrary.Arbitraries{{Value: reflect.ValueOf(uint64(10))}}}
			binder := func(arb arbitrary.Arbitrary) (arbitrary.Arbitrary, error) {
				return arbitrary.Arbitrary{}, nil
			}

			arb1.Shrinker = arbitrary.Shrinker(nil).Bind(binder, arb2)

			shrink, err := arb1.Shrinker(arb1, false)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if !reflect.DeepEqual(arb2, shrink) {
				t.Fatalf("Shrink should be arb2 if property holds")
			}
		},
		"BoundShrink": func(t *testing.T) {
			arb := arbitrary.Arbitrary{
				Value: reflect.ValueOf(int(20)),
				Precursors: arbitrary.Arbitraries{
					{
						Value:    reflect.ValueOf(uint64(5)),
						Shrinker: shrinker.Uint64(constraints.Uint64Default()),
					},
				},
			}

			binder := func(arbitrary.Arbitrary) (arbitrary.Arbitrary, error) {
				return arb, nil
			}

			arb.Shrinker = arb.Shrinker.Bind(binder, arb)

			shrink, err := arb.Shrinker(arb, false)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if arb.Value.Int() != shrink.Value.Int() {
				t.Fatalf("Shrink should be arb2 if property holds")
			}
		},
		"ShrinkIsTheSameTypeAsPassedArbitrary": func(t *testing.T) {
			arb1 := arbitrary.Arbitrary{
				Value: reflect.ValueOf(uint64(10)),
				Precursors: arbitrary.Arbitraries{
					{Value: reflect.ValueOf(uint64(5))},
				},
			}

			binder := func(arb arbitrary.Arbitrary) (arbitrary.Arbitrary, error) {
				return arbitrary.Arbitrary{
					Value:      arb1.Value,
					Precursors: arbitrary.Arbitraries{arb},
					Shrinker:   shrinker.Uint64(constraints.Uint64Default()),
				}, nil
			}

			arb1.Shrinker = shrinker.Uint64(constraints.Uint64Default()).Bind(binder, arb1)

			shrink := arb1
			for shrink.Shrinker != nil {
				var err error
				shrink, err = shrink.Shrinker(shrink, true)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}

				if shrink.CompareType(arb1) != nil {
					t.Fatalf("invalid shrink type")
				}
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}

}
func TestShrinkerTransformAfter(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"OnNilShrinker": func(t *testing.T) {
			transform := func(in arbitrary.Arbitrary) arbitrary.Arbitrary {
				return in
			}
			if arbitrary.Shrinker(nil).TransformAfter(transform) != nil {
				t.Fatalf("arbitrary.Shrinker should be nil")
			}
		},
		"NilTransform": func(t *testing.T) {
			shrinker := arbitrary.Shrinker(func(arbitrary.Arbitrary, bool) (arbitrary.Arbitrary, error) {
				return arbitrary.Arbitrary{}, nil
			}).TransformAfter(nil)
			if _, err := shrinker(arbitrary.Arbitrary{}, true); err == nil {
				t.Fatalf("Expected error when transfrom is nil")
			}
		},
		"ShrinkingError": func(t *testing.T) {
			randomError := fmt.Errorf("randomError")
			transform := func(in arbitrary.Arbitrary) arbitrary.Arbitrary {
				return in
			}

			shrinker := shrinker.Fail(randomError).TransformAfter(transform)
			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if !errors.Is(err, randomError) {
				t.Fatalf("Expected error: %s", randomError)
			}
		},
		"Transform": func(t *testing.T) {
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(uint64(100))}
			transform := func(in arbitrary.Arbitrary) arbitrary.Arbitrary {
				in.Elements = make(arbitrary.Arbitraries, 10)
				return in
			}

			arb.Shrinker = shrinker.Uint64(constraints.Uint64Default()).TransformAfter(transform)
			shrink, err := arb.Shrinker(arb, true)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if len(shrink.Elements) == 0 {
				t.Fatal("Excepted transformAfter to change shrink arbitrary")
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}

func TestShrinkerTransformBefore(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"OnNilShrinker": func(t *testing.T) {
			transform := func(in arbitrary.Arbitrary) arbitrary.Arbitrary {
				return in
			}
			if arbitrary.Shrinker(nil).TransformBefore(transform) != nil {
				t.Fatalf("arbitrary.Shrinker should be nil")
			}
		},
		"NilTransform": func(t *testing.T) {
			shrinker := arbitrary.Shrinker(func(arbitrary.Arbitrary, bool) (arbitrary.Arbitrary, error) {
				return arbitrary.Arbitrary{}, nil
			}).TransformBefore(nil)
			if _, err := shrinker(arbitrary.Arbitrary{}, true); err == nil {
				t.Fatalf("Expected error when transfrom is nil")
			}
		},
		"ShrinkingError": func(t *testing.T) {
			randomError := fmt.Errorf("randomError")
			transform := func(in arbitrary.Arbitrary) arbitrary.Arbitrary {
				return in
			}

			shrinker := shrinker.Fail(randomError).TransformBefore(transform)
			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if !errors.Is(err, randomError) {
				t.Fatalf("Expected error: %s", randomError)
			}
		},
		"Transform": func(t *testing.T) {
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(uint64(100))}
			transform := func(in arbitrary.Arbitrary) arbitrary.Arbitrary {
				in.Elements = make(arbitrary.Arbitraries, 10)
				return in
			}

			shrinker := arbitrary.Shrinker(func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
				if len(arb.Elements) == 0 {
					t.Fatal("Excepted transformBefore to change shrink arbitrary")
				}
				return arb, nil
			}).TransformBefore(transform)
			_, err := shrinker(arb, true)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
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
			shrinker := arbitrary.Shrinker(nil).Or(shrinker.Fail(someError))

			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if !errors.Is(err, someError) {
				t.Fatalf("expected error returned by Fail shrinker")
			}
		},
		"PropertyFailed": func(t *testing.T) {
			error1 := fmt.Errorf("error1")
			error2 := fmt.Errorf("error2")
			shrinker := shrinker.Fail(error1).Or(shrinker.Fail(error2))

			_, err := shrinker(arbitrary.Arbitrary{}, true)
			// Testing if base shrinker is called, doesn't matter that it
			// throws an error.
			if !errors.Is(err, error1) {
				t.Fatalf("expected error: %s", error1)
			}
		},
		"PropertySucceed": func(t *testing.T) {
			error1 := fmt.Errorf("error1")
			error2 := fmt.Errorf("error2")
			shrinker := shrinker.Fail(error1).Or(shrinker.Fail(error2))

			_, err := shrinker(arbitrary.Arbitrary{}, false)
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

func TestShrinkerRetry(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"RetryOnNilShrinker": func(t *testing.T) {
			if arbitrary.Shrinker(nil).Retry(100, 100, arbitrary.Arbitrary{}) != nil {
				t.Errorf("Calling Retry on a nil shrinker should return nil")
			}
		},
		"ShrinkingError": func(t *testing.T) {
			arb := arbitrary.Arbitrary{}
			shrinker := shrinker.Fail(fmt.Errorf("random error")).Retry(10, 0, arb)

			if _, err := shrinker(arb, true); err == nil {
				t.Fatalf("Expected error")
			}
		},
		"UseAllRetries": func(t *testing.T) {
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(uint64(1000))}
			arb.Shrinker = shrinker.Uint64(constraints.Uint64Default()).Retry(5, 5, arb)

			var err error
			var shrink = arb
			for shrink.Shrinker != nil {
				shrink, err = shrink.Shrinker(shrink, false)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}

			if shrink.Value.Uint() != math.MaxUint64 {
				t.Fatalf("Expected shrink value to be max uint64")
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}

func TestShrinkerValidate(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"OnNilShrinker": func(t *testing.T) {
			validation := func(in arbitrary.Arbitrary) error {
				return nil
			}
			if arbitrary.Shrinker(nil).Validate(validation) != nil {
				t.Fatalf("arbitrary.Shrinker should be nil")
			}
		},
		"NilValidation": func(t *testing.T) {
			shrinker := arbitrary.Shrinker(func(arbitrary.Arbitrary, bool) (arbitrary.Arbitrary, error) {
				return arbitrary.Arbitrary{}, nil
			}).Validate(nil)
			if _, err := shrinker(arbitrary.Arbitrary{}, true); err == nil {
				t.Fatalf("Expected error when validation is nil")
			}
		},
		"ShrinkingError": func(t *testing.T) {
			randomError := fmt.Errorf("randomError")
			validation := func(in arbitrary.Arbitrary) error {
				return nil
			}

			shrinker := shrinker.Fail(randomError).Validate(validation)
			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if !errors.Is(err, randomError) {
				t.Fatalf("Expected error: %s", randomError)
			}
		},
		"ValidationError": func(t *testing.T) {
			validationError := fmt.Errorf("validationError")
			validation := func(in arbitrary.Arbitrary) error {
				return validationError
			}

			shrinker := arbitrary.Shrinker(func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
				return arbitrary.Arbitrary{}, nil
			}).Validate(validation)

			_, err := shrinker(arbitrary.Arbitrary{}, true)
			if !errors.Is(err, validationError) {
				t.Fatalf("Expected error: %s", validationError)
			}
		},
		"ShrinkValue": func(t *testing.T) {
			validation := func(in arbitrary.Arbitrary) error {
				return nil
			}

			arb := arbitrary.Arbitrary{
				Value: reflect.ValueOf(uint64(0)),
				Shrinker: arbitrary.Shrinker(func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
					return arb, nil
				}).Validate(validation),
			}
			shrink, err := arb.Shrinker(arb, true)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if !reflect.DeepEqual(arb.Value, shrink.Value) {
				t.Fatalf("Expected shrink to be of the same value as arb")
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
