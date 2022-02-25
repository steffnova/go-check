package shrinker

import "github.com/steffnova/go-check/arbitrary"

func Invalid(err error) Shrinker {
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, Shrinker, error) {
		return arbitrary.Arbitrary{}, nil, err
	}
}
