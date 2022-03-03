package generator

import (
	"fmt"
	"math"

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

	switch {
	case constraint.Min < -math.MaxFloat64:
		return Invalid(fmt.Errorf("lower range value can't be lower then %f", -math.MaxFloat64))
	case constraint.Max > math.MaxFloat64:
		return Invalid(fmt.Errorf("upper range value can't be greater then %f", math.MaxFloat64))
	case constraint.Max < constraint.Min:
		return Invalid(fmt.Errorf("lower range value can't be greater then upper range value"))
	case constraint.Max <= math.Copysign(0, -1):
		return Uint64(constraints.Uint64{Min: -math.Float64bits(constraint.Max), Max: -math.Float64bits(constraint.Min)}).
			Map(func(x uint64) float64 {
				return -math.Float64frombits(x)
			})
	case constraint.Min >= math.Copysign(0, 1):
		return Uint64(constraints.Uint64{Min: math.Float64bits(constraint.Min), Max: math.Float64bits(constraint.Max)}).
			Map(func(x uint64) float64 {
				return math.Float64frombits(x)
			})
	default:
		return Weighted(
			[]uint64{
				uint64(math.Float64bits(-constraint.Min)) + 1,
				uint64(math.Float64bits(constraint.Max)) + 1,
			},
			Uint64(constraints.Uint64{Min: 0, Max: math.Float64bits(-constraint.Min)}).
				Map(func(x uint64) float64 {
					return -math.Float64frombits(x)
				}),
			Uint64(constraints.Uint64{Min: 0, Max: math.Float64bits(constraint.Max)}).
				Map(func(x uint64) float64 {
					return math.Float64frombits(x)
				}),
			// Weighted{
			// 	Weight: uint(math.Float64bits(-constraint.Min)) + 1,

			// },
			// Weighted{
			// 	Weight: uint(math.Float64bits(constraint.Max)) + 1,
			// 	Gen:
			// },
		)
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

	return Float64(constraints.Float64{
		Min: float64(constraint.Min),
		Max: float64(constraint.Max),
	}).Map(func(n float64) float32 {
		return float32(n)
	})
}
