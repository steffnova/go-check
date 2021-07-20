package generator

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

// Arbitrary is a function signature for Arbitrary generator. Invoking it creates
// Type generator in case of success or returns an error if creation fails. Target
// parameter should be used to verify the type of arbitrary, generator should generate.
type Arbitrary func(target reflect.Type) (Type, error)

// Map creates new Arbitrary, where generator's type is defined by mapper.
//
// Mapper must be a function that has one input and one output. It's input type must
// match type of generator returned by receiver Arbitrary. Output type defines the new
// type of created Type generator.
func (arbGenerator Arbitrary) Map(mapper interface{}) Arbitrary {
	return func(target reflect.Type) (Type, error) {
		val := reflect.ValueOf(mapper)
		switch {
		case val.Kind() != reflect.Func:
			return Type{}, fmt.Errorf("mapper must be a function")
		case val.Type().NumOut() != 1:
			return Type{}, fmt.Errorf("mapper must have 1 output value")
		case val.Type().NumIn() != 1:
			return Type{}, fmt.Errorf("mapper must have 1 input value")
		}

		generator, err := arbGenerator(val.Type().In(0))
		switch {
		case err != nil:
			return Type{}, fmt.Errorf("failed to create base generator. %s", err)
		case val.Type().In(0) != generator.Type:
			return Type{}, fmt.Errorf("mappers input parameter's type must match arbitrary's type")
		case val.Type().Out(0).Kind() != target.Kind():
			return Type{}, fmt.Errorf("mappers output parameter's kind must match target's kind")
		default:
			return Type{
				Type: reflect.ValueOf(mapper).Type().Out(0),
				Generate: func(rand *rand.Rand) arbitrary.Type {
					arbType := generator.Generate(rand)
					outputs := reflect.ValueOf(mapper).Call([]reflect.Value{arbType.Value()})
					return arbitrary.Mapped{
						Base:   arbType,
						Mapper: mapper,
						Val:    outputs[0],
					}
				},
			}, nil
		}
	}
}

func (arbGenerator Arbitrary) Filter(predicate interface{}) Arbitrary {
	return func(target reflect.Type) (Type, error) {
		generator, err := arbGenerator(target)
		switch val := reflect.ValueOf(predicate); {
		case err != nil:
			return Type{}, fmt.Errorf("failed to create base generator. %s", err)
		case val.Kind() != reflect.Func:
			return Type{}, fmt.Errorf("predicate must be a function")
		case val.Type().NumIn() != 1:
			return Type{}, fmt.Errorf("predicate must have one input value")
		case val.Type().In(0) != generator.Type:
			return Type{}, fmt.Errorf("predicate's input type must match generator's type")
		case val.Type().NumOut() != 1:
			return Type{}, fmt.Errorf("predicate must have one output value")
		case val.Type().Out(0).Kind() != reflect.Bool:
			return Type{}, fmt.Errorf("predicate must have bool as a output value")
		}

		return Type{
			Type: generator.Type,
			Generate: func(rand *rand.Rand) arbitrary.Type {
				for {
					arbType := generator.Generate(rand)
					outputs := reflect.ValueOf(predicate).Call([]reflect.Value{arbType.Value()})
					if outputs[0].Bool() {
						return arbType
					}
				}
			},
		}, nil
	}
}
