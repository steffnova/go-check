package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

// Slice is a shrinker for slice. Slice is shrinked by two dimensions: elements and size.
// Shrinking is first done by elements, where in each shrink itteration all elements are
// shrunk at the same time. When elements can no longer be srunk, size is being shrunk by
// removing one element at a time. Convergance speed for shrinker is O(n*m), n is slice
// size and m is convergance speed of slice elements.
func Slice(slice SliceShrink, index int, limits constraints.Length) Shrinker {
	switch {
	case slice.Type.Kind() != reflect.Slice:
		return Invalid(fmt.Errorf("slice shrinker cannot shrink: %s", slice.Type.Kind().String()))
	case index < 0 || index > len(slice.Elements):
		return Invalid(fmt.Errorf("index: %d is out of slice range", index))
	case limits.Min == len(slice.Elements)-index:
		return SliceElements(slice)
	default:
		return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
			elements := []Shrink{}
			elements = append(elements, slice.Elements[:index]...)
			elements = append(elements, slice.Elements[index+1:]...)

			nextShrink := SliceShrink{
				Type:     slice.Type,
				Elements: elements,
			}

			shrinker1 := Slice(nextShrink, index, limits)
			shrinker2 := Slice(slice, index+1, limits)

			return nextShrink.Value(), shrinker1.Or(shrinker2), nil
		}
	}
}

func SliceElements(slice SliceShrink) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		for index, element := range slice.Elements {
			if element.Shrinker == nil {
				continue
			}

			elementValue, elementShrinker, err := element.Shrinker(propertyFailed)
			if err != nil {
				return reflect.Value{}, nil, fmt.Errorf("failed to shrink slice element with index: %d, %s", index, err)
			}

			if !elementValue.Type().AssignableTo(slice.Type.Elem()) {
				return reflect.Value{}, nil, fmt.Errorf("failed to assign shrink type: %s to slice element with index %d", elementValue.Type(), index)
			}

			slice.Elements[index] = Shrink{
				Value:    elementValue,
				Shrinker: elementShrinker,
			}

			return slice.Value(), SliceElements(slice), nil
		}
		return slice.Value(), nil, nil
	}
}
