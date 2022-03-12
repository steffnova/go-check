package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Int64 returns generator for int64 types. Range of int64 values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// int64 range [math.MinInt64, math.MaxInt64] is used instead. Error is returned if
// generator's target is not int64 type or limits.Min is greater than limits.Max.
func Int64(limits ...constraints.Int64) Generator {
	constraint := constraints.Int64Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		negativeMapper := arbitrary.Mapper(reflect.TypeOf(uint64(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(int64(-in.Uint())).Convert(target)
		})

		positiveMapper := arbitrary.Mapper(reflect.TypeOf(uint64(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(int64(in.Uint())).Convert(target)
		})

		switch {
		case target.Kind() != reflect.Int64:
			return nil, fmt.Errorf("can't use Int64 generator for %s type", target)
		case constraint.Min > constraint.Max:
			return nil, fmt.Errorf("lower limit: %d cannot be greater than upper limit: %d", constraint.Min, constraint.Max)
		case constraint.Max < 0:
			return Uint64(constraints.Uint64{Min: uint64(-constraint.Max), Max: uint64(-constraint.Min)}).
				Map(negativeMapper)(target, bias, r)
		case constraint.Min >= 0:
			return Uint64(constraints.Uint64{Min: uint64(constraint.Min), Max: uint64(constraint.Max)}).
				Map(positiveMapper)(target, bias, r)
		default:
			return Weighted(
				[]uint64{uint64(-(constraint.Min)), uint64(constraint.Max) + 1},
				Uint64(constraints.Uint64{Min: 0, Max: uint64(-constraint.Min)}).
					Map(negativeMapper),
				Uint64(constraints.Uint64{Min: 0, Max: uint64(constraint.Max)}).
					Map(positiveMapper),
			)(target, bias, r)
		}
	}

}

// Int32 returns generator for int32 types. Range of int32 values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// int64 range [math.MinInt32, math.MaxInt32] is used instead. Error is returned if
// generator's target is not int32 type or limits.Min is greater than limits.Max.
func Int32(limits ...constraints.Int32) Generator {
	constraint := constraints.Int32Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		if target.Kind() != reflect.Int32 {
			return nil, fmt.Errorf("can't use Int32 generator for %s type", target)
		}

		mapper := arbitrary.Mapper(reflect.TypeOf(int64(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(int32(in.Int())).Convert(target)
		})
		return Int64(constraints.Int64{
			Min: int64(constraint.Min),
			Max: int64(constraint.Max),
		}).Map(mapper)(target, bias, r)
	}
}

// Int16 returns generator for int16 types. Range of int16 values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// int16 range [math.MinInt16, math.MaxInt16] is used instead. Error is returned if
// generator's target is not int16 type or limits.Min is greater than limits.Max.
func Int16(limits ...constraints.Int16) Generator {
	constraint := constraints.Int16Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		if target.Kind() != reflect.Int16 {
			return nil, fmt.Errorf("can't use Int16 generator for %s type", target)
		}

		mapper := arbitrary.Mapper(reflect.TypeOf(int64(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(int16(in.Int())).Convert(target)
		})
		return Int64(constraints.Int64{
			Min: int64(constraint.Min),
			Max: int64(constraint.Max),
		}).Map(mapper)(target, bias, r)
	}
}

// Int8 returns generator for int8 types. Range of int8 values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// int8 range [math.MinInt8, math.MaxInt8] is used instead. Error is returned if
// generator's target is not int8 type or limits.Min is greater than limits.Max.
func Int8(limits ...constraints.Int8) Generator {
	constraint := constraints.Int8Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		if target.Kind() != reflect.Int8 {
			return nil, fmt.Errorf("can't use Int8 generator for %s type", target)
		}

		mapper := arbitrary.Mapper(reflect.TypeOf(int64(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(int8(in.Int())).Convert(target)
		})
		return Int64(constraints.Int64{
			Min: int64(constraint.Min),
			Max: int64(constraint.Max),
		}).Map(mapper)(target, bias, r)
	}
}

// Int returns generator for int types. Range of int values that can be
// generated is defined by "limits" parameter.  If no limits are provided default
// int range is used instead. Error is returned if generator's target is not int
// type or limits.Min is greater than limits.Max.
func Int(limits ...constraints.Int) Generator {
	constraint := constraints.IntDefault()
	if len(limits) > 0 {
		constraint.Min, constraint.Max = limits[0].Min, limits[0].Max
	}

	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		if target.Kind() != reflect.Int {
			return nil, fmt.Errorf("can't use Int generator for %s type", target)
		}

		mapper := arbitrary.Mapper(reflect.TypeOf(int64(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(int(in.Int())).Convert(target)
		})
		return Int64(constraints.Int64{
			Min: int64(constraint.Min),
			Max: int64(constraint.Max),
		}).Map(mapper)(target, bias, r)
	}
}
