package generator

import (
	"fmt"
	"math"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Chan returns generator of chan types. Channel's buffer size range is defined with "limits"
// parameter. arbitrary.Arbitraryd channel is empty and can be used for all 3 channel types (chan, <-chan and
// chan <-). Error is returned if generator's target is not chan type or limit's lower bound i higher
// than limit's upper bound.
func Chan(limits ...constraints.Length) arbitrary.Generator {
	constraint := constraints.LengthDefault()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		if target.Kind() != reflect.Chan {
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "Chan")
		}
		if constraint.Min > constraint.Max {
			return arbitrary.Arbitrary{}, fmt.Errorf("%w. Minimal length value %d can't be greater than max length value %d", arbitrary.ErrorInvalidConstraints, constraint.Min, constraint.Max)
		}
		if constraint.Max > uint64(math.MaxInt64) {
			return arbitrary.Arbitrary{}, fmt.Errorf("%w. Max length %d can't be greater than %d", arbitrary.ErrorInvalidConstraints, constraint.Max, uint64(math.MaxInt64))
		}

		return arbitrary.Arbitrary{
			Value: reflect.MakeChan(
				reflect.ChanOf(reflect.BothDir, target.Elem()),
				int(r.Uint64(constraints.Uint64{
					Min: uint64(constraint.Min),
					Max: uint64(constraint.Max),
				})),
			),
		}, nil
	}
}
