package shrinker

import (
	"github.com/steffnova/go-check/arbitrary"
)

func Collection() arbitrary.Shrinker {
	return Chain(
		func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
			shrinker := Chain(
				CollectionSizeRemoveBack(len(arb.Elements)-1),
				CollectionSizeRemoveFront(0),
			)
			return shrinker(arb, propertyFailed)
		},
		func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
			return CollectionElements(arb)(arb, propertyFailed)
		},
	)
}
