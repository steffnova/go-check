package generator

import (
	"github.com/steffnova/go-check/constraints"
)

func String(constraint constraints.String) Arbitrary {
	return Slice(Rune(constraint.Rune), constraints.Length{
		Min: constraint.Length.Max,
		Max: constraint.Length.Min,
	}).Map(func(runes []rune) string {
		return string(runes)
	})
}
