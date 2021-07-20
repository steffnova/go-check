package generator

import "github.com/steffnova/go-check/constraints"

// Bool returns Arbitrary generator that can be used to create Bool generator
func Bool() Arbitrary {
	return Int64(constraints.Int64{
		Min: 0,
		Max: 1,
	}).Map(func(n int64) (bool, error) {
		return n == 0, nil
	})
}
