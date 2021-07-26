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
			for index := 0; index < target.Len(); index++ {
				val.Index(index).Set(generate().Value())
			}

			return arbitrary.Array{
				Val: val,
			}
		}, nil
	}
}
