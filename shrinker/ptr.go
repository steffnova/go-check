package shrinker

import "reflect"

// Ptr is a shrinker for pointers. ElementShrinker is shrinker for type to which
// ptr points to. Shrinking process consists of shrinking the value to which pointer
// points to, and shrinking of pointer to nil.
func Ptr(ptr reflect.Value, elementShrinker Shrinker) Shrinker {
	var shrinker Shrinker
	var lastValidPtr reflect.Value

	shrinker = func(propertyFailed bool) (reflect.Value, Shrinker) {
		var val reflect.Value
		if elementShrinker != nil {
			val, elementShrinker = elementShrinker(propertyFailed)
			ptr.Elem().Set(val)
			return ptr, shrinker
		}

		switch {
		case !ptr.IsNil():
			// Shrinking to nil could make a property not fail. The last valid pointer
			// value needs to be saved in order to revert nil pointer if needed.
			lastValidPtr, ptr = ptr, reflect.Zero(ptr.Type())
			return ptr, shrinker
		case propertyFailed:
			return ptr, nil
		default:
			return lastValidPtr, nil
		}
	}

	return shrinker
}
