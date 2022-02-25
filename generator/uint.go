package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Uint64 is Generator that creates uint64 Generator. Range in which uint64 value is generated
// is defined by limits parameter that specifies range's minimal and maximum value (min and max
// are included in range). If no constraints are provided default range for uint64 is used
// [math.MinUint64, math.MaxUint64]. Even though limits is a variadic argument only the first
// value is used for defining constraints. Error is returned if target's reflect.Kind is not Uint64.
func Uint64(limits ...constraints.Uint64) Generator {
	constraint := constraints.Uint64Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		if target.Kind() != reflect.Uint64 {
			return nil, fmt.Errorf("target arbitrary's kind must be Uint64. Got: %s", target.Kind())
		}
		return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
			n := r.Uint64(constraint)
			nVal := reflect.ValueOf(n).Convert(target)
			return arbitrary.Arbitrary{
				Value: nVal,
			}, shrinker.Uint64(constraint)
		}, nil
	}
}

func Uint32(limits ...constraints.Uint32) Generator {
	constraint := constraints.Uint32Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return Uint64(constraints.Uint64{
		Min: uint64(constraint.Min),
		Max: uint64(constraint.Max),
	}).Map(func(n uint64) uint32 {
		return uint32(n)
	})
}

func Uint16(limits ...constraints.Uint16) Generator {
	constraint := constraints.Uint16Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return Uint64(constraints.Uint64{
		Max: uint64(constraint.Max),
		Min: uint64(constraint.Min),
	}).Map(func(n uint64) uint16 {
		return uint16(n)
	})
}

// Uint8 is Arbitrary that creates uint8 Generator. Range in which Uint8 value is generated
// is defined by limits parameter that specifies range's minimal and maximum value (min and
// max are included in range). If no constraints are provided default range for Uint8 is
// used [math.MinUint8, math.MaxUint8]. Even though limits is a variadic argument only the
// first value is used for defining constraints. Error is returned if target's reflect.Kind
// is not Uint8.
func Uint8(limits ...constraints.Uint8) Generator {
	constraint := constraints.Uint8Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return Uint64(constraints.Uint64{
		Max: uint64(constraint.Max),
		Min: uint64(constraint.Min),
	}).Map(func(n uint64) uint8 {
		return uint8(n)
	})
}

// Uint is Arbitrary that creates uint Generator. Range in which Uint value is generated
// is defined by limits parameter that specifies range's minimal and maximum value (min and
// max are included in range). If no constraints are provided default range for Uint is
// used [math.MinUint, math.MaxUint]. Even though limits is a variadic argument only the
// first value is used for defining constraints. Error is returned if target's reflect.Kind
// is not Uint.
func Uint(limits ...constraints.Uint) Generator {
	constraint := constraints.UintDefault()
	if len(limits) > 0 {
		constraint = limits[0]
	}

	return Uint64(constraints.Uint64{
		Max: uint64(constraint.Max),
		Min: uint64(constraint.Min),
	}).Map(func(n uint64) uint {
		return uint(n)
	})
}
