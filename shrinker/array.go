package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

func Array(original arbitrary.Arbitrary) arbitrary.Shrinker {
	switch {
	case original.Value.Kind() != reflect.Array:
		return Fail(fmt.Errorf("Array shrinker can's shrink: %s", original.Value.Kind()))
	case original.Value.Len() != len(original.Elements):
		return Fail(fmt.Errorf("invalid number of elements. Expected: %d", len(original.Elements)))
	default:
		shrinkers := make([]arbitrary.Shrinker, len(original.Elements))
		for index, element := range original.Elements {
			shrinkers[index] = element.Shrinker
		}

		return CollectionElements(original).
			TransformAfter(arbitrary.NewArray(original.Value.Type())).
			Validate(arbitrary.ValidateArray())
	}
}
