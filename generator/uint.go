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
func Uint64(limits ...constraints.Uint64) arbitrary.Generator {
	constraint := constraints.Uint64Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		if target.Kind() != reflect.Uint64 {
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "Uint64")
		}
		if constraint.Min > constraint.Max {
			return arbitrary.Arbitrary{}, fmt.Errorf("%w. Lower limit: %d cannot be greater than upper limit: %d", arbitrary.ErrorInvalidConstraints, constraint.Min, constraint.Max)
		}

		n := r.Uint64(constraint)
		nVal := reflect.ValueOf(n).Convert(target)
		return arbitrary.Arbitrary{
			Value:    nVal,
			Shrinker: shrinker.Uint64(constraint),
		}, nil
	}
}

// Uint32 returns generator for uint32 types. Range of int32 values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// uint32 range [0, math.MaxUint32] is used instead. Error is returned if
// generator's target is not uint32 type or limits.Min is greater than limits.Max.
func Uint32(limits ...constraints.Uint32) arbitrary.Generator {
	constraint := constraints.Uint32Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		if target.Kind() != reflect.Uint32 {
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "Uint32")
		}

		mapper := arbitrary.Mapper(reflect.TypeOf(uint64(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(uint32(in.Uint())).Convert(target)
		})
		return Uint64(constraints.Uint64{
			Min: uint64(constraint.Min),
			Max: uint64(constraint.Max),
		}).Map(mapper)(target, bias, r)
	}

}

// Uint16 returns generator for uint16 types. Range of int16 values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// uint16 range [0, math.MaxUint16] is used instead. Error is returned if
// generator's target is not uint16 type or limits.Min is greater than limits.Max.
func Uint16(limits ...constraints.Uint16) arbitrary.Generator {
	constraint := constraints.Uint16Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		if target.Kind() != reflect.Uint16 {
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "Uint16")
		}

		mapper := arbitrary.Mapper(reflect.TypeOf(uint64(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(uint16(in.Uint())).Convert(target)
		})
		return Uint64(constraints.Uint64{
			Min: uint64(constraint.Min),
			Max: uint64(constraint.Max),
		}).Map(mapper)(target, bias, r)
	}
}

// Uint8 returns generator for uint8 types. Range of int8 values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// uint8 range [0, math.MaxUint8] is used instead. Error is returned if
// generator's target is not uint8 type or limits.Min is greater than limits.Max.
func Uint8(limits ...constraints.Uint8) arbitrary.Generator {
	constraint := constraints.Uint8Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		if target.Kind() != reflect.Uint8 {
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "Uint8")
		}
		mapper := arbitrary.Mapper(reflect.TypeOf(uint64(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(uint8(in.Uint())).Convert(target)
		})
		return Uint64(constraints.Uint64{
			Min: uint64(constraint.Min),
			Max: uint64(constraint.Max),
		}).Map(mapper)(target, bias, r)
	}
}

// UInt returns generator for uint types. Range of uint values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// uint range is used instead. Error is returned if generator's target is not uint
// type or limits.Min is greater than limits.Max.
func Uint(limits ...constraints.Uint) arbitrary.Generator {
	constraint := constraints.UintDefault()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		if target.Kind() != reflect.Uint {
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "Uint")
		}

		mapper := arbitrary.Mapper(reflect.TypeOf(uint64(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(uint(in.Uint())).Convert(target)
		})

		return Uint64(constraints.Uint64{
			Min: uint64(constraint.Min),
			Max: uint64(constraint.Max),
		}).Map(mapper)(target, bias, r)
	}
}
