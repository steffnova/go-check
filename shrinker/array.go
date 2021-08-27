package shrinker

import (
	"fmt"
	"reflect"
)

// Array is a shrinker for array. Array is shrinked by shrinking it's elements one at a time
// Convergence speed for shrinker is O(n*m), n is array size and m is convergance speed of
// array elements.
func Array(val reflect.Value, shrinkers []Shrinker) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		if val.Kind() != reflect.Array {
			return reflect.Value{}, nil, fmt.Errorf("array shrinker cannot shrink %s", val.Kind().String())
		}

		newArray := reflect.New(val.Type()).Elem()
		reflect.Copy(newArray, val)

		for index, shrinker := range shrinkers {
			if shrinker == nil {
				continue
			}

			val, shrinker, err := shrinker(propertyFailed)
			if err != nil {
				return reflect.Value{}, nil, fmt.Errorf("failed to shrink array element. %w", err)
			}

			newArray.Index(index).Set(val)
			shrinkers[index] = shrinker
			return newArray, Array(newArray, shrinkers), nil
		}

		return val, nil, nil
	}
}
