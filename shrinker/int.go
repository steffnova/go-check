package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

// Int64 is a shrinker for int64. Val is the shrinking target and limits are constraints in which
// val will be shrunk. If val.Int() it will be shrunk towards limits.Min or 0, whichever is higher.
// Otherwise it will be shrunk towards 0 or limits.Max, whichever is lower.
//
// An error is returned if val's underlying type is not int64, limits.Min is greater than limits.Max or
// if val's value is out of limit bounds.
func Int64(val reflect.Value, limits constraints.Int64) Shrinker {
	switch {
	case val.Kind() != reflect.Int64:
		return Invalid(fmt.Errorf("int64 shrinker cannot shrink %s", val.Kind().String()))
	case limits.Min > limits.Max:
		return Invalid(fmt.Errorf("lower limit: %d cannot be greater than upper limit: %d", limits.Min, limits.Max))
	case val.Int() < limits.Min || val.Int() > limits.Max:
		return Invalid(fmt.Errorf("n: %v is out of limit constraints: {Min: %v, Max: %v}", val.Int(), limits.Min, limits.Max))
	case val.Int() > 0:
		if limits.Min < 0 {
			limits.Min = 0
		}
		return Uint64(reflect.ValueOf(uint64(val.Int())), constraints.Uint64{
			Min: uint64(limits.Min),
			Max: uint64(limits.Max),
		}).Map(func(n uint64) int64 {
			return int64(n)
		}).Convert(val.Type())
	default:
		if limits.Max > 0 {
			limits.Max = 0
		}
		return Uint64(reflect.ValueOf(uint64(-val.Int())), constraints.Uint64{
			Min: uint64(-limits.Max),
			Max: uint64(-limits.Min),
		}).Map(func(n uint64) int64 {
			return -int64(n)
		}).Convert(val.Type())
	}
}
