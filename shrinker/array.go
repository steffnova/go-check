package shrinker

import (
	"reflect"
)

// Array is a shrinker for array. Array is shrinked by shrinking it's elements one at a time
// Convergence speed for shrinker is O(n*m), n is array size and m is convergance speed of
// array elements.
func Array(original reflect.Value, shrinkers []Shrinker) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker) {
		newArray := reflect.New(original.Type()).Elem()
		reflect.Copy(newArray, original)

		for index, shrinker := range shrinkers {
			if shrinker == nil {
				continue
			}
			val, shrinker := shrinker(propertyFailed)

			newArray.Index(index).Set(val)
			shrinkers[index] = shrinker
			return newArray, Array(newArray, shrinkers)
		}

		return original, nil
	}
}
