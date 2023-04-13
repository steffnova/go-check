package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Constant returns generator that always generates the value passed to it via "constant"
// parameter. Error is returned if value passed to generator doesn't match generator's target.
func Constant(constant interface{}) Generator {
	if constant == nil {
		return Nil()
	}
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		switch {
		case target.Kind() == reflect.TypeOf(constant).Kind():
			fallthrough
		case target.Kind() == reflect.Interface && reflect.TypeOf(constant).Implements(target):
			return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
				return arbitrary.Arbitrary{
					Value: reflect.ValueOf(constant),
				}, nil
			}, nil
		default:
			return nil, fmt.Errorf("%w. Constant %s doesn't match the target's type: %s", ErrorInvalidTarget, reflect.TypeOf(constant).Kind().String(), target.String())
		}
	}
}

// Constant returns generator that returns one of the constants. Error is returned if
// number of constants is 0 or chosen constant doesn't match generator's target
func ConstantFrom(constant interface{}, constants ...interface{}) Generator {
	generators := make([]Generator, len(constants)+1)
	for index, constant := range append([]interface{}{constant}, constants...) {
		generators[index] = Constant(constant)
	}

	return OneFrom(generators[0], generators[1:]...)
}
