package generator

import (
	"github.com/steffnova/go-check/constraints"
)

// Complex128 is Arbitrary that creates complex128 Generator. Range in which complex128 value is
// generated is defined by limits parameter that specifies range in which both real and imaginary
// part of complex number are generated. Range is a float64 range with minimal and maximum value
// (min and max are included in range). If no constraints are provided default range for float64
// is used [-math.MaxFloat64, math.MaxFloat64] for both parts. Even though limits is a variadic
// argument only the first value is used for defining constraints. Error is returned if target's
// reflect.Kind is not Complex128 or constraints are out of range (-Inf, +Inf, Nan).
func Complex128(limits ...constraints.Complex128) Generator {
	constraint := constraints.Complex128Default()
	if len(limits) != 0 {
		constraint = limits[0]
	}

	return ArrayFrom(
		Float64(constraint.Real),
		Float64(constraint.Imaginary),
	).Map(func(parts [2]float64) complex128 {
		return complex128(complex(parts[0], parts[1]))
	})
}

// Complex64 is Arbitrary that creates complex64 Generator. Range in which complex64 value is
// generated is defined by limits parameter that specifies range in which both real and imaginary
// part of complex number are generated. Range is a float32 range with minimal and maximum value
// (min and max are included in range). If no constraints are provided default range for float64
// is used [-math.MaxFloat32, math.MaxFloat32] for both parts. Even though limits is a variadic
// argument only the first value is used for defining constraints. Error is returned if target's
// reflect.Kind is not Complex64 or constraints are out of range (-Inf, +Inf, Nan).
func Complex64(limits ...constraints.Complex64) Generator {
	constraint := constraints.Complex64Default()
	if len(limits) != 0 {
		constraint = limits[0]
	}

	return Complex128(constraints.Complex128{
		Real: constraints.Float64{
			Min: float64(constraint.Real.Min),
			Max: float64(constraint.Real.Max),
		},
		Imaginary: constraints.Float64{
			Min: float64(constraint.Imaginary.Min),
			Max: float64(constraint.Imaginary.Max),
		},
	}).Map(func(n complex128) complex64 {
		return complex64(n)
	})
}
