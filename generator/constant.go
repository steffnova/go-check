package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Constant returns generator that always generates the value passed to it via "constant"
// parameter. Error is returned if value passed to generator doesn't match generator's target.
func Constant(constant interface{}) arbitrary.Generator {
	if constant == nil {
		return Nil()
	}
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		switch {
		case target.Kind() == reflect.TypeOf(constant).Kind():
			fallthrough
		case target.Kind() == reflect.Interface && reflect.TypeOf(constant).Implements(target):
			return arbitrary.Arbitrary{
				Value: reflect.ValueOf(constant),
			}, nil
		default:
			return arbitrary.Arbitrary{}, fmt.Errorf("%w. Constant %s doesn't match the target's type: %s", arbitrary.ErrorInvalidTarget, reflect.TypeOf(constant).Kind().String(), target.String())
		}
	}
}

// Constant returns generator that returns one of the constants. Error is returned if
// number of constants is 0 or chosen constant doesn't match generator's target
func ConstantFrom(constant interface{}, constants ...interface{}) arbitrary.Generator {
	generators := make([]arbitrary.Generator, len(constants)+1)
	for index, constant := range append([]interface{}{constant}, constants...) {
		generators[index] = Constant(constant)
	}

	return OneFrom(generators[0], generators[1:]...)
}
