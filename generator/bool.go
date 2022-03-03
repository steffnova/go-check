package generator

import "github.com/steffnova/go-check/constraints"

// Bool returns generator of bool types. Error is returned if generator's
// target is not bool type.
func Bool() Generator {
	return Uint64(constraints.Uint64{
		Min: 0,
		Max: 1,
	}).Map(func(n int64) bool {
		return n != 0
	})
}
