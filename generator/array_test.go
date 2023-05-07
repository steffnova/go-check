package generator

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/constraints"
)

func TestArrayFrom(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				ArrayFrom(Int()),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s':", ErrorInvalidTarget)
			}
		},
		"InvalidNumberOfElementarbitrary.Generators": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(n [3]uint) {},
				ArrayFrom(Uint(), Uint()),
			))

			if !errors.Is(err, ErrorInvalidCollectionSize) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidCollectionSize)
			}
		},
		"InvalidElementarbitrary.GeneratorType": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(n [3]uint) {},
				ArrayFrom(Int(), Int(), Int()),
			))

			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s':", ErrorInvalidTarget)
			}
		},
		"arbitrary.GeneratorPerElement": func(t *testing.T) {
			// This test case aims to confirm whether each element within an array is
			//  utilizing its designated generator. In order to validate this behavior,
			// each element must abide by the constraints defined by its respective generator.
			arrayConstraints := []constraints.Uint{
				{Min: 0, Max: 100},
				{Min: 1000, Max: 10000},
				{Min: 100000, Max: 10000000},
			}

			err := Stream(0, 100, Streamer(
				func(n [3]uint) {
					for index := range n {
						if n[index] < arrayConstraints[index].Min || n[index] > arrayConstraints[index].Max {
							t.Fatalf("Element[%d] is not within constraints %#v", index, arrayConstraints[index])
						}
					}
				},
				ArrayFrom(
					Uint(arrayConstraints[0]),
					Uint(arrayConstraints[1]),
					Uint(arrayConstraints[2]),
				),
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
