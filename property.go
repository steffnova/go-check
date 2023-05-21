package check

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

type property func(constraints.Bias, arbitrary.Random) (Details, error)

type Details struct {
	NumberOfExamples   int
	NumberOfIterations int
	NumberOfShrinks    uint
	FailureReason      error
	FailureInput       arbitrary.Arbitraries
}

// Property defines a new property by specifing predicate and property generators.
// Predicate must be a function that can have any number of input values, and must
// have only one output value of error type. Number of predicate's input parameters
// must match number of generators.
func Property(predicate interface{}, generators ...arbitrary.Generator) property {
	return func(bias constraints.Bias, r arbitrary.Random) (Details, error) {
		predicateVal := reflect.ValueOf(predicate)

		switch t := predicateVal.Type(); {
		case t.Kind() != reflect.Func:
			return Details{}, fmt.Errorf("predicate must be a function")
		case t.NumIn() != len(generators):
			return Details{}, fmt.Errorf("number of predicate input parameters (%d) doesn't match number of generators (%d)", t.NumIn(), len(generators))
		case t.NumOut() != 1:
			return Details{}, fmt.Errorf("number of predicate output parameters must be 1")
		case !t.Out(0).Implements(reflect.TypeOf((*error)(nil)).Elem()):
			return Details{}, fmt.Errorf("predicate's output parameter type must be error")
		}

		inputs := make(arbitrary.Arbitraries, len(generators))
		for index, generator := range generators {
			arb, err := generator(predicateVal.Type().In(index), bias, r)
			if err != nil {
				return Details{}, fmt.Errorf("failed to use generator for property parameter at index %d. %s", index+1, err)
			}
			inputs[index] = arb
		}

		outputs := predicateVal.Call(inputs.Values())
		if outputs[0].IsZero() {
			return Details{}, nil
		}

		numberOfShrinks := uint(0)
		for index := range inputs {
			for inputs[index].Shrinker != nil {
				propertyFailed := !outputs[0].IsZero()
				input, err := inputs[index].Shrinker(inputs[index], propertyFailed)
				if err != nil {
					return Details{}, fmt.Errorf("failed to shrink input with index: %d. %s", index, err)
				}

				if propertyFailed && !reflect.DeepEqual(inputs[index].Value.Interface(), input.Value.Interface()) {
					numberOfShrinks++
				}
				inputs[index] = input
				outputs = predicateVal.Call(inputs.Values())
			}
		}

		return Details{
			NumberOfShrinks: uint(numberOfShrinks),
			FailureReason:   outputs[0].Interface().(error),
			FailureInput:    inputs,
		}, nil
	}
}

type property2 func(r arbitrary.Random, iterations int) (Details, error)

func Property2(in inputsGenerator, p predicate, examples ...example) property2 {
	return func(r arbitrary.Random, iterations int) (Details, error) {
		targets, runner := p()
		for index, example := range examples {
			inputs, err := example(targets)
			if err != nil {
				return Details{}, fmt.Errorf("%d. example is invalid: %w", index+1, err)
			}
			err = runner(inputs)
			if err != nil {
				return Details{
					NumberOfExamples: index,
					FailureReason:    err,
					FailureInput:     inputs,
				}, nil
			}
		}

		for i := 0; i < iterations; i++ {
			bias := constraints.Bias{
				Size:    iterations,
				Scaling: iterations - i,
			}

			arbs, shrinker, err := in(targets, bias, r)
			if err != nil {
				return Details{}, err
			}

			predicateErr := runner(arbs)
			if predicateErr == nil {
				continue
			}

			numberOfShrinks := uint(0)

			for shrinker != nil {
				var shrinkingErr error
				propertyFailed := predicateErr != nil
				arbs, shrinker, shrinkingErr = shrinker(arbs, propertyFailed)
				if shrinkingErr != nil {
					return Details{}, err
				}
				predicateErr = runner(arbs)
			}

			return Details{
				FailureInput:       arbs,
				FailureReason:      predicateErr,
				NumberOfShrinks:    uint(numberOfShrinks),
				NumberOfExamples:   len(examples),
				NumberOfIterations: i,
			}, nil
		}

		return Details{}, nil
	}
}
