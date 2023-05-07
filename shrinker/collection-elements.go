package shrinker

import "github.com/steffnova/go-check/arbitrary"

func CollectionElements(firstRun bool, source ...arbitrary.Shrinker) arbitrary.Shrinker {
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
		if firstRun {
			for index := range arb.Elements {
				arb.Elements[index].Shrinker = source[index]
			}
		}

		shrinkers := make([]arbitrary.Shrinker, len(source))
		copy(shrinkers, source)
		canShrink := false

		for index, element := range arb.Elements {
			if element.Shrinker == nil {
				continue
			}
			canShrink = true

			shrink, err := element.Shrinker(element, propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, err
			}
			arb.Elements[index] = shrink
			shrinkers[index] = shrink.Shrinker
		}

		if canShrink {
			arb.Shrinker = CollectionElements(false, shrinkers...)
			return arb, nil
		}

		return arb, nil
	}
}
