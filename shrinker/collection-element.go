package shrinker

import (
	"github.com/steffnova/go-check/arbitrary"
)

func CollectionElement(shrinkers ...Shrinker) Shrinker {
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
		nextShrinkers := make([]Shrinker, len(shrinkers))
		copy(nextShrinkers, shrinkers)
		for index, shrinker := range nextShrinkers {
			if shrinker == nil {
				continue
			}

			var err error
			arb.Elements[index], nextShrinkers[index], err = shrinker(arb.Elements[index], propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, nil, err
			}

			return arb, CollectionElement(nextShrinkers...), nil
		}

		return arb, nil, nil
	}
}
