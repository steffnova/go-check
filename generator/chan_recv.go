package generator

import (
	"fmt"
	"math"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// [ChanRecv] is a generator of read-only channels. The generated channel is closed, buffered,
// and has a capacity that falls within the range specified by [constraints.Length]. It is filled
// with values of the channel's element type. The 'gen' parameter [arbitrary.Generator] is used
// to fill the channel with values and must match the type of the channel's elements. Although the
// 'limits' parameter is variadic, only the first value passed is used. If the 'limits' parameter is
// omitted, [constraints.LengthDefault] is used instead. The following example demonstrates how to use
// ChanRecv to generate a <-chan int that can have between 10 and 20 int elements within it:
//
//	generator.ChanRecv(generator.Int(), constraints.Length{Min: 10, Max: 20})
//
// ChanRecv will return an error:
//   - If the generator's target is not <-chan T
//   - If the 'limits' parameter is invalid: [constraints.Length.Min] > [constraints.Length.Max], [constraints.Length.Max] > [math.MaxInt64]
func ChanRecv(element arbitrary.Generator, limits ...constraints.Length) arbitrary.Generator {
	constraint := constraints.LengthDefault()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		if target.Kind() != reflect.Chan || target.ChanDir() != reflect.RecvDir {
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "ChanReadOnly")
		}

		if constraint.Min > constraint.Max {
			return arbitrary.Arbitrary{}, fmt.Errorf("%w. Minimal length value %d can't be greater than max length value %d", arbitrary.ErrorInvalidConstraints, constraint.Min, constraint.Max)
		}
		if constraint.Max > uint64(math.MaxInt64) {
			return arbitrary.Arbitrary{}, fmt.Errorf("%w. Max length %d can't be greater than %d", arbitrary.ErrorInvalidConstraints, constraint.Max, uint64(math.MaxInt64))
		}

		mapper := arbitrary.Mapper(reflect.SliceOf(target.Elem()), reflect.ChanOf(reflect.RecvDir, target.Elem()), func(v reflect.Value) reflect.Value {
			val := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, target.Elem()), v.Len())
			for i := 0; i < v.Len(); i++ {
				val.Send(v.Index(i))
			}
			val.Close()
			return val
		})

		return Slice(element, constraint).Map(mapper)(target, bias, r)
	}
}
