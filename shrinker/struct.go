package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

// Struct is a shrinker for struct. Struct shrinking consists of shrinking individual
// fields one by one. Error is returned if struct type is not struct, length of fieldShrinks
// is not equal to number of struct fields or if any of the struct fields return an error
// during shrinking.
func Struct(structType reflect.Type, fieldShrinks []Shrink) Shrinker {
	switch {
	case structType.Kind() != reflect.Struct:
		return Invalid(fmt.Errorf("struct shrinker cannot shrink: %s", structType.Kind().String()))
	case structType.NumField() != len(fieldShrinks):
		return Invalid(fmt.Errorf("number shrinks must match number of struct fields"))
	default:
		sliceType := reflect.ArrayOf(len(fieldShrinks), arbitrary.Type)

		mapFn := func(in reflect.Value) reflect.Value {
			out := reflect.New(structType).Elem()
			for index := 0; index < in.Len(); index++ {
				fieldValue := in.Index(index).Interface()
				out.Field(index).Set(reflect.ValueOf(fieldValue))
			}
			return out
		}

		return Array(sliceType, fieldShrinks).Map(arbitrary.Mapper(sliceType, structType, mapFn))
	}
}
