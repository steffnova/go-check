package shrinker

import (
	"fmt"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func CollectionSize(arbs []arbitrary.Arbitrary, shrinkers []Shrinker, index int, limits constraints.Length) Shrinker {
	switch {
	case len(shrinkers) != len(arbs):
		return Fail(fmt.Errorf("shrinker, nodes miss match"))
	case index < 0 || index > int(limits.Max):
		return Fail(fmt.Errorf("number of indexes out of range"))
	case int(limits.Min) == len(arbs)-index:
		return Chain(
			CollectionElement(shrinkers...),
			CollectionElements(shrinkers...),
		)
	default:
		return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
			nextShrinkers := []Shrinker{}
			nextShrinkers = append(nextShrinkers, shrinkers[:index]...)
			nextShrinkers = append(nextShrinkers, shrinkers[index+1:]...)

			nextArbs := []arbitrary.Arbitrary{}
			nextArbs = append(nextArbs, arbs[:index]...)
			nextArbs = append(nextArbs, arbs[index+1:]...)

			shrinker1 := CollectionSize(nextArbs, nextShrinkers, index, limits)
			shrinker2 := CollectionSize(arbs, shrinkers, index+1, limits)
			shrinker := shrinker1.Or(shrinker2)

			return arbitrary.Arbitrary{Elements: nextArbs}, shrinker, nil
		}
	}
}
