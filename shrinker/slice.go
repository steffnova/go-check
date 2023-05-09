package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func Slice(original arbitrary.Arbitrary, con constraints.Length) arbitrary.Shrinker {
	switch {
	case original.Value.Kind() != reflect.Slice:
		return Fail(fmt.Errorf("slice shrinker cannot shrink %s", original.Value.Kind().String()))
	case original.Value.Len() != len(original.Elements):
		return Fail(fmt.Errorf("number of elements %d must match size of the array %d", len(original.Elements), original.Value.Len()))
	default:
		filter := arbitrary.FilterPredicate(original.Value.Type(), func(in reflect.Value) bool {
			return in.Len() >= int(con.Min) && in.Len() <= int(con.Max)
		})

		return CollectionRemoveFront(0).
			Validate(arbitrary.ValidateSlice()).
			TransformAfter(arbitrary.NewSlice(original.Value.Type())).
			Filter(filter)
	}
}
