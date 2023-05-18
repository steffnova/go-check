package shrinker

import (
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestPtr(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"OriginalNotAPtr": func(t *testing.T) {
			val := uint64(0)
			arb := arbitrary.Arbitrary{Value: reflect.ValueOf(val)}
			shrinker := Ptr(arb, constraints.Ptr{NilFrequency: 5})
			if _, err := shrinker(arb, true); err == nil {
				t.Fatalf("Expected error when original arbitrary is not pointer")
			}
		},
		"ShrinkingNotNil": func(t *testing.T) {
			val := uint64(1000)
			valPtr := &val

			arb := arbitrary.Arbitrary{
				Value: reflect.ValueOf(valPtr),
				Elements: arbitrary.Arbitraries{
					arbitrary.Arbitrary{
						Value:    reflect.ValueOf(val),
						Shrinker: Uint64(constraints.Uint64Default()),
					},
				},
			}
			arb.Shrinker = Ptr(arb, constraints.Ptr{NilFrequency: 0})

			for arb.Shrinker != nil {
				var err error
				arb, err = arb.Shrinker(arb, true)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}

			shrink := arb.Value.Interface().(*uint64)

			if shrink == nil || *shrink != 0 {
				t.Fatalf("Shrinking failed")
			}
		},
		"ShrinkingNil": func(t *testing.T) {
			val := uint64(1000)
			valPtr := &val

			arb := arbitrary.Arbitrary{
				Value: reflect.ValueOf(valPtr),
				Elements: arbitrary.Arbitraries{
					arbitrary.Arbitrary{
						Value:    reflect.ValueOf(val),
						Shrinker: Uint64(constraints.Uint64Default()),
					},
				},
			}
			arb.Shrinker = Ptr(arb, constraints.Ptr{NilFrequency: 5})

			for arb.Shrinker != nil {
				var err error
				arb, err = arb.Shrinker(arb, true)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}

			if !arb.Value.IsZero() {
				t.Fatalf("Shrinking failed")
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
