package shrinker

import (
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
)

func TestCollectionSizeRemoveFront(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"IndexIsNegative": func(t *testing.T) {
			arb := arbitrary.Arbitrary{
				Elements: arbitrary.Arbitraries{
					{
						Value:    reflect.ValueOf(uint64(200)),
						Shrinker: nil,
					},
					{
						Value:    reflect.ValueOf(uint64(200)),
						Shrinker: nil,
					},
				},
			}
			arb.Shrinker = CollectionSizeRemoveFront(-1)

			_, err := arb.Shrinker(arb, true)
			if err == nil {
				t.Fatalf("Expected error")
			}
		},
		"RemoveFirstElement": func(t *testing.T) {
			elements := make([]arbitrary.Arbitrary, 10)
			for index := range elements {
				elements[index] = arbitrary.Arbitrary{
					Value:    reflect.ValueOf(uint64(index)),
					Shrinker: nil,
				}
			}
			arb := arbitrary.Arbitrary{
				Shrinker: CollectionSizeRemoveBack(0),
				Elements: elements,
			}

			shrink, err := arb.Shrinker(arb, true)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if len(shrink.Elements) != len(elements)-1 {
				t.Fatalf("Number of elements should be reduced by 1")
			}

			for index := range shrink.Elements {
				if !reflect.DeepEqual(shrink.Elements[index].Value, arb.Elements[index+1].Value) {
					t.Fatalf("Remaining elements in shrink do not match their counterpart in original arbitrary")
				}
			}
		},
		"Remove2ndElement": func(t *testing.T) {
			elements := make([]arbitrary.Arbitrary, 10)
			for index := range elements {
				elements[index] = arbitrary.Arbitrary{
					Value:    reflect.ValueOf(uint64(index)),
					Shrinker: nil,
				}
			}
			arb := arbitrary.Arbitrary{
				Shrinker: CollectionSizeRemoveFront(0),
				Elements: elements,
			}

			// First shrinker invocation removes the last element.
			// Second invocation reverts the change and removes the 2nd element from the back.
			shrink, err := arb.Shrinker(arb, true)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			shrink, err = shrink.Shrinker(arb, false)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			arb.Elements = append(arb.Elements[:1], arb.Elements[2:]...)
			for index := range arb.Elements {
				if !reflect.DeepEqual(shrink.Elements[index].Value, arb.Elements[index].Value) {
					t.Fatalf("Remaining elements in shrink do not match their counterpart in original arbitrary")
				}
			}
		},
		"ShrinkingFinishes": func(t *testing.T) {
			elements := make([]arbitrary.Arbitrary, 10)
			for index := range elements {
				elements[index] = arbitrary.Arbitrary{
					Value:    reflect.ValueOf(uint64(index)),
					Shrinker: nil,
				}
			}
			arb := arbitrary.Arbitrary{
				Shrinker: CollectionSizeRemoveBack(len(elements) - 1),
				Elements: elements,
			}

			predicate := func(in arbitrary.Arbitrary) bool {
				for _, element := range arb.Elements {
					if element.Value.Uint() == 3 {
						return true
					}
				}
				return false
			}

			// Shrinking process should remove all elements except
			// the one with the value 3
			for arb.Shrinker != nil {
				var err error
				arb, err = arb.Shrinker(arb, predicate(arb))
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}

			if len(arb.Elements) != 1 {
				t.Fatalf("Expected only one element")
			}
			if arb.Elements[0].Value.Uint() != uint64(3) {
				t.Fatalf("Expected element value to be 3, got %d", arb.Elements[0].Value.Uint())
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
