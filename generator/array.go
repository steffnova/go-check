package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Array returns generator for array types. Array element's generator is specified by "element"
// parameter. Error is returned if generator's target is not array type or element's generator
// returns an error.
func Array(element arbitrary.Generator) arbitrary.Generator {
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		generators := make([]arbitrary.Generator, target.Len())
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
func ArrayFrom(elements ...arbitrary.Generator) arbitrary.Generator {
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		if target.Kind() != reflect.Array {
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "Array")
		}
		if target.Len() != len(elements) {
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidCollectionSize(target.Len(), len(elements))
		}

		value := reflect.New(target).Elem()
		arbitraries := make([]arbitrary.Arbitrary, target.Len())

		for index := range elements {
			arb, err := elements[index](target.Elem(), bias, r)
			if err != nil {
				return arbitrary.Arbitrary{}, fmt.Errorf("%w. Failed to use gerator for array's (%s) element %d.", err, target.Kind(), index)
			}
			arbitraries[index] = arb
			value.Index(index).Set(arbitraries[index].Value)
		}

		return arbitrary.Arbitrary{
			Value:    value,
			Elements: arbitraries,
			// Shrinker: shrinker.Array(value, shrinkers),
		}, nil
	}
}
