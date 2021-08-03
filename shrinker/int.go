package shrinker

import (
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

// Int64 is a shrinker for int64. X is the shrinking target and limits are
// constraints in which x will be shrunk. If x >= 0 it will be shrunk towards
// limits.Min or 0 whichever is higher. If x < 0 it will be shrunk towards 0
// or limits.Max, whichever is higher.
func Int64(n int64, limits constraints.Int64) Shrinker {
	switch {
	case n >= 0 && limits.Min < 0:
		limits.Min = 0
	case n < 0 && limits.Max > 0:
		limits.Max = 0
	}

	return func(propertyFailed bool) (reflect.Value, Shrinker) {
		switch {
		case limits.Max == limits.Min:
			return reflect.ValueOf(n), nil
		case n >= 0:
			if propertyFailed {
				limits.Max = n
			} else {
				limits.Min = n + 1
			}
			shrinked := limits.Min + (limits.Max-limits.Min)/2
			return reflect.ValueOf(shrinked), Int64(shrinked, limits)
		default:
			if propertyFailed {
				limits.Min = n
			} else {
				limits.Max = n - 1
			}
			shrinked := limits.Max - (limits.Max-limits.Min)/2
			return reflect.ValueOf(shrinked), Int64(shrinked, limits)
		}
	}
}
