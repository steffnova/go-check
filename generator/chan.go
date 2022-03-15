package generator

import (
	"fmt"
	"math"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Chan returns generator of chan types. Channel's buffer size range is defined with "limits"
// parameter. Generated channel is empty and can be used for all 3 channel types (chan, <-chan and
// chan <-). Error is returned if generator's target is not chan type or limit's lower bound i higher
// than limit's upper bound.
func Chan(limits ...constraints.Length) Generator {
	constraint := constraints.LengthDefault()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		if target.Kind() != reflect.Chan {
			return nil, fmt.Errorf("can't use Chan generator for %s type", target)
		}
		if constraint.Min > constraint.Max {
			return nil, fmt.Errorf("minimal length value %d can't be greater than max length value %d", constraint.Min, constraint.Max)
		}
		if constraint.Max > uint64(math.MaxInt64) {
			return nil, fmt.Errorf("max length %d can't be greater than %d", constraint.Max, uint64(math.MaxInt64))
		}
		return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
			return arbitrary.Arbitrary{
				Value: reflect.MakeChan(
					reflect.ChanOf(reflect.BothDir, target.Elem()),
					int(r.Uint64(constraints.Uint64{
						Min: uint64(constraint.Min),
						Max: uint64(constraint.Max),
					})),
				),
			}, nil
		}, nil
	}
}
