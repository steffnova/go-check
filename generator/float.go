package generator

import (
	"fmt"
	"math"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Float64 returns generator for float64 types. Range of float64 values that can be generated is
// defined by "limits" parameter. If no limits are provided default float64 range
// [-math.MaxFloat64, math.MaxFloat64] is used instead. Error is returned if generator's target
// is not float64 type, "limits" paramter has invalid values (-Inf, NaN, +Inf), or limits.Min is
// greater than limits.Max.
func Float64(limits ...constraints.Float64) Generator {
	constraint := constraints.Float64Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		mapper := arbitrary.Mapper(reflect.TypeOf(uint64(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(math.Float64frombits(in.Uint())).Convert(target)
		})

		switch {
		case target.Kind() != reflect.Float64:
			return nil, fmt.Errorf("can't use Float64 generator for %s type", target)
		case constraint.Min < -math.MaxFloat64:
			return nil, fmt.Errorf("lower range value can't be lower then %f", -math.MaxFloat64)
		case constraint.Max > math.MaxFloat64:
			return nil, fmt.Errorf("upper range value can't be greater then %f", math.MaxFloat64)
		case constraint.Max < constraint.Min:
			return nil, fmt.Errorf("lower range value can't be greater then upper range value")
		case constraint.Min >= math.Copysign(0, 1):
			return Uint64(
				constraints.Uint64{
					Min: math.Float64bits(constraint.Min),
					Max: math.Float64bits(constraint.Max),
				}).Map(mapper)(target, bias, r)
		case constraint.Max <= math.Copysign(0, -1):
			return Uint64(constraints.Uint64{
				Min: math.Float64bits(math.Copysign(constraint.Max, -1)),
				Max: math.Float64bits(constraint.Min),
			}).Map(mapper)(target, bias, r)
		default:
			return Weighted(
				[]uint64{
					uint64(math.Float64bits(math.Copysign(constraint.Min, 1))) + 1,
					uint64(math.Float64bits(constraint.Max)) + 1,
				},
				Uint64(constraints.Uint64{
					Min: math.Float64bits(math.Copysign(0, -1)),
					Max: math.Float64bits(constraint.Min),
				}).Map(mapper),
				Uint64(constraints.Uint64{
					Min: 0,
					Max: math.Float64bits(constraint.Max),
				}).Map(mapper),
			)(target, bias, r)
		}
	}
}

// Float32 returns generator for float32 types. Range of float64 values that can be generated is
// defined by "limits" paramter. If no constraints are provided default float32 range
// [-math.MaxFloat32, math.MaxFloat32] is used instead. Error is returned if generator's target
// is not float32 type, "limits" paramter has invalid values (-Inf, NaN, +Inf), or limits.Min is
// greater than limits.Max.
func Float32(limits ...constraints.Float32) Generator {
	constraint := constraints.Float32Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}

	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		mapper := arbitrary.Mapper(reflect.TypeOf(uint32(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(math.Float32frombits(uint32(in.Uint()))).Convert(target)
		})

		switch {
		case target.Kind() != reflect.Float32:
			return nil, fmt.Errorf("can't use Float32 generator for %s type", target)
		case constraint.Min < -math.MaxFloat32:
			return nil, fmt.Errorf("lower range value can't be lower then %f", -math.MaxFloat64)
		case constraint.Max > math.MaxFloat32:
			return nil, fmt.Errorf("upper range value can't be greater then %f", math.MaxFloat64)
		case constraint.Max < constraint.Min:
			return nil, fmt.Errorf("lower range value can't be greater then upper range value")
		case constraint.Min >= 0:
			return Uint32(constraints.Uint32{
				Min: math.Float32bits(constraint.Min),
				Max: math.Float32bits(constraint.Max),
			}).Map(mapper)(target, bias, r)
		case constraint.Max <= 0:
			return Uint32(constraints.Uint32{
				Min: math.Float32bits(float32(math.Copysign(float64(constraint.Max), -1))),
				Max: math.Float32bits(constraint.Min),
			}).Map(mapper)(target, bias, r)
		default:
			return Weighted(
				[]uint64{
					uint64(math.Float32bits(-constraint.Min)) + 1,
					uint64(math.Float32bits(constraint.Max)) + 1,
				},
				Uint32(constraints.Uint32{
					Min: math.Float32bits(float32(math.Copysign(0, -1))),
					Max: math.Float32bits(constraint.Min),
				}).Map(mapper),
				Uint32(constraints.Uint32{
					Min: 0,
					Max: math.Float32bits(constraint.Max),
				}).Map(mapper),
			)(target, bias, r)
		}
	}
}
