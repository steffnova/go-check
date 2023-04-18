package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Rune returns generator for rune types. Range of rune values that can be
// generated is defined by "limits" parameter. If no limits are provided default
// [0, 0x10ffff] code point range is used which includes all Unicode16 characters.
// Error is returned if minimal code point is greater than maximal code point,
// minimal code point is lower than 0 or maximal code point is greater than 0x10ffff
func Rune(limits ...constraints.Rune) Generator {
	constraint := constraints.RuneDefault()
	if len(limits) != 0 {
		constraint = limits[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		switch {
		case target.Kind() != reflect.Int32:
			return nil, NewErrorInvalidTarget(target, "Rune")
		case constraint.MinCodePoint > constraint.MaxCodePoint:
			return nil, fmt.Errorf("%w. Minimal code point %d can't be greater than maximal code point: %d", ErrorInvalidConstraints, constraint.MinCodePoint, constraint.MaxCodePoint)
		case constraint.MinCodePoint < 0:
			return nil, fmt.Errorf("%w. Minimal code point must be greater then or equal to 0", ErrorInvalidConstraints)
		case constraint.MaxCodePoint > 0x10ffff:
			return nil, fmt.Errorf("%w. Maximal code point must be lower then or equal to 0x10ffff", ErrorInvalidConstraints)
		default:
			mapper := arbitrary.Mapper(reflect.TypeOf(int32(0)), target, func(in reflect.Value) reflect.Value {
				return reflect.ValueOf(rune(int32(in.Int()))).Convert(target)
			})
			return Int32(constraints.Int32{
				Max: constraint.MaxCodePoint,
				Min: constraint.MinCodePoint,
			}).Map(mapper)(target, bias, r)
		}
	}
}
