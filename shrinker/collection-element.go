package shrinker

import (
	"fmt"

	"github.com/steffnova/go-check/arbitrary"
)

func CollectionElement(firstRun bool, shrinkers ...arbitrary.Shrinker) arbitrary.Shrinker {
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
		if firstRun {
			for index := range arb.Elements {
				arb.Elements[index].Shrinker = shrinkers[index]
			}
		}
		nextShrinkers := make([]arbitrary.Shrinker, len(shrinkers))
		copy(nextShrinkers, shrinkers)
		for index, element := range arb.Elements {
			if element.Shrinker == nil {
				continue
			}

			var err error
			shrink, err := element.Shrinker(element, propertyFailed)
			// arb.Elements[index], nextShrinkers[index], err = shrinker(arb.Elements[index], propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, err
			}
			nextShrinkers[index] = shrink.Shrinker
			arb.Elements[index] = shrink
			arb.Shrinker = CollectionElement(false, nextShrinkers...)

			return arb, nil
		}

		return arb, nil
	}
}

type collectionValue func(elements arbitrary.Arbitraries) (arbitrary.Arbitrary, error)

func CollectionOneElement() arbitrary.Shrinker {
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
		// arb = arb.Copy()
		for index, element := range arb.Elements {
			if element.Shrinker == nil {
				continue
			}
			fmt.Printf("Shrinker at index %d is not nil\n", index)

			shrink, err := element.Shrinker(element, propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, nil
			}
			fmt.Println(shrink.Value)

			arb.Elements[index] = shrink
			// arb.Shrinker = CollectionOneElement()
			return arb, nil
		}

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
			return arb, nil
		}

		arb.Shrinker = CollectionAllElements()
		return arb, nil
	}
}
