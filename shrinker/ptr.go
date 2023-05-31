package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func Ptr(original arbitrary.Arbitrary, limit constraints.Ptr) arbitrary.Shrinker {
	if original.Value.Kind() != reflect.Ptr {
		return Fail(fmt.Errorf("Ptr shrinker can't shrink %s", original.Value.Type()))
	}

	predicate := arbitrary.FilterPredicate(original.Value.Type(), func(v reflect.Value) bool {
		return limit.NilFrequency != 0 || !v.IsZero()
	})

	return Chain(
		CollectionSizeRemoveFront(0),
		CollectionOneElement(),
	).
		TransformAfter(arbitrary.NewPtr(original.Value.Type())).
		Filter(predicate)

}
