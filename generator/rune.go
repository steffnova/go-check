package generator

import (
	"github.com/steffnova/go-check/constraints"
)

func Rune(constraint constraints.Rune) Arbitrary {
	return Int32(constraints.Int32{
		Max: constraint.MaxCodePoint,
		Min: constraint.MinCodePoint,
	}).Map(func(n int32) rune {
		return rune(n)
	})
}
