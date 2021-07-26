package generator

import (
	"reflect"
)

// OneOf is Arbitrary that will create Generator for one of the passed
// arbitraries.
func OneOf(first Arbitrary, other ...Arbitrary) Arbitrary {
	return func(target reflect.Type, r Random) (Generator, error) {
		arbitraries := append([]Arbitrary{first}, other...)

		index := int(r.Int64(0, int64(len(arbitraries)-1)))
		arb := arbitraries[index]

		gen, err := arb(target, r.Split())
		if err != nil {
			return nil, err
		}

		return gen, nil
	}
}
