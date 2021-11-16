package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

// Struct is a shrinker for Go's struct. Struct shrinking consists of shrinking individual
// fields one by one. Convergance speed for shrinker is O(n*m), n is number of Struct fields
// and m is convergance speed of field type.
func Struct(structType reflect.Type, fieldShrinks []Shrink) Shrinker {
	switch {
	case structType.Kind() != reflect.Struct:
		return Invalid(fmt.Errorf("struct shrinker cannot shrink: %s", structType.Kind().String()))
	case structType.NumField() != len(fieldShrinks):
		return Invalid(fmt.Errorf("number shrinks must match number of struct fields"))
	default:
		sliceType := reflect.TypeOf([]interface{}{})

		mapFn := func(in reflect.Value) reflect.Value {
			out := reflect.New(structType).Elem()
			for index := 0; index < in.Len(); index++ {
				fieldValue := in.Index(index).Interface()
				out.Field(index).Set(reflect.ValueOf(fieldValue))
			}
			return out
		}

		return SliceElements(SliceShrink{
			Type:     sliceType,
			Elements: fieldShrinks,
		}).Map(arbitrary.Mapper(sliceType, structType, mapFn))
	}
}
