package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

func Array(shrinker Shrinker) Shrinker {
	if shrinker == nil {
		return nil
	}
	return func(val arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
		switch {
		case val.Value.Kind() != reflect.Array:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("array shrinker cannot shrink %s", val.Value.Kind().String())
		case val.Value.Len() != len(val.Elements):
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("number of elements must match size of the array")
		default:
			next, shrinker, err := shrinker(val, propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, nil, err
			}

			next.Value = reflect.New(val.Value.Type()).Elem()
			for index, element := range val.Elements {
				next.Value.Index(index).Set(element.Value)
			}

			return next, Array(shrinker), nil
		}
	}
}
