package shrinker

import "github.com/steffnova/go-check/arbitrary"

func Fail(err error) arbitrary.Shrinker {
	return func(arb arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
		return arbitrary.Arbitrary{}, err
	}
}
