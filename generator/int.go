package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Int64 is Arbitrary that creates int64 Generator. Range in which int64 value is generated
// is defined by limits parameter that specifies range's minimal and maximum value (min and
// max are included in range). If no constraints are provided default range for int64 is
// used [math.MinInt64, math.MaxInt64]. Even though limits is a variadic argument only the
// first value is used for defining constraints. Error is returned if target's reflect.Kind
// is not Int64.
func Int64(limits ...constraints.Int64) Arbitrary {
	constraint := constraints.Int64Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return func(target reflect.Type, r Random) (Generator, error) {
		if target.Kind() != reflect.Int64 {
			return nil, fmt.Errorf("target arbitrary's kind must be Int64. Got: %s", target.Kind())
		}
		return func() (reflect.Value, shrinker.Shrinker) {
			n := r.Int64(constraint.Min, constraint.Max)
			nVal := reflect.ValueOf(n).Convert(target)
			return nVal, shrinker.Int64(nVal, constraint)
		}, nil
	}
}

// Int32 is Arbitrary that creates int32 Generator. Range in which int32 value is generated
// is defined by limits parameter that specifies range's minimal and maximum value (min and
// max are included in range). If no constraints are provided default range for int32 is
// used [math.MinInt32, math.MaxInt32]. Even though limits is a variadic argument only the
// first value is used for defining constraints. Error is returned if target's reflect.Kind
// is not Int32.
func Int32(limits ...constraints.Int32) Arbitrary {
	constraint := constraints.Int32Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return Int64(constraints.Int64{
		Max: int64(constraint.Max),
		Min: int64(constraint.Min),
	}).Map(func(n int64) int32 {
		return int32(n)
	})
}

// Int16 is Arbitrary that creates int16 Generator. Range in which int16 value is generated
// is defined by limits parameter that specifies range's minimal and maximum value (min and
// max are included in range). If no constraints are provided default range for int16 is
// used [math.MinInt16, math.MaxInt16]. Even though limits is a variadic argument only the
// first value is used for defining constraints. Error is returned if target's reflect.Kind
// is not Int16.
func Int16(limits ...constraints.Int16) Arbitrary {
	constraint := constraints.Int16Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return Int64(constraints.Int64{
		Max: int64(constraint.Max),
		Min: int64(constraint.Min),
	}).Map(func(n int64) int16 {
		return int16(n)
	})
}

// Int8 is Arbitrary that creates int16 Generator. Range in which int8 value is generated
// is defined by limits parameter that specifies range's minimal and maximum value (min and
// max are included in range). If no constraints are provided default range for int8 is
// used [math.MinInt8, math.MaxInt8]. Even though limits is a variadic argument only the
// first value is used for defining constraints. Error is returned if target's reflect.Kind
// is not Int8.
func Int8(limits ...constraints.Int8) Arbitrary {
	constraint := constraints.Int8Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return Int64(constraints.Int64{
		Max: int64(constraint.Max),
		Min: int64(constraint.Min),
	}).Map(func(n int64) int8 {
		return int8(n)
	})
}

// Int is Arbitrary that creates int Generator. Range in which int value is generated is
// defined by limits parameter that specifies range's minimal and maximum value (min and
// max are included in range). If no constraints are provided default range for int is used
// [math.MinInt32, math.MaxInt32] for 32bit architecture or [math.MinInt64, math.MaxInt64]
// for 64bit architecture. Even though limits is a variadic argument first value is used for
// defining constraints. Error is returned if target's reflect.Kind is not Int.
func Int(limits ...constraints.Int) Arbitrary {
	constraint := constraints.IntDefault()
	if len(limits) > 0 {
		constraint.Min, constraint.Max = limits[0].Min, limits[0].Max
	}

	return Int64(constraints.Int64{
		Max: int64(constraint.Max),
		Min: int64(constraint.Min),
	}).Map(func(n int64) int {
		return int(n)
	})
}
