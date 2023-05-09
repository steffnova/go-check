package shrinker

import (
	"fmt"

	"github.com/steffnova/go-check/arbitrary"
)

func Array(original arbitrary.Arbitrary) arbitrary.Shrinker {
	switch {
	case original.Value.Len() != len(original.Elements):
		return Fail(fmt.Errorf("Invalid number of elements. Expected: %d", len(original.Elements)))
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
