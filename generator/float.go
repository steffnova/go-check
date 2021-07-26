package generator

import (
	"fmt"
	"math"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Float64 is Arbitrary that creates float64 Generator. Range in which float64 value is generated
// is defined by limits parameter that specifies range's minimal and maximum value (min and max are
// included in range). If no constraints are provided default range for float64 is used
// [-math.MaxFloat64, math.MaxFloat64]. Even though limits is a variadic argument only the
// first value is used for defining constraints. Error is returned if target's reflect.Kind
// is not Float64 or constraints are out of range (-Inf, +Inf, Nan).
func Float64(limits ...constraints.Float64) Arbitrary {
	constraint := constraints.Float64Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}

	return func(target reflect.Type, r Random) (Generator, error) {
		if target.Kind() != reflect.Float64 {
			return nil, fmt.Errorf("target arbitrary's kind must be Int64. Got: %s", target.Kind())
		}
		if constraint.Min < -math.MaxFloat64 {
			return nil, fmt.Errorf("lower range value can't be lower then %f", -math.MaxFloat64)
		}
		if constraint.Max > math.MaxFloat64 {
			return nil, fmt.Errorf("upper range value can't be greater then %f", math.MaxFloat64)
		}
		if constraint.Max < constraint.Min {
			return nil, fmt.Errorf("lower range value can't be greater then upper range value")
		}

		return func() arbitrary.Type {
			return arbitrary.Float64{
				Constraint: constraint,
				N:          r.Float64(constraint.Min, constraint.Max),
			}
		}, nil
	}
}

// Float32 is Arbitrary that creates float32 Generator. Range in which float32 value is generated
// is defined by limits parameter that specifies range's minimal and maximum value (min and max are
// included in range). If no constraints are provided default range for float32 is used
// [-math.MaxFloat32, math.MaxFloat32]. Even though limits is a variadic argument only the
// first value is used for defining constraints. Error is returned if target's reflect.Kind
// is not Float32 or constraints are out of range (-Inf, +Inf, Nan).
func Float32(limits ...constraints.Float32) Arbitrary {
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
