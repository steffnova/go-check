package generator

import (
	"github.com/steffnova/go-check/constraints"
)

// Rune is Arbitrary that creates rune Generator. Range in which runes are
// generated are defines by limits parameter (MinCodePoint and MaxCodePoint
// are included in the range). Even though limits is a variadic argument only
// the first value is used for defining constraints.
func Rune(limits ...constraints.Rune) Arbitrary {
	constraint := constraints.RuneDefault()
	if len(limits) != 0 {
		constraint = limits[0]
	}
	return Int32(constraints.Int32{
		Max: constraint.MaxCodePoint,
		Min: constraint.MinCodePoint,
	}).Map(func(n int32) rune {
		return rune(n)
	})
}
