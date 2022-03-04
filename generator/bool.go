package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Bool returns generator of bool types. Error is returned if generator's
// target is not bool type.
func Bool() Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		switch {
		case target.Kind() != reflect.Bool:
			return nil, fmt.Errorf("can't use Bool generator for %s type", target)
		default:
			mapper := arbitrary.Mapper(reflect.TypeOf(uint64(0)), target, func(in reflect.Value) reflect.Value {
				return reflect.ValueOf(in.Uint() != 0).Convert(target)
			})
			return Uint64(constraints.Uint64{
				Min: 0,
				Max: 1,
			}).Map(mapper)(target, bias, r)
		}
	}
}
