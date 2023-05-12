package shrinker

import (
	"fmt"

	"github.com/steffnova/go-check/arbitrary"
)

func CollectionSizeRemoveBack(index int) arbitrary.Shrinker {
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
		switch {
		case index >= len(arb.Elements):
			return arbitrary.Arbitrary{}, fmt.Errorf("index is out of range")
		case index < 0 && propertyFailed:
			reduced := arb.Copy()
			reduced.Shrinker = nil
			reduced.Shrinker = CollectionElements(reduced)
			return reduced, nil
		default:
			reduced := arb.Copy()
			elements := []arbitrary.Arbitrary{}
			elements = append(elements, arb.Elements[:index]...)
			elements = append(elements, arb.Elements[index+1:]...)

			revertRemoval := func(in arbitrary.Arbitrary) arbitrary.Arbitrary {
				in = in.Copy()
				in.Elements = arb.Elements
				return in
			}

			shrinker1 := CollectionSizeRemoveBack(index - 1)
			shrinker2 := CollectionSizeRemoveBack(index - 1).TransformOnceBefore(revertRemoval)
			shrinker := shrinker1.Or(shrinker2)

			reduced.Elements = elements
			reduced.Shrinker = shrinker
			return reduced, nil
		}
	}
}
