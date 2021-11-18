package shrinker

import (
	"fmt"
	"reflect"
)

// Array is a shrinker for array. Array is shrinked by shrinking it's elements one at a time.
// An error is returned if arrayType is not an array, length of elements is not equal the
// array size or if any of the array elements return an error during shrinking.
func Array(arrayType reflect.Type, elements []Shrink) Shrinker {
	createArray := func(elements []Shrink) reflect.Value {
		array := reflect.New(arrayType).Elem()
		for index, element := range elements {
			array.Index(index).Set(element.Value)
		}
		return array
	}

	switch {
	case arrayType.Kind() != reflect.Array:
		return Invalid(fmt.Errorf("array shrinker cannot shrink %s", arrayType.Kind().String()))
	case arrayType.Len() != len(elements):
		return Invalid(fmt.Errorf("number of shrinkable elements must match size of the array"))
	default:
		return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
			for index, element := range elements {
				if element.Shrinker == nil {
					continue
				}

				elementValue, elementShrinker, err := element.Shrinker(propertyFailed)
				if err != nil {
					return reflect.Value{}, nil, fmt.Errorf("failed to shrink slice element with index: %d, %s", index, err)
				}

				if !elementValue.Type().AssignableTo(arrayType.Elem()) {
					return reflect.Value{}, nil, fmt.Errorf("failed to assign shrink type: %s to array element with index %d", elementValue.Type(), index)
				}

				elements[index] = Shrink{
					Value:    elementValue,
					Shrinker: elementShrinker,
				}

				return createArray(elements), Array(arrayType, elements), nil
			}
			return createArray(elements), nil, nil
		}
	}
}
