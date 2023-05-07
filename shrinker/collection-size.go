package shrinker

import (
	"fmt"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func CollectionSize(arbs []arbitrary.Arbitrary, index int, limits constraints.Length) arbitrary.Shrinker {
	shrinkers := make([]arbitrary.Shrinker, len(arbs))
	for index := range arbs {
		shrinkers[index] = arbs[index].Shrinker
	}
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
		switch {
		// case len(shrinkers) != len(arbs):
		// 	return arbitrary.Arbitrary{}, nil, fmt.Errorf("shrinker, nodes miss match")
		case index < 0 || index > int(limits.Max):
			return arbitrary.Arbitrary{}, fmt.Errorf("number of indexes out of range")
		case int(limits.Min) == len(arbs)-index:
			return arbitrary.Arbitrary{
				Elements: arbs,
				Shrinker: Chain(
					CollectionOneElement(),
					CollectionAllElements(),
				)}, nil
		default:
			nextArbs := []arbitrary.Arbitrary{}
			nextArbs = append(nextArbs, arbs[:index]...)
			nextArbs = append(nextArbs, arbs[index+1:]...)

			shrinker1 := CollectionSize(nextArbs, index, limits)
			shrinker2 := CollectionSize(arbs, index+1, limits)
			shrinker := shrinker1.Or(shrinker2)

			return arbitrary.Arbitrary{
				Elements: nextArbs,
				Shrinker: shrinker,
			}, nil
		}
	}
}
