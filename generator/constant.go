package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

// Constant is arbitrary that creates generator of constant value. Value returned
// by generator is the one passed in constant parameter. Unlike other generators
// Constant can be used to satisfy target of interface kind. Nil value can be passed
// as a constant parameter, if it is targets's zero value (Chan, Slice, Ptr, Interface,
// Func and Map). Error is returned if constant is nil and nil is not target's zero
// value, target is an interface but constant doesn't implement it, and finally if
// target's kinds doesn't match constant's kind.
func Constant(constant interface{}) Arbitrary {
	return func(target reflect.Type, r Random) (Generator, error) {
		switch {
		case constant == nil:
			return Nil()(target, r)
		case target.Kind() == reflect.TypeOf(constant).Kind():
			fallthrough
		case target.Kind() == reflect.Interface && reflect.TypeOf(constant).Implements(target):
			return func() arbitrary.Type {
				return arbitrary.Constant{
					C: reflect.ValueOf(constant),
				}
			}, nil
		default:
			return nil, fmt.Errorf("constant %s doesn't match the target's type: %s", reflect.TypeOf(constant).Kind().String(), target.String())
		}
	}

}

// ConstantFrom is arbitrary that creates generator of constant value from
// one of the passed values. It requires at least one constant which is
// defined by first parameter. Additional constants can be passed in variadic
// parameter other. Created generator has the same rules applies as the one
// returned from Constant arbitrary.
func ConstantFrom(first interface{}, other ...interface{}) Arbitrary {
	constants := make([]Arbitrary, len(other))
	for index, constant := range other {
		constants[index] = Constant(constant)
	}

	return OneFrom(Constant(first), constants...)
}
