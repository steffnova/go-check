package generator

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
)

// Complex128 is generator for complex128 types. Range of complex128 values that can be generated
// is defined by "limits" parameter. If no constraints are provided default range is used
// [-math.MaxFloat64, math.MaxFloat64] for both real and imaginary part of complex128. Error
// is returned if generator's target is not complex128 type or constraints for real or imaginary
// part of complex128 are invalid.
func Complex128(limits ...constraints.Complex128) Generator {
	constraint := constraints.Complex128Default()
	if len(limits) != 0 {
		constraint = limits[0]
	}

	switch {
	case constraint.Real.Min > constraint.Real.Max:
		return Invalid(fmt.Errorf("lower limit of complex's real part can't be higher that it's upper limit"))
	case constraint.Imaginary.Min > constraint.Real.Max:
		return Invalid(fmt.Errorf("lower limit of complex's imaginary part can't be higher that it's upper limit"))
	default:
		return ArrayFrom(
			Float64(constraint.Real),
			Float64(constraint.Imaginary),
		).Map(func(parts [2]float64) complex128 {
			return complex128(complex(parts[0], parts[1]))
		})
	}

}

// Complex64 is generator for complex64 types. Range of complex64 values that can be generated
// is defined by limits parameter. If no constraints are provided default range is used
// [-math.MaxFloat32, math.MaxFloat32] for both real and imaginary part of complex64. Error
// is returned if generator's target is not complex64 type or constraints for real or imaginary
// part of complex64 are invalid.
func Complex64(limits ...constraints.Complex64) Generator {
	constraint := constraints.Complex64Default()
	if len(limits) != 0 {
		constraint = limits[0]
	}

	switch {
	case constraint.Real.Min > constraint.Real.Max:
		return Invalid(fmt.Errorf("lower limit of complex's real part can't be higher that it's upper limit"))
	case constraint.Imaginary.Min > constraint.Real.Max:
		return Invalid(fmt.Errorf("lower limit of complex's imaginary part can't be higher that it's upper limit"))
	default:
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
}
