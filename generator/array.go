package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Array returns generator for array types. Array element's generator is specified by "element"
// parameter. Error is returned if generator's target is not array type or element's generator
// returns an error.
func Array(element Generator) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		generators := make([]Generator, target.Len())
		for index := range generators {
			generators[index] = element
		}

		return ArrayFrom(generators...)(target, bias, r)
	}
}

// ArrayFrom returns generator of array types. Unlike Array where one generator is used for all
// elements of array, ArrayFrom accepts a generator for each individual element which is
// specified with element parameter. This behavior allows imposing different constraints
// for each array element. Error is returned if generator's target is not array type, number
// of element generators doesn't match the size of the array, or any of the element generators
// return an error.
func ArrayFrom(elements ...Generator) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		if target.Kind() != reflect.Array {
			return nil, NewErrorInvalidTarget(target, "Array")
		}
		if target.Len() != len(elements) {
			return nil, NewErrorInvalidCollectionSize(target.Len(), len(elements))
		}

		generators := make([]Generate, target.Len())
		for index := range generators {
			generator, err := elements[index](target.Elem(), bias, r)
			if err != nil {
				return nil, fmt.Errorf("%w. Failed to use gerator for array's (%s) element %d.", err, target.Kind(), index)
			}
			generators[index] = generator
		}

		return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
			arb := arbitrary.Arbitrary{
				Value:    reflect.New(target).Elem(),
				Elements: make(arbitrary.Arbitraries, target.Len()),
			}

			shrinkers := make([]shrinker.Shrinker, target.Len())

			for index, generator := range generators {
				arb.Elements[index], shrinkers[index] = generator()
				arb.Value.Index(index).Set(arb.Elements[index].Value)
			}

			return arb, shrinker.Array(arb, shrinkers)
		}, nil
	}
}
