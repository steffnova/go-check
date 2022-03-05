package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

func Map(shrinker Shrinker) Shrinker {
	if shrinker == nil {
		return nil
	}
	return func(val arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
		switch {
		case val.Value.Kind() != reflect.Map:
			return arbitrary.Arbitrary{}, nil, fmt.Errorf("map shrinker cannot shrink %s", val.Value.Kind().String())
		default:
			next, shrinker, err := shrinker(val, propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, nil, err
			}

			next.Value = reflect.MakeMap(val.Value.Type())
			for _, node := range next.Elements {
				key, value := node.Elements[0], node.Elements[1]
				next.Value.SetMapIndex(key.Value, value.Value)
			}

			return next, Map(shrinker), nil
		}
	}
}
