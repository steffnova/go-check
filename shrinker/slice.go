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
func Slice(val reflect.Value, shrinkers []Shrinker, limits constraints.Length) Shrinker {
	return sliceElements(val, shrinkers...).Compose(sliceSize(val, 0, limits))
}

func sliceElement(val reflect.Value, index int, shrinker Shrinker) Shrinker {
	if shrinker == nil {
		return nil
	}

	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		switch {
		case val.Kind() != reflect.Slice:
			return reflect.Value{}, nil, fmt.Errorf("slice element shrinker cannot shrink: %s", val.Kind().String())
		case val.Len() <= index:
			return reflect.Value{}, nil, fmt.Errorf("cannot shrink element with index: %d out of range", index)
		}

		elementVal, shrinker, err := shrinker(propertyFailed)
		if err != nil {
			return reflect.Value{}, nil, fmt.Errorf("failed to shrink slice element: %w", err)
		}

		val.Index(index).Set(elementVal)
		return val, sliceElement(val, index, shrinker), nil
	}
}

func sliceElements(val reflect.Value, shrinkers ...Shrinker) Shrinker {
	var chainShrinker Shrinker
	for index, shrinker := range shrinkers {
		chainShrinker = chainShrinker.Compose(sliceElement(val, index, shrinker))
	}
	return chainShrinker
}

func sliceSize(val reflect.Value, index int, limits constraints.Length) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		switch {
		case val.Kind() != reflect.Slice:
			return reflect.Value{}, nil, fmt.Errorf("slice size shrinker cannot shrink: %s", val.Kind().String())
		case index < 0 || index > val.Len():
			return reflect.Value{}, nil, fmt.Errorf("index: %d is out of slice range", index)
		case limits.Min == val.Len()-index:
			return val, nil, nil
		default:
			shrink := reflect.MakeSlice(val.Type(), 0, val.Len()-1)
			shrink = reflect.AppendSlice(shrink, val.Slice(0, index))
			shrink = reflect.AppendSlice(shrink, val.Slice(index+1, val.Len()))

			shrinker1 := sliceSize(shrink, index, limits)
			shrinker2 := sliceSize(val, index+1, limits)

			return shrink, shrinker1.WithFallback(shrinker2), nil
		}
	}
}
