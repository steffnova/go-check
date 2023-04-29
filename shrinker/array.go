package shrinker

import (
	"github.com/steffnova/go-check/arbitrary"
)

func Array(original arbitrary.Arbitrary, shrinkers []Shrinker) Shrinker {
	switch {
	default:
		return Chain(
			CollectionElement(shrinkers...),
			CollectionElements(shrinkers...),
		).
			Transform(arbitrary.NewArray(original.Value.Type())).
			Validate(arbitrary.ValidateArray())
	}
}
