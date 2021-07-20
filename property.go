package check

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"

	"github.com/steffnova/go-check/generator"
)

type Error func() string

func (pe Error) Error() string {
	return pe()
}

func ErrorForInputs(inputs []reflect.Value) Error {
	return func() string {
		inputData := make([]string, len(inputs))
		for index, input := range inputs {
			inputData[index] = fmt.Sprintf("<%s> %#v", input.Type().String(), input.Interface())
		}

		return fmt.Sprintf("Property failed for inputs: [\n\t%s\n]", strings.Join(inputData, ",\n\t"))
	}
}

type run func(rand *rand.Rand) error

type property func() (run, error)

func Property(predicate interface{}, arbGenerators ...generator.Arbitrary) property {
	return func() (run, error) {
		generators := make([]generator.Type, len(arbGenerators))

		switch val := reflect.ValueOf(predicate); {
		case val.Kind() != reflect.Func:
			return nil, fmt.Errorf("predicate must be a function")
		case val.Type().NumIn() != len(arbGenerators):
			return nil, fmt.Errorf("number of predicate input parameters (%d) doesn't match number of generators (%d)", val.Type().NumIn(), len(generators))
		case val.Type().NumOut() != 1:
			return nil, fmt.Errorf("number of predicate output parameters must be 1")
		case !val.Type().Out(0).Implements(reflect.TypeOf((*error)(nil)).Elem()):
			return nil, fmt.Errorf("predicate's output parameter type must be error")
		default:
			for index, arbGenerator := range arbGenerators {
				generator, err := arbGenerator(val.Type().In(index))
				if err != nil {
					return nil, fmt.Errorf("failed to create type generator at index [%d]. %s", index, err)
				}
				if !generator.Type.ConvertibleTo(val.Type().In(index)) {
					return nil, fmt.Errorf("generator's arbitrary type (%s) can't be assigned to predicate's input type (%s)", generator.Type, val.Type().In(index))
				}
				generators[index] = generator
			}
		}

		return func(r *rand.Rand) error {
			inputs := make([]reflect.Value, len(generators))
			for index, generator := range generators {
				inputs[index] = generator.Generate(r).Value()
			}

			outputs := reflect.ValueOf(predicate).Call(inputs)
			if !outputs[0].IsZero() {
				// TODO: shrink on error
				return fmt.Errorf("%s. \nProperty error: %s", ErrorForInputs(inputs).Error(), outputs[0].Interface().(error))
			}

			return nil
		}, nil
	}
}
