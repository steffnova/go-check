package check

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/steffnova/go-check/generator"
	"github.com/steffnova/go-check/shrinker"
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

type run func() error

type property func(generator.Random) (run, error)

func Property(predicate interface{}, arbGenerators ...generator.Arbitrary) property {
	return func(r generator.Random) (run, error) {
		generators := make([]generator.Generator, len(arbGenerators))

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
				generate, err := arbGenerator(val.Type().In(index), r)
				if err != nil {
					return nil, fmt.Errorf("failed to create type generator at index [%d]. %s", index, err)
				}
				generators[index] = generate
			}
		}

		return func() error {
			inputs := make([]reflect.Value, len(generators))
			shrinkers := make([]shrinker.Shrinker, len(generators))

			for index, generate := range generators {
				inputs[index], shrinkers[index] = generate()
			}

			outputs := reflect.ValueOf(predicate).Call(inputs)
			if outputs[0].IsZero() {
				return nil
			}

			numberOfShrinks := 0
			for index, shrinker := range shrinkers {
				for shrinker != nil {
					oldValue := inputs[index]
					failed := !outputs[0].IsZero()
					inputs[index], shrinker = shrinker(failed)
					if failed && !reflect.DeepEqual(oldValue.Interface(), inputs[index].Interface()) {
						numberOfShrinks++
					}
					outputs = reflect.ValueOf(predicate).Call(inputs)
				}
			}
			// TODO: shrink on error
			return fmt.Errorf("%s. \nShrunked %d times. \nProperty error: %s", ErrorForInputs(inputs).Error(), numberOfShrinks, outputs[0].Interface().(error))

		}, nil
	}
}
