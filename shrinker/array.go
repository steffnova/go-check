package shrinker

import (
	"fmt"

	"github.com/steffnova/go-check/arbitrary"
)

func Array(original arbitrary.Arbitrary, shrinkers []Shrinker) Shrinker {
	switch {
	case original.Value.Len() != len(original.Elements):
		return Fail(fmt.Errorf("Invalid number of elements. Expected: %d", len(original.Elements)))
	case len(original.Elements) != len(shrinkers):
		return Fail(fmt.Errorf("Number of shrinkers: %d must match number of elements: %d", len(shrinkers), len(original.Elements)))
	default:
		return Chain(
			CollectionElement(shrinkers...),
			CollectionElements(shrinkers...),
		).
			transformAfter(arbitrary.NewArray(original.Value.Type())).
			Validate(arbitrary.ValidateArray())
	}
}
