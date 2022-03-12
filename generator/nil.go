package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Nil returns generator for types that can have nil values. Supported types
// are: chan, slice, map, func, interface and pointers. Error is returned if
// target is not one of the supported types.
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
			return nil, fmt.Errorf("can't use Nil generator for %s type", target)
		}
	}
}
