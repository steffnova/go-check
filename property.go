package check

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
	"github.com/steffnova/go-check/shrinker"
)

type property func(constraints.Bias, generator.Random) error

// Property defines a new property by specifing predicate and property generators.
// Predicate must be a function that can have any number of input values, and must
// have only one output value of error type. Number of predicate's input parameters
// must match number of generators.
func Property(predicate interface{}, arbGenerators ...generator.Arbitrary) property {
	return func(bias constraints.Bias, r generator.Random) error {
		generators := make([]generator.Generator, len(arbGenerators))
		predicateVal := reflect.ValueOf(predicate)

		switch t := reflect.TypeOf(predicate); {
		case t.Kind() != reflect.Func:
			return fmt.Errorf("predicate must be a function")
		case t.NumIn() != len(arbGenerators):
			return fmt.Errorf("number of predicate input parameters (%d) doesn't match number of generators (%d)", t.NumIn(), len(generators))
		case t.NumOut() != 1:
			return fmt.Errorf("number of predicate output parameters must be 1")
		case !t.Out(0).Implements(reflect.TypeOf((*error)(nil)).Elem()):
			return fmt.Errorf("predicate's output parameter type must be error")
		default:
			for index, arbGenerator := range arbGenerators {
				generate, err := arbGenerator(t.In(index), bias, r)
				if err != nil {
					return fmt.Errorf("failed to create type generator at index [%d]. %s", index, err)
				}
				generators[index] = generate
			}
		}

		inputs := make([]reflect.Value, len(generators))
		shrinkers := make([]shrinker.Shrinker, len(generators))

		for index, generate := range generators {
			inputs[index], shrinkers[index] = generate()
		}

		outputs := predicateVal.Call(inputs)
		if outputs[0].IsZero() {
			return nil
		}

		numberOfShrinks := 0
		for index, shrinker := range shrinkers {
			for shrinker != nil {
				propertyFailed := !outputs[0].IsZero()
				shrink, nextShrinker, err := shrinker(propertyFailed)
				if err != nil {
					return fmt.Errorf("failed to shrink input with index: %d. %s", index, err)
				}

				if propertyFailed && !reflect.DeepEqual(inputs[index].Interface(), shrink.Interface()) {
					numberOfShrinks++
				}
				inputs[index], shrinker = shrink, nextShrinker
				outputs = predicateVal.Call(inputs)
			}
		}

		return fmt.Errorf(strings.Join([]string{
			propertyFailed(inputs).Error(),
			fmt.Sprintf("Shrink %d time(s)", numberOfShrinks),
			fmt.Sprintf("Failure reason: %s", outputs[0].Interface().(error)),
		}, "\n"))
	}
}
