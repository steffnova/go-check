package shrinker

import "github.com/steffnova/go-check/arbitrary"

func Chain(shrinkers ...arbitrary.Shrinker) arbitrary.Shrinker {
	if len(shrinkers) == 0 {
		return nil
	}

	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
		for index, shrinker := range shrinkers {
			if shrinker == nil {
				continue
			}
			arb, err := shrinker(arb, propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, err
			}
			shrinkers[index] = arb.Shrinker
			arb.Shrinker = Chain(shrinkers[index:]...)
			return arb, nil
		}

		arb.Shrinker = nil
		return arb, nil
	}
}
