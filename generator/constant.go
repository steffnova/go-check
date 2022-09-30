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
	return func(target reflect.Type, bias constraints.Bias, r Random) (arbitrary.Arbitrary, shrinker.Shrinker, error) {
		switch {
		case constant == nil:
			return Nil()(target, bias, r)
		case target.Kind() == reflect.TypeOf(constant).Kind():
			fallthrough
		case target.Kind() == reflect.Interface && reflect.TypeOf(constant).Implements(target):
			return arbitrary.Arbitrary{
				Value: reflect.ValueOf(constant),
			}, nil, nil
		default:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("constant %s doesn't match the target's type: %s", reflect.TypeOf(constant).Kind().String(), target.String())
		}
	}
}

// Constant returns generator that returns one of the constants. Error is returned if
// number of constants is 0 or chosen constant doesn't match generator's target
func ConstantFrom(constants ...interface{}) Generator {
	if len(constants) == 0 {
		return Invalid(fmt.Errorf("number of constants must be greater than 0"))
	}
	generators := make([]Generator, len(constants))
	for index, constant := range constants {
		generators[index] = Constant(constant)
	}

	return OneFrom(generators...)
}
