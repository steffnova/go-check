package property

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

type runner func(arbitrary.Arbitraries) error

type predicate func() ([]reflect.Type, runner)

// Predicate creates new predicate used by [Property] (see [Define]). The definition parameter
// must be a function, that can have arbitrary number of input parameters and a single output
// parameter of error type.
func Predicate(definition any) predicate {
	return func() ([]reflect.Type, runner) {
		predicateVal := reflect.ValueOf(definition)
		switch {
		case definition == nil:
			return nil, func(arbitrary.Arbitraries) error {
				return fmt.Errorf("definition is nil")
			}
		case predicateVal.Type().Kind() != reflect.Func:
			return nil, func(arbitrary.Arbitraries) error {
				return fmt.Errorf("  predicate must be a function")
			}
		case predicateVal.Type().NumOut() != 1:
			return nil, func(arbitrary.Arbitraries) error {
				return fmt.Errorf("number of predicate output parameters must be 1")
			}
		case !predicateVal.Type().Out(0).Implements(reflect.TypeOf((*error)(nil)).Elem()):
			return nil, func(arbitrary.Arbitraries) error {
				return fmt.Errorf("predicate's output parameter type must be error")
			}
		}

		targets := make([]reflect.Type, predicateVal.Type().NumIn())
		for index := 0; index < predicateVal.Type().NumIn(); index++ {
			targets[index] = predicateVal.Type().In(index)
		}

		return targets, func(arbs arbitrary.Arbitraries) error {
			if predicateVal.Type().NumIn() != len(arbs) {
				return fmt.Errorf("number of predicate input parameters (%d) doesn't match number of arbs (%d)", predicateVal.Type().NumIn(), len(arbs))
			}

			output := predicateVal.Call(arbs.Values())
			if !output[0].IsZero() {
				return output[0].Interface().(error)
			}
			return nil
		}

	}
}
