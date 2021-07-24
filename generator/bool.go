package generator

import "github.com/steffnova/go-check/constraints"

// Bool returns Arbitrary that create bool Generator.
func Bool() Arbitrary {
	return Int64(constraints.Int64{
		Min: 0,
		Max: 1,
	}).Map(func(n int64) bool {
		return n == 0
	})
}
