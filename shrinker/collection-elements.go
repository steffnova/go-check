package shrinker

import "github.com/steffnova/go-check/arbitrary"

func CollectionElements(source ...Shrinker) Shrinker {
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
		shrinkers := make([]Shrinker, len(source))
		copy(shrinkers, source)
		canShrink := false

		for index, shrinker := range shrinkers {
			if shrinker == nil {
				continue
			}
			canShrink = true

			var err error
			arb.Elements[index], shrinkers[index], err = shrinker(arb.Elements[index], propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, nil, err
			}
		}

		if canShrink {
			return arb, CollectionElements(shrinkers...), nil
		}

		return arb, nil, nil
	}
}
