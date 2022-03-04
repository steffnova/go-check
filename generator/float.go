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
		negativeMapper := arbitrary.Mapper(reflect.TypeOf(uint64(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(-math.Float64frombits(in.Uint())).Convert(target)
		})
		positiveMapper := arbitrary.Mapper(reflect.TypeOf(uint64(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(math.Float64frombits(in.Uint())).Convert(target)
		})

		switch {
		case constraint.Min < -math.MaxFloat64:
			return nil, fmt.Errorf("lower range value can't be lower then %f", -math.MaxFloat64)
		case constraint.Max > math.MaxFloat64:
			return nil, fmt.Errorf("upper range value can't be greater then %f", math.MaxFloat64)
		case constraint.Max < constraint.Min:
			return nil, fmt.Errorf("lower range value can't be greater then upper range value")
		case constraint.Max <= math.Copysign(0, -1):
			return Uint64(constraints.Uint64{Min: -math.Float64bits(constraint.Max), Max: -math.Float64bits(constraint.Min)}).
				Map(negativeMapper)(target, bias, r)
		case constraint.Min >= math.Copysign(0, 1):
			return Uint64(constraints.Uint64{Min: math.Float64bits(constraint.Min), Max: math.Float64bits(constraint.Max)}).
				Map(positiveMapper)(target, bias, r)
		default:
			return Weighted(
				[]uint64{
					uint64(math.Float64bits(-constraint.Min)) + 1,
					uint64(math.Float64bits(constraint.Max)) + 1,
				},
				Uint64(constraints.Uint64{Min: 0, Max: math.Float64bits(-constraint.Min)}).
					Map(negativeMapper),
				Uint64(constraints.Uint64{Min: 0, Max: math.Float64bits(constraint.Max)}).
					Map(positiveMapper),
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
		mapper := arbitrary.Mapper(reflect.TypeOf(float32(0)), target, func(in reflect.Value) reflect.Value {
			return reflect.ValueOf(float32(in.Float())).Convert(target)
		})
		return Float64(constraints.Float64{
			Min: float64(constraint.Min),
			Max: float64(constraint.Max),
		}).Map(mapper)(target, bias, r)
	}
}
