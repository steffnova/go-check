package shrinker

import (
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

// Float64 is a shrinker for float64. N is the shrinking target and limits are constraints
// in which n will be shrunk. If n > 0 it will be shrunk towards limits.Min or 0, whichever
// is higher. If n < 0 it will be shrunk towards 0 or limits.Max, whichever is lower. If
// n == 0 it will be returned with no shrinker as it is converging point in ranges
// [math.Float64Min, 0] and [0, math.Float64Max]
func Float64(n float64, limits constraints.Float64) Shrinker {
	switch {
	case n > 0:
		return float64Positive(n, limits)
	case n < 0:
		return float64Negative(n, limits)
	default:
		return func(propertyFailed bool) (reflect.Value, Shrinker) {
			return reflect.ValueOf(float64(0)), nil
		}
	}
}

// float64Positive is a shrinker for positive float64 numbers. N is the shrinking target and limits are
// constraints in which n will be shrunk. N will be shrunk towards limits.Min or 0 whichever is higher.
func float64Positive(n float64, limits constraints.Float64) Shrinker {
	if limits.Min < 0 {
		limits.Min = 0
	}
	return func(propertyFailed bool) (reflect.Value, Shrinker) {
		switch {
		case propertyFailed:
			limits.Max = n
		default:
			limits.Min = n
		}

		shrinked := limits.Max - (limits.Max-limits.Min)/2
		if shrinked == n {
			return reflect.ValueOf(limits.Max), nil
		}

		return reflect.ValueOf(shrinked), float64Positive(shrinked, limits)
	}
}

// float64Negative is a shrinker for negative float64 numbers. N is the shrinking target and limits are
// constraints in which n will be shrunk. N will be shrunk towards limits.Max or 0 whichever is lower.
func float64Negative(n float64, limits constraints.Float64) Shrinker {
	limits.Max, limits.Min = -limits.Min, -limits.Max
	return float64Positive(-n, limits).Map(func(n float64) float64 {
		return -n
	})
}
