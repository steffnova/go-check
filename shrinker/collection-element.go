package shrinker

import (
	"github.com/steffnova/go-check/arbitrary"
)

func CollectionOneElement() arbitrary.Shrinker {
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
		arb = arb.Copy()
		for index, element := range arb.Elements {
			if element.Shrinker == nil {
				continue
			}

			shrink, err := element.Shrinker(element, propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, nil
			}

			arb.Elements[index] = shrink
			arb.Shrinker = CollectionOneElement()

			return arb, nil
		}

		arb.Shrinker = nil
		return arb, nil
	}
}

func CollectionAllElements() arbitrary.Shrinker {
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
		arb = arb.Copy()
		canShrink := false

		for index, element := range arb.Elements {
			if element.Shrinker == nil {
				continue
			}

			canShrink = true
			shrink, err := element.Shrinker(element, propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, nil
			}

			arb.Elements[index] = shrink
		}

		if !canShrink {
			arb.Shrinker = nil
			return arb, nil
		}

		arb.Shrinker = CollectionAllElements()
		return arb, nil
	}
}
