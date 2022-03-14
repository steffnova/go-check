package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
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
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		switch {
		case target.Kind() != reflect.Complex128:
			return nil, fmt.Errorf("can't use Complex128 generator for %s type", target)
		case constraint.Real.Min > constraint.Real.Max:
			return nil, fmt.Errorf("lower limit of complex's real part can't be higher that it's upper limit")
		case constraint.Imaginary.Min > constraint.Imaginary.Max:
			return nil, fmt.Errorf("lower limit of complex's imaginary part can't be higher that it's upper limit")
		default:
			mapper := arbitrary.Mapper(reflect.TypeOf([2]float64{}), target, func(in reflect.Value) reflect.Value {
				parts := in.Interface().([2]float64)
				return reflect.ValueOf(complex(parts[0], parts[1])).Convert(target)
			})
			return ArrayFrom(
				Float64(constraint.Real),
				Float64(constraint.Imaginary),
			).Map(mapper)(target, bias, r)
		}
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
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		switch {
		case target.Kind() != reflect.Complex64:
			return nil, fmt.Errorf("can't use Complex64 generator for %s type", target)
		case constraint.Real.Min > constraint.Real.Max:
			return nil, fmt.Errorf("lower limit of complex's real part can't be higher that it's upper limit")
		case constraint.Imaginary.Min > constraint.Imaginary.Max:
			return nil, fmt.Errorf("lower limit of complex's imaginary part can't be higher that it's upper limit")
		default:
			mapper := arbitrary.Mapper(reflect.TypeOf([2]float32{}), target, func(in reflect.Value) reflect.Value {
				parts := in.Interface().([2]float32)
				return reflect.ValueOf(complex(parts[0], parts[1])).Convert(target)
			})
			return ArrayFrom(
				Float32(constraint.Real),
				Float32(constraint.Imaginary),
			).Map(mapper)(target, bias, r)
		}
	}
}
