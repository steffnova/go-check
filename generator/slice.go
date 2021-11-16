package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Slice returns Arbitrary that creates slice Generator. Slice's element values
// are generate with Arbitrary provided in element parameter. Range in which Slice's
// size generated is defined my limits parameter. Even though limits is a variadic
// argument only the first value is used for defining constraints. Error is returned
// if target's reflect.Kind is not Slice, if creation of Generator for slice's elements
// fails or creation of Generator for slice's size fail.
func Slice(element Arbitrary, limits ...constraints.Length) Arbitrary {
	return func(target reflect.Type, r Random) (Generator, error) {
		constraint := constraints.LengthDefault()
		if len(limits) != 0 {
			constraint = limits[0]
		}
		if target.Kind() != reflect.Slice {
			return nil, fmt.Errorf("targets kind must be Slice. Got: %s", target.Kind())
		}

		generator, err := element(target.Elem(), r)
		if err != nil {
			return nil, fmt.Errorf("failed to create generator for slice elements: %s", err)
		}

		return func() (reflect.Value, shrinker.Shrinker) {
			size := r.Int64(int64(constraint.Min), int64(constraint.Max))
			elements := make([]shrinker.Shrink, size)

			for index := range elements {
				elementValue, elementShrinker := generator()
				elements[index] = shrinker.Shrink{
					Value:    elementValue,
					Shrinker: elementShrinker,
				}
			}

			shrinkable := shrinker.SliceShrink{
				Type:     target,
				Elements: elements,
			}

			return shrinkable.Value(), shrinker.Slice(shrinkable, 0, constraint)
		}, nil
	}

}
