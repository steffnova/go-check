package generator

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
)

func TestStruct(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"InvalidTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Struct(),
			))

			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidTarget)
			}
		},
		"InvalidFieldName": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(struct{}) {},
				Struct(map[string]arbitrary.Generator{"X": Int()}),
			))

			if !errors.Is(err, ErrorInvalidConfig) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConfig)
			}
		},
		"InvalidFieldTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(struct{ X int }) {},
				Struct(map[string]arbitrary.Generator{"X": Uint()}),
			))

			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidTarget)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type testStruct struct{ X int }
			err := Stream(0, 100, Streamer(
				func(testStruct) {},
				Struct(),
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
