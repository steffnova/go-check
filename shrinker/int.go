package shrinker

import (
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

// Int64 is a shrinker for int64. N is the shrinking target and limits are
// constraints in which n will be shrunk. If n > 0 it will be shrunk towards
// limits.Min or 0 whichever is higher. If n < 0 it will be shrunk towards 0
// or limits.Max, whichever is lower. If n == 0 it will be returned with no
// shrinker as it is converging point in ranges [math.Int64Min, 0] and [0, math.Int64Max]
func Int64(n int64, limits constraints.Int64) Shrinker {
	switch {
	case n > 0:
		return int64Positive(n, limits)
	case n < 0:
		return int64Negative(n, limits)
	default:
		return func(propertyFailed bool) (reflect.Value, Shrinker) {
			return reflect.ValueOf(int64(0)), nil
		}
	}

}

// int64Positive is a shrinker of positive int64 numbers. All numbers
// are shrunk towards 0 or limits.Min whichever is higher.
func int64Positive(n int64, limits constraints.Int64) Shrinker {
	if limits.Min < 0 {
		limits.Min = 0
	}
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
		return reflect.ValueOf(shrinked), int64Positive(shrinked, limits)
	}
}

// int64Negative is a shrinker of negative int64 numbers. All numbers
// are shrunk towards 0 or limits.Max whichever is lower.
func int64Negative(n int64, limits constraints.Int64) Shrinker {
	if limits.Max > 0 {
		limits.Max = 0
	}
	return func(propertyFailed bool) (reflect.Value, Shrinker) {
		switch {
		case limits.Max == limits.Min:
			return reflect.ValueOf(n), nil
		case propertyFailed:
			limits.Min = n
		default:
			limits.Max = n - 1
		}

		shrinked := limits.Max - (limits.Max-limits.Min)/2
		return reflect.ValueOf(shrinked), int64Negative(shrinked, limits)
	}
}
