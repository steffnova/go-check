package shrinker

import (
	"fmt"
	"reflect"
)

// Ptr is a shrinker for pointers. Shrinker is shrinker for type to which ptr points to
// points to. Shrinking process consists of shrinking the pointer's value, and shrinking
// of pointer to nil
func Ptr(val reflect.Value, shrinker Shrinker) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		switch {
		case val.Kind() != reflect.Ptr:
			return reflect.Value{}, nil, fmt.Errorf("ptr shrinker cannot shrink: %s", val.Kind().String())
		default:
			return ptr(val, val, shrinker)(propertyFailed)
		}
	}
}

func ptr(val reflect.Value, lastVal reflect.Value, shrinker Shrinker) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		if shrinker != nil {
			ptrVal, shrinker, err := shrinker(propertyFailed)
			if err != nil {
				return reflect.Value{}, nil, fmt.Errorf("failed to shrink ptr's value. %w", err)
			}
			val.Elem().Set(ptrVal)
			return val, ptr(val, val, shrinker), nil
		}

		switch {
		case !val.IsNil():
			// Shrinking to nil could make a property not fail. The last valid pointer
			// value needs to be saved in order to revert nil pointer if needed.
			lastVal, val = val, reflect.Zero(val.Type())
			return val, ptr(val, lastVal, shrinker), nil
		case propertyFailed:
			return val, nil, nil
		default:
			return lastVal, nil, nil
		}
	}
}
