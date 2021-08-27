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
func Slice(target reflect.Type, original []reflect.Value, elementShrinkers []Shrinker, limits constraints.Length) Shrinker {
	candidates := []reflect.Value{}
	candidateIndex := 0
	shrinker := Shrinker(nil)

	shrinker = func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		value := reflect.MakeSlice(target, 0, 0)
		elementsShrunk := true
		var err error
		for index, elementShrinker := range elementShrinkers {
			if elementShrinker == nil {
				continue
			}
			elementsShrunk = false
			original[index], elementShrinkers[index], err = elementShrinker(propertyFailed)
			if err != nil {
				return reflect.Value{}, nil, fmt.Errorf("failed to shrink slice element at index: %d. %w", index, err)
			}
			return reflect.Append(value, original...), Slice(target, original, elementShrinkers, limits), nil
		}

		switch {
		case !elementsShrunk:
			// If elements are not shrunk to smalles possible value, keep shrinking them
			return reflect.Append(value, original...), shrinker, nil
		case !propertyFailed:
			// If element removed from a slice causes the property to succeed, it needs to
			// be added to a candidates as it is essential element for property failure
			candidates = append(candidates, original[candidateIndex-1])
			return reflect.Append(value, append(candidates, original[candidateIndex:]...)...), shrinker, nil
		case limits.Min == len(original)-(candidateIndex-len(candidates)):
			// Shrinking should stop if we've reached a slice's minimal length defined by constraint.
			// The last shrunk value is aggregation of candidates and remaning original elements
			return reflect.Append(value, append(candidates, original[candidateIndex:]...)...), nil, nil
		case candidateIndex == len(original):
			// Shrinking should stop if candidate index has passed through all elements.
			// Candidates is final shrunk value of original slice
			return reflect.Append(value, candidates...), nil, nil
		default:
			// In all other cases property keeps failing so slice size shrinking continues
			// TODO: See if size shrinking speed can be increased
			candidateIndex++
			return reflect.Append(value, append(candidates, original[candidateIndex:]...)...), shrinker, nil
		}
	}
	return shrinker
}
