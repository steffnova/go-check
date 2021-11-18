package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Map is a shrinker for map. Map is shrinked by three dimensions: size, keys and values.
// Shrinking is first done by size and afterwords by map's elements. Element's key is shrunk
// before value. Shrinking will fail if mapType is not a map, number of map elements is not
// whithin limits (min and max map's size), or shrinking of one of map elements fails.
func Map(mapType reflect.Type, elements [][2]Shrink, limits constraints.Length) Shrinker {
	switch {
	case mapType.Kind() != reflect.Map:
		return Invalid(fmt.Errorf("map shrinker cannot shrink %s", mapType.String()))
	case limits.Min > len(elements) || limits.Max < len(elements):
		return Invalid(fmt.Errorf("number of map elements: %d is outside of range [%d, %d]", len(elements), limits.Min, limits.Max))
	default:
		mapSliceType := reflect.TypeOf([][2]interface{}{})

		mapper := arbitrary.Mapper(mapSliceType, mapType, func(in reflect.Value) reflect.Value {
			out := reflect.MakeMapWithSize(mapType, in.Len())
			for index := 0; index < in.Len(); index++ {
				key := in.Index(index).Index(0).Interface()
				value := in.Index(index).Index(1).Interface()
				out.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
			}
			return out
		})

		filter := arbitrary.FilterPredicate(mapType, func(in reflect.Value) bool {
			return in.Len() >= limits.Min
		})

		sliceShrink := SliceShrink{
			Type:     mapSliceType,
			Elements: make([]Shrink, len(elements)),
		}

		filterDefault := reflect.MakeMapWithSize(mapType, len(elements))

		for index, element := range elements {
			val := reflect.New(mapSliceType.Elem()).Elem()
			val.Index(0).Set(element[0].Value)
			val.Index(1).Set(element[1].Value)
			sliceShrink.Elements[index] = Shrink{
				Value:    val,
				Shrinker: Array(mapSliceType.Elem(), []Shrink{element[0], element[1]}),
			}

			filterDefault.SetMapIndex(element[0].Value, element[1].Value)
		}

		return Slice(sliceShrink, 0, limits).Map(mapper).Filter(filterDefault, filter)
	}
}
