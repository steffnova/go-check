package shrinker

import (
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestUint(t *testing.T) {
	testCase := map[string]func(t *testing.T){
		"Kind": func(t *testing.T) {
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(int(0))}
			_, err := Uint64(constraints.Uint64Default())(arb, false)

			if err == nil {
				t.Fatalf("Expected error because arb is int and shrinker is Uint64")
			}
		},
		"Constraints": func(t *testing.T) {
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(uint64(10))}
			limits := constraints.Uint64{Max: 2, Min: 10}

			_, err := Uint64(limits)(arb, false)

			if err == nil {
				t.Fatalf("Expected error because limit's upper bound is higher then it's lower bound")
			}
		},
		"NotWithinConstraints": func(t *testing.T) {
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(uint64(1))}
			limits := constraints.Uint64{Max: 20, Min: 10}

			_, err := Uint64(limits)(arb, false)

			if err == nil {
				t.Fatalf("Expected error because limit's upper bound is higher then it's lower bound")
			}
		},
		"Shrink": func(t *testing.T) {
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(uint64(50))}
			limits := constraints.Uint64{Max: 100, Min: 0}

			shrink, err := Uint64(limits)(arb, true)

			if err != nil {
				t.Fatalf("Unexpected error : %s", err)
			}

			if shrink.Value.Uint() > arb.Value.Uint() {
				t.Fatalf("Shrunk value: %d is bigger then original: %d", shrink.Value.Uint(), arb.Value.Uint())
			}
		},
		"Unshrink": func(t *testing.T) {
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(uint64(50))}
			limits := constraints.Uint64{Max: 100, Min: 0}

			shrink, err := Uint64(limits)(arb, false)

			if err != nil {
				t.Fatalf("Unexpected error : %s", err)
			}

			if shrink.Value.Uint() < arb.Value.Uint() {
				t.Fatalf("Shrunk value: %d is smaller then original: %d", shrink.Value.Uint(), arb.Value.Uint())
			}
		},
		"ShrinkTowardsLowerBound": func(t *testing.T) {
			limits := constraints.Uint64{Max: 100, Min: 20}
			arb := arbitrary.Arbitrary{
				Value:    reflect.ValueOf(uint64(50)),
				Shrinker: Uint64(limits),
			}

			for arb.Shrinker != nil {
				var err error
				arb, err = arb.Shrinker(arb, true)
				if err != nil {
					t.Fatalf("Unexpected error : %s", err)
				}
			}

			if arb.Value.Uint() != limits.Min {
				t.Errorf("Shrunk value: %d should be exactly the value of lower bound limit: %d", arb.Value.Uint(), limits.Min)
			}
		},
		"ShrinkTowardsUpperBound": func(t *testing.T) {
			limits := constraints.Uint64{Max: 100, Min: 20}
			arb := arbitrary.Arbitrary{
				Value:    reflect.ValueOf(uint64(50)),
				Shrinker: Uint64(limits),
			}

			for arb.Shrinker != nil {
				var err error
				arb, err = arb.Shrinker(arb, false)
				if err != nil {
					t.Fatalf("Unexpected error : %s", err)
				}
			}

			if arb.Value.Uint() != limits.Max {
				t.Errorf("Shrunk value: %d should be exactly the value of upper bound limit: %d", arb.Value.Uint(), limits.Max)
			}
		},
	}

	for name, testCase := range testCase {
		t.Run(name, testCase)
	}
}
