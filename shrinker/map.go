package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func Map(original arbitrary.Arbitrary, con constraints.Length) arbitrary.Shrinker {
	switch {
	case original.Value.Kind() != reflect.Map:
		return Fail(fmt.Errorf("map shrinker cannot shrink %s", original.Value.Kind().String()))
	case original.Value.Len() != len(original.Elements):
		return Fail(fmt.Errorf("number of map's key-value pairs %d must match size of the map %d", len(original.Elements), original.Value.Len()))
	default:
		for index, element := range original.Elements {
			keyShrinker := element.Elements[0].Shrinker
			valueShrinker := element.Elements[1].Shrinker
			original.Elements[index].Shrinker = Chain(
				CollectionElement(true, keyShrinker, valueShrinker),
				CollectionElements(true, keyShrinker, valueShrinker),
			)
		}

		filter := arbitrary.FilterPredicate(original.Value.Type(), func(in reflect.Value) bool {
			return in.Len() >= int(con.Min)
		})

		return CollectionSize(original.Elements, 0, con).
			Validate(arbitrary.ValidateMap()).
			TransformAfter(arbitrary.NewMap(original.Value.Type())).
			Filter(filter)
	}
}
