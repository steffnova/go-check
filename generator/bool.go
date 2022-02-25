package generator

import "github.com/steffnova/go-check/constraints"

func Bool() Generator {
	return Uint64(constraints.Uint64{
		Min: 0,
		Max: 1,
	}).Map(func(n int64) bool {
		return n != 0
	})
}
