package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Nil is Arbitrary that creates nil Generator. Generates nil value for
// chan, slice, map, func, interface and ptr types. Error is returned
// if target is not one of the supported types.
func Nil() Arbitrary {
	return func(target reflect.Type, bias constraints.Bias, _ Random) (Generator, error) {
		switch target.Kind() {
		case reflect.Chan, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface, reflect.Ptr:
			return func() (reflect.Value, shrinker.Shrinker) {
				return reflect.Zero(target), nil
			}, nil
		default:
			return nil, fmt.Errorf("nil is not a valid value for target: %s", target.String())
		}
	}
}
