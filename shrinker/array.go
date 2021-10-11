package shrinker

import (
	"fmt"
	"reflect"
)

// Array is a shrinker for array. Array is shrinked by shrinking it's elements one at a time
// Convergence speed for shrinker is O(n*m), n is array size and m is convergance speed of
// array elements.
func Array(val reflect.Value, shrinkers []Shrinker) Shrinker {
	mapperSignature := reflect.FuncOf(
		[]reflect.Type{val.Slice(0, val.Len()).Type()},
		[]reflect.Type{val.Type()},
		false,
	)

	mapper := reflect.MakeFunc(mapperSignature, func(arg []reflect.Value) []reflect.Value {
		newArray := reflect.New(val.Type()).Elem()
		for index, slice := 0, arg[0]; index < slice.Len(); index++ {
			newArray.Index(index).Set(slice.Index(index))
		}
		return []reflect.Value{newArray}
	})

	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		switch {
		case val.Kind() != reflect.Array:
			return reflect.Value{}, nil, fmt.Errorf("array shrinker cannot shrink %s", val.Kind().String())
		case val.Len() != len(shrinkers):
			return reflect.Value{}, nil, fmt.Errorf("size of array must match the number of shrinkers")
		default:
			shrinker := sliceElements(val.Slice(0, val.Len()), shrinkers...).Map(val.Type(), mapper.Interface())
			return shrinker(propertyFailed)
		}
	}
}
