package property

import (
	"testing"

	"github.com/steffnova/go-check/arbitrary"
)

func TestPredicate(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"DefintionNil": func(t *testing.T) {
			_, runner := Predicate(nil)()
			if err := runner(arbitrary.Arbitraries{}); err == nil {
				t.Error("Expected error")
			}
		},
		"DefinitonNotAFunction": func(t *testing.T) {
			_, runner := Predicate(0)()
			if err := runner(arbitrary.Arbitraries{}); err == nil {
				t.Error("Expected error")
			}
		},
		"DefinitionMoreThen1Output": func(t *testing.T) {
			_, runner := Predicate(func() (int, int) {
				return 0, 0
			})()
			if err := runner(arbitrary.Arbitraries{}); err == nil {
				t.Error("Expected error")
			}
		},
		"DefinitionHasLessThen1Output": func(t *testing.T) {
			_, runner := Predicate(func() {})()
			if err := runner(arbitrary.Arbitraries{}); err == nil {
				t.Error("Expected error")
			}
		},
		"DefinitionOutputIsNotError": func(t *testing.T) {
			_, runner := Predicate(func() int { return 0 })()
			if err := runner(arbitrary.Arbitraries{}); err == nil {
				t.Error("Expected error")
			}
		},
		"TooManyArbitrariesPassedToPredicateRunner": func(t *testing.T) {
			_, runner := Predicate(func() error {
				return nil
			})()

			if err := runner(make(arbitrary.Arbitraries, 5)); err == nil {
				t.Error("Expected error")
			}
		},
		"InsufficientArbitrariesPassedToPredicateRunner": func(t *testing.T) {
			_, runner := Predicate(func(int, int, int) error {
				return nil
			})()

			if err := runner(arbitrary.Arbitraries{}); err == nil {
				t.Error("Expected error")
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}

}
