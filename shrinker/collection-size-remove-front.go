package shrinker

import (
	"fmt"

	"github.com/steffnova/go-check/arbitrary"
)

func CollectionSizeRemoveFront(index int) arbitrary.Shrinker {
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
		switch {
		case index < 0:
			return arbitrary.Arbitrary{}, fmt.Errorf("index is out of range")
		case index >= len(arb.Elements) && propertyFailed:
			reduced := arb.Copy()
			reduced.Shrinker = CollectionElements(reduced)
			return reduced, nil
		default:
			reduced := arb.Copy()
			elements := []arbitrary.Arbitrary{}
			elements = append(elements, arb.Elements[:index]...)
			elements = append(elements, arb.Elements[index+1:]...)

			revertRemoval := func(in arbitrary.Arbitrary) arbitrary.Arbitrary {
				in.Elements = arb.Elements
				return in
			}

			shrinker1 := CollectionSizeRemoveFront(index)
			shrinker2 := CollectionSizeRemoveFront(index + 1).TransformOnceBefore(revertRemoval)
			shrinker := shrinker1.Or(shrinker2)

			reduced.Elements = elements
			reduced.Shrinker = shrinker
			return reduced, nil
		}
	}
}
