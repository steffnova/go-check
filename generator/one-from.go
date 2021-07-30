package generator

import (
	"reflect"
)

// OneFrom is Arbitrary that will create Generator frome one of the passed
// arbitraries.
func OneFrom(first Arbitrary, other ...Arbitrary) Arbitrary {
	return func(target reflect.Type, r Random) (Generator, error) {
		arbitraries := append([]Arbitrary{first}, other...)

		index := int(r.Int64(0, int64(len(arbitraries)-1)))
		return arbitraries[index](target, r)
	}
}
