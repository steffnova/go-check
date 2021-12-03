package generator

import (
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

// OneFrom is Arbitrary that will create Generator frome one of the passed
// arbitraries.
func OneFrom(first Arbitrary, other ...Arbitrary) Arbitrary {
	return func(target reflect.Type, r Random) (Generator, error) {
		arbitraries := append([]Arbitrary{first}, other...)

		index := int(r.Int64(constraints.Int64{
			Min: 0,
			Max: int64(len(arbitraries) - 1),
		}))
		return arbitraries[index](target, r)
	}
}
