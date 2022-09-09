package shrinker

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/steffnova/go-check/arbitrary"
)

// Struct is a shrinker for struct. Struct shrinking consists of shrinking individual
// fields one by one. Error is returned if struct type is not struct, length of fieldShrinks
// is not equal to number of struct fields or if any of the struct fields return an error
// during shrinking.
func Struct(shrinker Shrinker) Shrinker {
	if shrinker == nil {
		return nil
	}
	return func(val arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
		switch {
		case val.Value.Kind() != reflect.Struct:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("array shrinker cannot shrink %s", val.Value.Kind().String())
		case val.Value.NumField() != len(val.Elements):
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("number of elements must match size of the array")
		default:
			next, shrinker, err := shrinker(val, propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, nil, err
			}

			next.Value = reflect.New(val.Value.Type()).Elem()
			for index, element := range val.Elements {
				reflect.NewAt(
					next.Value.Field(index).Type(),
					unsafe.Pointer(next.Value.Field(index).UnsafeAddr()),
				).Elem().Set(element.Value)
			}

			return next, Struct(shrinker), nil
		}
	}
}
