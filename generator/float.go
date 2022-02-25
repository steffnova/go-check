package generator

import (
	"fmt"
	"math"

	"github.com/steffnova/go-check/constraints"
)

// Float64 is Arbitrary that creates float64 Generator. Range in which float64 value is generated
// is defined by limits parameter that specifies range's minimal and maximum value (min and max are
// included in range). If no constraints are provided default range for float64 is used
// [-math.MaxFloat64, math.MaxFloat64]. Even though limits is a variadic argument only the
// first value is used for defining constraints. Error is returned if target's reflect.Kind
// is not Float64 or constraints are out of range (-Inf, +Inf, Nan).
func Float64(limits ...constraints.Float64) Generator {
	constraint := constraints.Float64Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}

	switch {
	case constraint.Min < -math.MaxFloat64:
		return InvalidGen(fmt.Errorf("lower range value can't be lower then %f", -math.MaxFloat64))
	case constraint.Max > math.MaxFloat64:
		return InvalidGen(fmt.Errorf("upper range value can't be greater then %f", math.MaxFloat64))
	case constraint.Max < constraint.Min:
		return InvalidGen(fmt.Errorf("lower range value can't be greater then upper range value"))
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
		return OneFromWeighted(
			Weighted{
				Weight: uint(math.Float64bits(-constraint.Min)) + 1,
				Gen: Uint64(constraints.Uint64{Min: 0, Max: math.Float64bits(-constraint.Min)}).
					Map(func(x uint64) float64 {
						return -math.Float64frombits(x)
					}),
			},
			Weighted{
				Weight: uint(math.Float64bits(constraint.Max)) + 1,
				Gen: Uint64(constraints.Uint64{Min: 0, Max: math.Float64bits(constraint.Max)}).
					Map(func(x uint64) float64 {
						return math.Float64frombits(x)
					}),
			},
		)
	}
}

// Float32 is Arbitrary that creates float32 Generator. Range in which float32 value is generated
// is defined by limits parameter that specifies range's minimal and maximum value (min and max are
// included in range). If no constraints are provided default range for float32 is used
// [-math.MaxFloat32, math.MaxFloat32]. Even though limits is a variadic argument only the
// first value is used for defining constraints. Error is returned if target's reflect.Kind
// is not Float32 or constraints are out of range (-Inf, +Inf, Nan).
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
