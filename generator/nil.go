package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Nil is Arbitrary that creates nil Generator. Generates nil value for
// chan, slice, map, func, interface and ptr types. Error is returned
// if target is not one of the supported types.
func Nil() Generator {
	return func(target reflect.Type, bias constraints.Bias, _ Random) (Generate, error) {
		switch target.Kind() {
		case reflect.Chan, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface, reflect.Ptr:
			return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
				return arbitrary.Arbitrary{
					Value: reflect.Zero(target),
				}, nil
			}, nil
		default:
			return nil, fmt.Errorf("nil is not a valid value for target: %s", target.String())
		}
	}
}
