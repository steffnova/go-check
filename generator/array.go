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
		if target.Kind() != reflect.Array {
			return nil, fmt.Errorf("can't use Array generator for %s type", target)
		}

		if _, err := element(target.Elem(), bias, r); err != nil {
			return nil, fmt.Errorf("can't use array generator for %s type. %s", target, err)
		}

		generators := make([]Generator, target.Len())
		for index := range generators {
			generators[index] = element
		}

		generator := ArrayFrom(generators...)

		return generator(target, bias, r)
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
			return nil, fmt.Errorf("target arbitrary's kind must be Array. Got: %s", target.Kind())
		}
		if target.Len() != len(elements) {
			return nil, fmt.Errorf("invalid number of arbs. Expected: %d", target.Len())
		}

		generators := make([]Generate, target.Len())
		for index := range generators {
			generator, err := elements[index](target.Elem(), bias, r)
			if err != nil {
				return nil, fmt.Errorf("failed to create element's generator. %s", err)
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

			return arb, shrinker.Array(shrinker.Chain(
				shrinker.CollectionElement(shrinkers...),
				shrinker.CollectionElements(shrinkers...),
			))
		}, nil
	}
}
