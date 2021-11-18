package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Slice is a shrinker for slice. Slice is shrinked by two dimensions: elements and size.
// Shrinking is first done by size, and then by elements. Error is returned if sliceType
// is not slice, length of elements is out of limits range [min, max] or if any of the
// elements returns an error during shrinking.
func Slice(sliceType reflect.Type, elements []Shrink, index int, limits constraints.Length) Shrinker {
	switch {
	case sliceType.Kind() != reflect.Slice:
		return Invalid(fmt.Errorf("slice shrinker cannot shrink: %s", sliceType.Kind().String()))
	case index < 0 || index > len(elements):
		return Invalid(fmt.Errorf("index: %d is out of slice range", index))
	case limits.Min == len(elements)-index:
		arrayType := reflect.ArrayOf(len(elements), arbitrary.Type)
		mapFn := func(in reflect.Value) reflect.Value {
			out := reflect.MakeSlice(sliceType, len(elements), len(elements))
			for index := 0; index < in.Len(); index++ {
				val := in.Index(index).Interface()
				out.Index(index).Set(reflect.ValueOf(val))
			}
			return out
		}
		return Array(arrayType, elements).Map(arbitrary.Mapper(arrayType, sliceType, mapFn))
	default:
		return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
			nextElements := []Shrink{}
			nextElements = append(nextElements, elements[:index]...)
			nextElements = append(nextElements, elements[index+1:]...)

			shrinker1 := Slice(sliceType, nextElements, index, limits)
			shrinker2 := Slice(sliceType, elements, index+1, limits)

			out := reflect.MakeSlice(sliceType, len(nextElements), len(nextElements))
			for index, element := range nextElements {
				out.Index(index).Set(element.Value)
			}

			return out, shrinker1.Or(shrinker2), nil
		}
	}
}
