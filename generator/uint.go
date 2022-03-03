package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Uint64 returns generator for uint64 types. Range of int64 values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// uint64 range [0, math.MaxUint64] is used instead. Error is returned if
// generator's target is not uint64 type or limits.Min is greater than limits.Max.
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

// Uint32 returns generator for uint32 types. Range of int32 values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// uint32 range [0, math.MaxUint32] is used instead. Error is returned if
// generator's target is not uint32 type or limits.Min is greater than limits.Max.
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

// Uint16 returns generator for uint16 types. Range of int16 values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// uint16 range [0, math.MaxUint16] is used instead. Error is returned if
// generator's target is not uint16 type or limits.Min is greater than limits.Max.
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

// Uint8 returns generator for uint8 types. Range of int8 values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// uint8 range [0, math.MaxUint8] is used instead. Error is returned if
// generator's target is not uint8 type or limits.Min is greater than limits.Max.
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

// UInt returns generator for uint types. Range of uint values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// uint range is used instead. Error is returned if generator's target is not uint
// type or limits.Min is greater than limits.Max.
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
