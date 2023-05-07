package arbitrary

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

// Generator returns Generate for a type specified by "target" parameter, that can be used to
// generate target's value using "bias" and "r".
type Generator func(target reflect.Type, bias constraints.Bias, r Random) (Arbitrary, error)

// Map (combinator) returns generator that maps generated value to a new one using mapper. Mapper
// must be a function that has one input and one output. Mapper's input type must match generated
// value's type and Mapper's output type must match generator's target type. Error is returned if
// mapper is invalid or if generator of mapper's input type returns an error.
func (generator Generator) Map(mapper interface{}) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (Arbitrary, error) {
		val := reflect.ValueOf(mapper)
		switch {
		case val.Kind() != reflect.Func:
			return Arbitrary{}, fmt.Errorf("%w. Mapper must be a function", ErrorMapper)
		case val.Type().NumOut() != 1:
			return Arbitrary{}, fmt.Errorf("%w. Mapper must have 1 output value", ErrorMapper)
		case val.Type().NumIn() != 1:
			return Arbitrary{}, fmt.Errorf("%w. Mapper must have 1 input value", ErrorMapper)
		case val.Type().Out(0).Kind() != target.Kind():
			return Arbitrary{}, fmt.Errorf("%w. Mappers output kind: %s must match target's kind. Got: %s", ErrorMapper, val.Type().Out(0).Kind(), target.Kind())
		}

		arb, err := generator(val.Type().In(0), bias, r)
		if err != nil {
			return Arbitrary{}, fmt.Errorf("Failed to use base generator. %w", err)
		}

		return Arbitrary{
			Value:      reflect.ValueOf(mapper).Call([]reflect.Value{arb.Value})[0],
			Precursors: []Arbitrary{arb},
			Shrinker:   arb.Shrinker.Map(mapper),
		}, nil
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
	return func(target reflect.Type, bias constraints.Bias, r Random) (Arbitrary, error) {

		switch val := reflect.ValueOf(predicate); {

		case val.Kind() != reflect.Func:
			return Arbitrary{}, fmt.Errorf("%w. Filter predicate must be a function", ErrorFilter)
		case val.Type().NumIn() != 1:
			return Arbitrary{}, fmt.Errorf("%w. Filter predicate must have one input value", ErrorFilter)
		case val.Type().NumOut() != 1:
			return Arbitrary{}, fmt.Errorf("%w. Filter predicate must have one output value", ErrorFilter)
		case val.Type().Out(0).Kind() != reflect.Bool:
			return Arbitrary{}, fmt.Errorf("%w. Filter predicate must have bool as a output type", ErrorFilter)
		}

		for {
			arb, err := generator(target, bias, r)
			if err != nil {
				return Arbitrary{}, fmt.Errorf("Failed to use base generator. %w", err)
			}

			outputs := reflect.ValueOf(predicate).Call([]reflect.Value{arb.Value})
			if outputs[0].Bool() {
				arb.Shrinker = arb.Shrinker.Filter(predicate)
				return arb, nil
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
	return func(target reflect.Type, bias constraints.Bias, r Random) (Arbitrary, error) {
		binderVal := reflect.ValueOf(binder)
		switch t := reflect.TypeOf(binder); {
		case t.Kind() != reflect.Func:
			return Arbitrary{}, fmt.Errorf("%w. Binder must be a function", ErrorBinder)
		case t.NumIn() != 1:
			return Arbitrary{}, fmt.Errorf("%w. Binder must have one input value", ErrorBinder)
		case t.NumOut() != 1:
			return Arbitrary{}, fmt.Errorf("%w. Binder must have one output values", ErrorBinder)
		case t.Out(0) != reflect.TypeOf(Generator(nil)):
			return Arbitrary{}, fmt.Errorf("%w. Binder's output type must be generator.Generator", ErrorBinder)
		}

		sourceArb, err := generator(binderVal.Type().In(0), bias, r)
		if err != nil {
			return Arbitrary{}, fmt.Errorf("Failed to use base generator: %w", err)
		}

		binder := func(source Arbitrary) (Arbitrary, error) {
			generator := binderVal.Call([]reflect.Value{source.Value})[0].Interface().(Generator)
			boundArb, err := generator(target, bias, r)
			if err != nil {
				return Arbitrary{}, fmt.Errorf("Generator Binding failed (%s -> %s): %w", binderVal.Type().In(0), target, err)
			}

			boundArb.Precursors = append(boundArb.Precursors, source)
			return boundArb, nil
		}

		// sourceArb, sourceShrinker := generate()
		boundVal, err := binder(sourceArb)
		if err != nil {
			return Arbitrary{}, err
		}

		boundVal.Shrinker = sourceArb.Shrinker.Bind(binder, boundVal)
		return boundVal, nil
	}
}
