package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Generate returns new arbitrary and it's coresponding shrinker
// type Generate func()

// Generator returns Generate for a type specified by "target" parameter, that can be used to
// generate target's value using "bias" and "r".
type Generator func(target reflect.Type, bias constraints.Bias, r Random) (arbitrary.Arbitrary, shrinker.Shrinker, error)

// Map (combinator) returns generator that maps generated value to a new one using mapper. Mapper
// must be a function that has one input and one output. Mapper's input type must match generated
// value's type and Mapper's output type must match generator's target type. Error is returned if
// mapper is invalid or if generator of mapper's input type returns an error.
func (generator Generator) Map(mapper interface{}) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (arbitrary.Arbitrary, shrinker.Shrinker, error) {
		val := reflect.ValueOf(mapper)
		switch {
		case val.Kind() != reflect.Func:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("mapper must be a function")
		case val.Type().NumOut() != 1:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("mapper must have 1 output value")
		case val.Type().NumIn() != 1:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("mapper must have 1 input value")
		case val.Type().Out(0).Kind() != target.Kind():
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("mappers output kind: %s must match target's kind. Got: %s", val.Type().Out(0).Kind(), target.Kind())
		}

		arb, shrinker, err := generator(val.Type().In(0), bias, r)
		if err != nil {
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("failed to create base value. %s", err)
		}

		return arbitrary.Arbitrary{
			Value:      reflect.ValueOf(mapper).Call([]reflect.Value{arb.Value})[0],
			Precursors: []arbitrary.Arbitrary{arb},
		}, shrinker.Map(mapper), nil
	}
}

// Filter (combinator) returns a generator that generates value only if predicate is satisfied.
// Predicate is a function that has one input and one output. Predicate's input parameter must
// match generator's target type, and output parameter must be bool. Error is returned
// predicate is invalid, or generator for predicate's input returns an error.
//
// NOTE: Returned generator will retry generation until predicate is satisfied, which
// can impact the generator's speed.
func (generator Generator) Filter(predicate interface{}) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (arbitrary.Arbitrary, shrinker.Shrinker, error) {
		switch val := reflect.ValueOf(predicate); {
		case val.Kind() != reflect.Func:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("predicate must be a function")
		case val.Type().NumIn() != 1:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("predicate must have one input value")
		case val.Type().NumOut() != 1:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("predicate must have one output value")
		case val.Type().Out(0).Kind() != reflect.Bool:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("predicate must have bool as a output type")
		default:
			for {
				arb, shrinker, err := generator(target, bias, r)
				if err != nil {
					return arbitrary.Arbitrary{}, nil, fmt.Errorf("failed to create base value. %s", err)
				}

				outputs := reflect.ValueOf(predicate).Call([]reflect.Value{arb.Value})
				if outputs[0].Bool() {
					return arb, shrinker.Filter(arb, predicate), nil
				}

				if bias.Scaling != 1 {
					bias.Scaling--
				}
			}
		}
	}
}

// Bind (combinator) returns bounded generator using "binder" parameter. Binder
// is a function that has one input and one output. Input's type must match
// generated value's type. Output type must be generator.Generator. Binder allows
// using generated value of one generator as an input to another generator. Error
// is returned if binder is invalid, generator returns an error or bound generator
// returns an error.
func (generator Generator) Bind(binder interface{}) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (arbitrary.Arbitrary, shrinker.Shrinker, error) {
		binderVal := reflect.ValueOf(binder)
		switch t := reflect.TypeOf(binder); {
		case t.Kind() != reflect.Func:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("binder must be a function")
		case t.NumIn() != 1:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("binder must have one input value")
		case t.NumOut() != 1:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("binder must have one output values")
		case t.Out(0) != reflect.TypeOf(Generator(nil)):
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("binder's output type must be generator.Generator")
		}

		sourceArb, sourceShrinker, err := generator(binderVal.Type().In(0), bias, r)
		if err != nil {
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("failed to create base value: %s", err)
		}

		boundGenerator := binderVal.Call([]reflect.Value{sourceArb.Value})[0].Interface().(Generator)
		boundVal, boundShrinker, err := boundGenerator(target, bias, r)
		if err != nil {
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("generator composition failed: %s", err)
		}

		binder := func(source arbitrary.Arbitrary) (arbitrary.Arbitrary, shrinker.Shrinker, error) {
			generator := binderVal.Call([]reflect.Value{source.Value})[0].Interface().(Generator)

			val, shrinker, err := generator(target, bias, r)
			if err != nil {
				return arbitrary.Arbitrary{}, nil, fmt.Errorf("generator binding failed: %s", err)
			}

			return val, shrinker, nil
		}

		boundVal.Precursors = []arbitrary.Arbitrary{sourceArb}

		return boundVal, sourceShrinker.
			Retry(100, 100, sourceArb).
			Bind(binder, boundShrinker, boundShrinker), nil
	}
}

type logger func(interface{}, reflect.Type, constraints.Bias)

func (generator Generator) Log(logger logger) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (arbitrary.Arbitrary, shrinker.Shrinker, error) {
		arb, shrinker, err := generator(target, bias, r)

		logger(arb.Value.Interface(), target, bias)

		return arb, shrinker, err
	}
}
