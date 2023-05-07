package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func Uint64(limits constraints.Uint64) arbitrary.Shrinker {
	return func(val arbitrary.Arbitrary, propertyFailed bool) (arbitrary.Arbitrary, error) {
		switch {
		case val.Value.Kind() != reflect.Uint64:
			return arbitrary.Arbitrary{}, fmt.Errorf("uint64 shrinker cannot shrink %s", val.Value.Kind().String())
		case limits.Min > limits.Max:
			return arbitrary.Arbitrary{}, fmt.Errorf("lower limit: %d cannot be greater than upper limit: %d", limits.Min, limits.Max)
		case val.Value.Uint() < limits.Min || val.Value.Uint() > limits.Max:
			return arbitrary.Arbitrary{}, fmt.Errorf("n: %v is out of limit constraints: {Min: %v, Max: %v}", val.Value.Uint(), limits.Min, limits.Max)
		case limits.Max == limits.Min:
			return val, nil
		case propertyFailed:
			limits.Max = val.Value.Uint()
		default:
			limits.Min = val.Value.Uint() + 1
		}

		shrink := limits.Max - ((limits.Max-limits.Min)/2 + (limits.Max-limits.Min)%2)

		return arbitrary.Arbitrary{
			Value:    reflect.ValueOf(shrink).Convert(val.Value.Type()),
			Shrinker: Uint64(limits),
		}, nil
	}
}
