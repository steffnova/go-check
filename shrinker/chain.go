package shrinker

import "github.com/steffnova/go-check/arbitrary"

func Chain(shrinkers ...Shrinker) Shrinker {
	if len(shrinkers) == 0 {
		return nil
	}

	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
		for index, shrinker := range shrinkers {
			if shrinker == nil {
				continue
			}
			arb, next, err := shrinker(arb, propertyFailed)
			if err != nil {
				return arbitrary.Arbitrary{}, nil, err
			}
			shrinkers[index] = next
			return arb, Chain(shrinkers[index:]...), nil
		}

		return arb, nil, nil
	}
}
