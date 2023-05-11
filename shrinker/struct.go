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
func Struct(original arbitrary.Arbitrary) arbitrary.Shrinker {
	switch {
	case original.Value.Kind() != reflect.Struct:
		return Fail(fmt.Errorf("struct shrinker cannot shrink %s", original.Value.Kind().String()))
	case original.Value.NumField() != len(original.Elements):
		return Fail(fmt.Errorf("number of struct arbitraries %d must match number of struct fields %d", len(original.Elements), original.Value.NumField()))
	default:
		shrinkers := make([]arbitrary.Shrinker, len(original.Elements))
		for index, element := range original.Elements {
			shrinkers[index] = element.Shrinker
		}

		return CollectionElements(original).
			TransformAfter(arbitrary.NewStruct(original.Value.Type())).
			Validate(arbitrary.ValidateStruct())
	}
}
