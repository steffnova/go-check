package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func Slice(original arbitrary.Arbitrary, shrinkers []Shrinker, con constraints.Length) Shrinker {
	switch {
	case original.Value.Kind() != reflect.Slice:
		return Fail(fmt.Errorf("slice shrinker cannot shrink %s", original.Value.Kind().String()))
	case original.Value.Len() != len(original.Elements):
		return Fail(fmt.Errorf("number of elements %d must match size of the array %d", len(original.Elements), original.Value.Len()))
	default:
		return CollectionSize(original.Elements, shrinkers, 0, con).
			Validate(arbitrary.ValidateSlice()).
			Transform(arbitrary.NewSlice(original.Value.Type()))
	}
}
