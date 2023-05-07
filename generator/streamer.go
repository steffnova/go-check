package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

type streamer func(arbitrary.Random) error

// Streamer returns new streamer that generates data using "generators" parameter and
// streams it to it's target. Target must be a function that has no output values.
// Number of input parameters can be arbitrary. Error is returned if "target" is not
// a function, number of generators doesn't match target's number of input parameters,
// target doesn't have 0 output values or if any of generators returns an error.
func Streamer(target interface{}, generators ...arbitrary.Generator) streamer {
	return func(r arbitrary.Random) error {
		targetVal := reflect.ValueOf(target)
		switch {
		case targetVal.Kind() != reflect.Func:
			return fmt.Errorf("%w. Stream's target must be a function", ErrorStream)
		case targetVal.Type().NumIn() != len(generators):
			return fmt.Errorf("%w. Number of generators %d must match number of input parameters :%d of stream target", ErrorStream, len(generators), targetVal.Type().NumIn())
		case targetVal.Type().NumOut() != 0:
			return fmt.Errorf("%w. Number of stream target's outputs must be 0", ErrorStream)
		}

		arbs := make(arbitrary.Arbitraries, len(generators))
		for index, generator := range generators {
			arb, err := generator(targetVal.Type().In(index), constraints.Bias{Size: 100, Scaling: 1}, r)
			if err != nil {
				return fmt.Errorf("failed to generate input parameter: %d. %w", index, err)
			}
			arbs[index] = arb
		}

		targetVal.Call(arbs.Values())
		return nil
	}
}
