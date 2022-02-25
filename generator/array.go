package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

func Array(element Generator) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		if target.Kind() != reflect.Array {
			return nil, fmt.Errorf("target arbitrary's kind must be Array. Got: %s", target.Kind())
		}

		mapper := arbitrary.Mapper(reflect.SliceOf(target.Elem()), target, func(in reflect.Value) reflect.Value {
			out := reflect.New(target).Elem()
			for index := 0; index < in.Len(); index++ {
				val := in.Index(index).Interface()
				out.Index(index).Set(reflect.ValueOf(val))
			}
			return out
		})

		generator := Slice(element, constraints.Length{
			Min: target.Len(),
			Max: target.Len(),
		}).Map(mapper)

		return generator(target, bias, r)
	}
}

// ArrayFrom returns Arbitrary that creates array Generator. Unlike Array, ArrayFrom
// accepts the variadic number of Arbitraries through arbs parameter, where each arb is
// used to generate one element of the array. This behavior allows imposing different
// constraints for each element in the array. Array's size is defined by Generator's target.
// Error is returned If target's kind is reflect.Array, len(arbs) doesn't match the size
// target array or Generator creation for any of the array's elements fails.
func ArrayFrom(arbs ...Generator) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		if target.Kind() != reflect.Array {
			return nil, fmt.Errorf("target arbitrary's kind must be Array. Got: %s", target.Kind())
		}
		if target.Len() != len(arbs) {
			return nil, fmt.Errorf("invalid number of arbs. Expected: %d", target.Len())
		}

		generators := make([]Generate, target.Len())
		for index := range generators {
			generator, err := arbs[index](target.Elem(), bias, r)
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
