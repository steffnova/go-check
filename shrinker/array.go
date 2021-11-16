package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

// Array is a shrinker for array. Array is shrinked by shrinking it's elements one at a time
// Convergence speed for shrinker is O(n*m), n is array size and m is convergance speed of
// array elements.
func Array(arrayType reflect.Type, elements []Shrink) Shrinker {
	switch {
	case arrayType.Kind() != reflect.Array:
		return Invalid(fmt.Errorf("array shrinker cannot shrink %s", arrayType.Kind().String()))
	case arrayType.Len() != len(elements):
		return Invalid(fmt.Errorf("number of shrinkable elements must match size of the array"))
	default:
		sliceType := reflect.SliceOf(arrayType.Elem())

		mapFn := func(in reflect.Value) reflect.Value {
			newArray := reflect.New(arrayType).Elem()
			for index := 0; index < arrayType.Len(); index++ {
				newArray.Index(index).Set(in.Index(index))
			}
			return newArray
		}

		return SliceElements(SliceShrink{
			Type:     reflect.SliceOf(arrayType.Elem()),
			Elements: elements,
		}).Map(arbitrary.Mapper(sliceType, arrayType, mapFn))
	}
}
