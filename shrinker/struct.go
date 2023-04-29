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
func Struct(original arbitrary.Arbitrary, shrinkers []Shrinker) Shrinker {
	switch {
	case original.Value.Kind() != reflect.Struct:
		return Fail(fmt.Errorf("struct shrinker cannot shrink %s", original.Value.Kind().String()))
	case original.Value.NumField() != len(original.Elements):
		return Fail(fmt.Errorf("number of struct arbitraries %d must match number of struct fields %d", len(original.Elements), original.Value.Len()))
	case len(original.Elements) != len(shrinkers):
		return Fail(fmt.Errorf("Number of shrinkers must match number of struct fields"))
	default:
		return Chain(
			CollectionElement(shrinkers...),
			CollectionElements(shrinkers...),
		).
			Transform(arbitrary.NewStruct(original.Value.Type())).
			Validate(arbitrary.ValidateStruct())
	}
}
