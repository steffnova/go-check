package generator

import (
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// String returns generator for string types. Range of slice size is defined by
// "limits" parameter. If "limits" parameter is not specified default [0, 100]
// range is used instead. Error is returned if generator's target is not a
// string type, or limits.Min > limits.Max
func String(limits ...constraints.String) Generator {
	constraint := constraints.StringDefault()
	if len(limits) != 0 {
		constraint = limits[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		mapper := arbitrary.Mapper(reflect.TypeOf([]rune{}), target, func(in reflect.Value) reflect.Value {
			return in.Convert(target)
		})
		return Slice(
			Rune(constraint.Rune),
			constraint.Length,
		).Map(mapper)(target, bias, r)
	}
}
