package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

// Array returns Arbitrary that creates array Generator. Array's element values
// are generate with Arbitrary provided in element parameter. Array's size is defined
// by Generator's target. Error is returned If target's kind is not reflect.Array
// or if Generator creation for array's elements fails.
func Array(element Arbitrary) Arbitrary {
	return func(target reflect.Type, r Random) (Generator, error) {
		if target.Kind() != reflect.Array {
			return nil, fmt.Errorf("target arbitrary's kind must be Array. Got: %s", target.Kind())
		}
		generate, err := element(target.Elem(), r.Split())
		if err != nil {
			return nil, fmt.Errorf("failed to crete generator. %s", err)
		}

		return func() arbitrary.Type {
			val := reflect.New(reflect.ArrayOf(target.Len(), target.Elem())).Elem()
			elements := make([]arbitrary.Type, target.Len())
			// for index := 0; index < target.Len(); index++ {
			for index := range elements {
				elements[index] = generate()
				val.Index(index).Set(elements[index].Value())
			}

			return arbitrary.Array{
				Elements: elements,
				Val:      val,
			}
		}, nil
	}
}

// ArrayFrom returns Arbitrary that creates array Generator. Unlike Array, ArrayFrom
// accepts the variadic number of Arbitraries through arbs parameter, where each arb is
// used to generate one element of the array. This behavior allows imposing different
// constraints for each element in the array. Array's size is defined by Generator's target.
// Error is returned If target's kind is reflect.Array, len(arbs) doesn't match the size
// target array or Generator creation for any of the array's elements fails.
func ArrayFrom(arbs ...Arbitrary) Arbitrary {
	return func(target reflect.Type, r Random) (Generator, error) {
		if target.Kind() != reflect.Array {
			return nil, fmt.Errorf("target arbitrary's kind must be Array. Got: %s", target.Kind())
		}
		if target.Len() != len(arbs) {
			return nil, fmt.Errorf("invalid number of arbs. Expected: %d", target.Len())
		}

		generators := make([]Generator, target.Len())
		for index := range generators {
			generator, err := arbs[index](target.Elem(), r.Split())
			if err != nil {
				return nil, fmt.Errorf("failed to create element's generator. %s", err)
			}
			generators[index] = generator
		}

		return func() arbitrary.Type {
			val := reflect.New(reflect.ArrayOf(target.Len(), target.Elem())).Elem()
			elements := make([]arbitrary.Type, target.Len())
			for index := 0; index < target.Len(); index++ {
				elements[index] = generators[index]()
				val.Index(index).Set(elements[index].Value())
			}

			return arbitrary.Array{
				Elements: elements,
				Val:      val,
			}
		}, nil
	}
}
