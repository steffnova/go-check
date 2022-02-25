package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

func Slice(shrinker Shrinker) Shrinker {
	if shrinker == nil {
		return nil
	}
	return func(val arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
		switch {
		case val.Value.Kind() != reflect.Slice:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("slice shrinker cannot shrink %s", val.Value.Kind().String())
		case val.Value.Len() != len(val.Elements):
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("number of elements %d must match size of the array %d", len(val.Elements), val.Value.Len())
		default:
			next, shrinker, err := shrinker(val, propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, nil, err
			}

			next.Value = reflect.MakeSlice(val.Value.Type(), len(next.Elements), len(next.Elements))
			for index, element := range next.Elements {
				next.Value.Index(index).Set(element.Value)
			}

			return next, Slice(shrinker), nil
		}
	}
}
