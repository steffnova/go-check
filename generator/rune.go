package generator

import (
	"fmt"

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
	switch {
	case constraint.MinCodePoint > constraint.MaxCodePoint:
		return Invalid(fmt.Errorf("minimal code point %d can't be greater than maximal code point: %d", constraint.MinCodePoint, constraint.MaxCodePoint))
	case constraint.MinCodePoint < 0:
		return Invalid(fmt.Errorf("minimal code point must be greater then or equal to 0"))
	case constraint.MaxCodePoint > 0x10ffff:
		return Invalid(fmt.Errorf("maximal code point must be lower then or equal to 0x10ffff"))
	default:
		return Int32(constraints.Int32{
			Max: constraint.MaxCodePoint,
			Min: constraint.MinCodePoint,
		}).Map(func(n int32) rune {
			return rune(n)
		})
	}
}
