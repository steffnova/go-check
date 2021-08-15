package shrinker

import (
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

// Uint64 is a shrinker for uint64. N is the shrinking target and limits are
// constraints in which n will be shrunk. N will be shrunk towards limits.Min
// or 0 whichever is higher.
func Uint64(n uint64, limits constraints.Uint64) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker) {
		switch {
		case limits.Max == limits.Min:
			return reflect.ValueOf(n), nil
		case propertyFailed:
			limits.Max = n
		default:
			limits.Min = n + 1
		}

		shrinked := limits.Max - ((limits.Max-limits.Min)/2 + (limits.Max-limits.Min)%2)
		return reflect.ValueOf(shrinked), Uint64(shrinked, limits)
	}
}
