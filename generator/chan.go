package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Chan returns Arbitrary that creates chan Generator. Range in which channel's buffer
// size is generated is defined my limits parameter. Even though limits is a variadic
// argument only the first value is used for defining constraints. Channel created by
// Generator is empty and can be used for all 3 types of channel (chan, <-chan and
// chan <-). Error is returned If target's kind is not reflect.Chan.
func Chan(limits ...constraints.Length) Generator {
	constraint := constraints.LengthDefault()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		if target.Kind() != reflect.Chan {
			return nil, fmt.Errorf("target arbitrary's kind must be Chan. Got: %s", target.Kind())
		}
		return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
			return arbitrary.Arbitrary{
				Value: reflect.MakeChan(
					reflect.ChanOf(reflect.BothDir, target.Elem()),
					int(r.Int64(constraints.Int64{
						Min: int64(constraint.Min),
						Max: int64(constraint.Max),
					})),
				),
			}, nil
		}, nil
	}
}
