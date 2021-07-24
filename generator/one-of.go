package generator

import (
	"math/rand"
	"reflect"
)

// OneOf is Arbitrary that will create Generator for one of the passed
// arbitraries.
func OneOf(first Arbitrary, other ...Arbitrary) Arbitrary {
	return func(target reflect.Type, r *rand.Rand) (Generator, error) {
		arbitraries := append([]Arbitrary{first}, other...)
		arb := arbitraries[r.Intn(len(arbitraries))]

		gen, err := arb(target, r)
		if err != nil {
			return nil, err
		}

		return gen, nil
	}
}
