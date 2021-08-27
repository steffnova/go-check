package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

// Uint64 is a shrinker for uint64. Val is the shrinking target and limits are constraints in which val
// will be shrunk. Val will be shrunk towards limits.Min or 0, whichever is higher.
//
// An error is returned if val's underlying type is not uint64, limits.Min is greater than limits.Max
// or if val's value is out of limit bounds.
func Uint64(val reflect.Value, limits constraints.Uint64) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		switch {
		case val.Kind() != reflect.Uint64:
			return reflect.Value{}, nil, fmt.Errorf("uint64 shrinker cannot shrink %s", val.Kind().String())
		case limits.Min > limits.Max:
			return reflect.Value{}, nil, fmt.Errorf("lower limit: %d cannot be greater than upper limit: %d", limits.Min, limits.Max)
		case val.Uint() < limits.Min || val.Uint() > limits.Max:
			return reflect.Value{}, nil, fmt.Errorf("n: %v is out of limit constraints: {Min: %v, Max: %v}", val.Uint(), limits.Min, limits.Max)
		case limits.Max == limits.Min:
			return val, nil, nil
		case propertyFailed:
			limits.Max = val.Uint()
		default:
			limits.Min = val.Uint() + 1
		}

		shrink := limits.Max - ((limits.Max-limits.Min)/2 + (limits.Max-limits.Min)%2)
		shrinkVal := reflect.ValueOf(shrink).Convert(val.Type())

		return shrinkVal, Uint64(shrinkVal, limits), nil
	}
}
