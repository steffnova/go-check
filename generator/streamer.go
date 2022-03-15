package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

type streamer func(Random) error

// Streamer returns new streamer that generates data using "generators" parameter and
// streams it to it's target. Target must be a function that has no output values.
// Number of input parameters can be arbitrary. Error is returned if "target" is not
// a function, number of generators doesn't match target's number of input parameters,
// target doesn't have 0 output values or if any of generators returns an error.
func Streamer(target interface{}, generators ...Generator) streamer {
	return func(r Random) error {
		targetVal := reflect.ValueOf(target)
		switch {
		case targetVal.Kind() != reflect.Func:
			return fmt.Errorf("stream's target must be a function")
		case targetVal.Type().NumIn() != len(generators):
			return fmt.Errorf("number of generators %d must match number of input parameters :%d of stream target", len(generators), targetVal.Type().NumIn())
		case targetVal.Type().NumOut() != 0:
			return fmt.Errorf("number of stream target's outputs must be 0")
		}

		arbs := make(arbitrary.Arbitraries, len(generators))
		for index, generator := range generators {
			generate, err := generator(targetVal.Type().In(index), constraints.Bias{Size: 100, Scaling: 1}, r)
			if err != nil {
				return fmt.Errorf("failed to generate input parameter: %d. %s", index, err)
			}
			arbs[index], _ = generate()
		}

		targetVal.Call(arbs.Values())
		return nil
	}
}
