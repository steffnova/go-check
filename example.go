package check

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

type example func([]reflect.Type) (arbitrary.Arbitraries, error)

func Example(inputs ...any) example {
	return func(targets []reflect.Type) (arbitrary.Arbitraries, error) {
		for index, input := range inputs {
			inputType := reflect.TypeOf(input)
			if inputType != targets[index] {
				return nil, fmt.Errorf("Input at index [%d] - %s doesn't match predicate input at index[%d] - %s", index, inputType, index, targets[index])
			}
		}

		arbs := make(arbitrary.Arbitraries, len(inputs))
		for index := range inputs {
			arbs[index] = arbitrary.Arbitrary{
				Value: reflect.ValueOf(inputs[index]),
			}
		}
		return arbs, nil
	}
}

func (e example) Filter(predicate any) example {
	return func(t []reflect.Type) (arbitrary.Arbitraries, error) {
		arbs, err := e(t)
		if err != nil {
			return nil, err
		}

		out := reflect.ValueOf(predicate).Call(arbs.Values())
		if !out[0].Bool() {
			return nil, fmt.Errorf("example failed filter predicate")
		}

		return arbs, nil
	}
}
