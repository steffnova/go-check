package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/shrinker"
)

// Generator generates random value and shrinker for it's type
type Generator func() (reflect.Value, shrinker.Shrinker)

// Arbitrary is Generator creator. It tries to create Generator for type specified
// by target parameter with provided Random instance as r parameter.
type Arbitrary func(target reflect.Type, r Random) (Generator, error)

// Map maps receiver Arbitrary (arb) to a new Arbitrary using mapper. Mapper must be a
// function that has one input and one output. Mappers's input type must satisfy
// target's type. Mapper's output defines the Generator type created by mapped
// Arbitrary.
func (arb Arbitrary) Map(mapper interface{}) Arbitrary {
	return func(target reflect.Type, r Random) (Generator, error) {
		val := reflect.ValueOf(mapper)
		switch {
		case val.Kind() != reflect.Func:
			return nil, fmt.Errorf("mapper must be a function")
		case val.Type().NumOut() != 1:
			return nil, fmt.Errorf("mapper must have 1 output value")
		case val.Type().NumIn() != 1:
			return nil, fmt.Errorf("mapper must have 1 input value")
		case val.Type().Out(0).Kind() != target.Kind():
			return nil, fmt.Errorf("mappers output parameter's kind must match target's kind. Got: %s", target.Kind())
		}

		generator, err := arb(val.Type().In(0), r)
		if err != nil {
			return nil, fmt.Errorf("failed to create base generator. %s", err)
		}

		return func() (reflect.Value, shrinker.Shrinker) {
			val, shrinker := generator()
			val = reflect.ValueOf(mapper).Call([]reflect.Value{val})[0]
			return val.Convert(target), shrinker.Map(mapper).Convert(target)
		}, nil
	}
}

// Filter creates new Arbitrary from receiver Arbitrary (arb) using predicate. Predicate
// is a function that has one input and one output. Input paramter must satisfy target's
// type, while output parameter must be bool. Generator returned by new Arbitrary will
// generate values that satisfy predicate.
//
// NOTE: This can highly impact Generator's time to generate a value as it will try to
// generate target's values unitl predicate is satisfied.
func (arb Arbitrary) Filter(predicate interface{}) Arbitrary {
	return func(target reflect.Type, r Random) (Generator, error) {
		generate, err := arb(target, r)
		switch val := reflect.ValueOf(predicate); {
		case err != nil:
			return nil, fmt.Errorf("failed to create base generator. %s", err)
		case val.Kind() != reflect.Func:
			return nil, fmt.Errorf("predicate must be a function")
		case val.Type().NumIn() != 1:
			return nil, fmt.Errorf("predicate must have one input value")
		case val.Type().NumOut() != 1:
			return nil, fmt.Errorf("predicate must have one output value")
		case val.Type().Out(0).Kind() != reflect.Bool:
			return nil, fmt.Errorf("predicate must have bool as a output value")
		}

		return func() (reflect.Value, shrinker.Shrinker) {
			for {
				val, _ := generate()
				outputs := reflect.ValueOf(predicate).Call([]reflect.Value{val})
				if outputs[0].Bool() {
					return val.Convert(target), nil
				}
			}
		}, nil
	}
}
