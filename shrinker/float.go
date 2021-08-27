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
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		switch {
		case val.Kind() != reflect.Float64:
			return reflect.Value{}, nil, fmt.Errorf("float64 shrinker cannot shrink %s", val.Kind().String())
		case limits.Min > limits.Max:
			return reflect.Value{}, nil, fmt.Errorf("limit's min: %f can't be greater than limit's max: %f", limits.Min, limits.Max)
		case val.Float() < limits.Min || val.Float() > limits.Max:
			return reflect.Value{}, nil, fmt.Errorf("n: %v is out of limit constraints: {Min: %f, Max: %f}", val.Float(), limits.Min, limits.Max)
		case val.Float() > 0:
			return float64Positive(val.Float(), val.Type(), limits)(propertyFailed)
		default:
			return float64Negative(-val.Float(), val.Type(), limits)(propertyFailed)
		}
	}
}

// float64Positive is a shrinker for positive float64 numbers. N is the shrinking value, target is a type
// to which shrunk value will be converted to and limits are constraints in which n will be shrunk . N will
// be shrunk towards limits.Min or 0 whichever is higher.
func float64Positive(n float64, target reflect.Type, limits constraints.Float64) Shrinker {
	if limits.Min <= 0 {
		limits.Min = 0
	}
	return Uint64(reflect.ValueOf(math.Float64bits(n)), constraints.Uint64{
		Min: math.Float64bits(limits.Min),
		Max: math.Float64bits(limits.Max),
	}).Map(target, func(floatBits uint64) float64 {
		return math.Float64frombits(floatBits)
	})
}

// float64Negative is a shrinker for negative float64 numbers. N is the shrinking value, target is a type
// to which shrunk value will be converted to and limits are constraints in which n will be shrunk . N will
// be shrunk towards limits.Max or 0 whichever is lower.
func float64Negative(n float64, target reflect.Type, limits constraints.Float64) Shrinker {
	if limits.Max >= 0 {
		limits.Max = math.Copysign(0, -1)
	}
	return Uint64(reflect.ValueOf(math.Float64bits(n)), constraints.Uint64{
		Min: math.Float64bits(-limits.Max),
		Max: math.Float64bits(-limits.Min),
	}).Map(target, func(floatBits uint64) float64 {
		return -math.Float64frombits(floatBits)
	})
}
