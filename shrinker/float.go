package shrinker

import (
	"fmt"
	"math"
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

// Float64 is a shrinker for float64. Val is the shrinking target and limits are constraints in which
// val will be shrunk. If val.Float() > 0 it will be shrunk towards limits.Min or 0, whichever is higher.
// Otherwise it will be shrunk towards 0 or limits.Max, whichever is lower.
//
// An error is returned if val's underlying type is not float64, limits.Min is greater than limits.Max
// or if val's value is out of limit bounds.
func Float64(val reflect.Value, limits constraints.Float64) Shrinker {
	switch {
	case val.Kind() != reflect.Float64:
		return Invalid(fmt.Errorf("float64 shrinker cannot shrink %s", val.Kind().String()))
	case limits.Min > limits.Max:
		return Invalid(fmt.Errorf("limit's min: %f can't be greater than limit's max: %f", limits.Min, limits.Max))
	case val.Float() < limits.Min || val.Float() > limits.Max:
		return Invalid(fmt.Errorf("n: %v is out of limit constraints: {Min: %f, Max: %f}", val.Float(), limits.Min, limits.Max))
	case val.Float() > 0:
		if limits.Min <= 0 {
			limits.Min = 0
		}
		return Uint64(reflect.ValueOf(math.Float64bits(val.Float())), constraints.Uint64{
			Min: math.Float64bits(limits.Min),
			Max: math.Float64bits(limits.Max),
		}).Map(func(floatBits uint64) float64 {
			return math.Float64frombits(floatBits)
		}).Convert(val.Type())
	default:
		if limits.Max >= 0 {
			limits.Max = math.Copysign(0, -1)
		}
		return Uint64(reflect.ValueOf(math.Float64bits(-val.Float())), constraints.Uint64{
			Min: math.Float64bits(-limits.Max),
			Max: math.Float64bits(-limits.Min),
		}).Map(func(floatBits uint64) float64 {
			return -math.Float64frombits(floatBits)
		}).Convert(val.Type())
	}
}
