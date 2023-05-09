package generator

import (
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Nil returns generator for types that can have nil values. Supported types
// are: chan, slice, map, func, interface and pointers. Error is returned if
// target is not one of the supported types.
func Nil() arbitrary.Generator {
	return func(target reflect.Type, bias constraints.Bias, _ arbitrary.Random) (arbitrary.Arbitrary, error) {
		switch target.Kind() {
		case reflect.Chan, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface, reflect.Ptr:
			return arbitrary.Arbitrary{
				Value: reflect.Zero(target),
			}, nil
		default:
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "Nil")
		}
	}
}
