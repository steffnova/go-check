package generator

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

// Generator generates random arbitrary.Type using pseudo-random
// number generator specified by rand parameter.
type Generator func(rand *rand.Rand) arbitrary.Type

// Arbitrary is Generator creator. It tries to create Generator
// for type specified by target parameter.
type Arbitrary func(target reflect.Type) (Generator, error)

// Map maps receiver Arbitrary (arb) to a new Arbitrary using mapper. Mapper must be a
// function that has one input and one output. Mappers's input type must satisfy
// target's type. Mapper's output defines the Generator type created by mapped
// Arbitrary.
func (arb Arbitrary) Map(mapper interface{}) Arbitrary {
	return func(target reflect.Type) (Generator, error) {
		val := reflect.ValueOf(mapper)
		switch {
		case val.Kind() != reflect.Func:
			return nil, fmt.Errorf("mapper must be a function")
		case val.Type().NumOut() != 1:
			return nil, fmt.Errorf("mapper must have 1 output value")
		case val.Type().NumIn() != 1:
			return nil, fmt.Errorf("mapper must have 1 input value")
		}

		generateMappedValue, err := arb(val.Type().In(0))
		switch {
		case err != nil:
			return nil, fmt.Errorf("failed to create base generator. %s", err)
		case val.Type().Out(0).Kind() != target.Kind():
			return nil, fmt.Errorf("mappers output parameter's kind must match target's kind")
		default:
			return func(rand *rand.Rand) arbitrary.Type {
				arbType := generateMappedValue(rand)
				outputs := reflect.ValueOf(mapper).Call([]reflect.Value{arbType.Value()})
				return arbitrary.Mapped{
					Base:   arbType,
					Mapper: mapper,
					Val:    outputs[0],
				}
			}, nil
		}
	}
}

// Filter creates new Arbitrary from receiver Arbitrary (arb) using predicate. Predicate
// is a function that has one input and one output. Input paramter must satisfy target's
// type, while output parameter must be bool. Generator returned by new Arbitrary will
// generate values that satisfy predicate.
//
// NOTE: This can highly impact Generator's time to generate arbitrary.Type as it will
// try to generate target's values unitl predicate is satisfied.
func (arb Arbitrary) Filter(predicate interface{}) Arbitrary {
	return func(target reflect.Type) (Generator, error) {
		generate, err := arb(target)
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

		return func(rand *rand.Rand) arbitrary.Type {
			for {
				arbType := generate(rand)
				outputs := reflect.ValueOf(predicate).Call([]reflect.Value{arbType.Value()})
				if outputs[0].Bool() {
					return arbType
				}
			}
		}, nil
	}
}
