package shrinker

import (
	"reflect"
)

// Struct is a shrinker for Go's struct. Struct shrinking consists of shrinking individual
// fields one by one. Convergance speed for shrinker is O(n*m), n is number of Struct fields
// and m is convergance speed of field type.
func Struct(original reflect.Value, fieldShrinkers map[string]Shrinker) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker) {
		for i := 0; i < original.Type().NumField(); i++ {
			fieldName := original.Type().Field(i).Name

			shrinker := fieldShrinkers[fieldName]
			if shrinker == nil {
				continue
			}

			val, shrinker := shrinker(propertyFailed)

			original.FieldByName(fieldName).Set(val)
			fieldShrinkers[fieldName] = shrinker

			return original, Struct(original, fieldShrinkers)
		}

		return original, nil
	}
}
