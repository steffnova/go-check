package shrinker

import (
	"fmt"
	"reflect"
)

// Struct is a shrinker for Go's struct. Struct shrinking consists of shrinking individual
// fields one by one. Convergance speed for shrinker is O(n*m), n is number of Struct fields
// and m is convergance speed of field type.
func Struct(val reflect.Value, fieldShrinkers map[string]Shrinker) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		if val.Kind() != reflect.Struct {
			return reflect.Value{}, nil, fmt.Errorf("struct shrinker cannot shrink: %s", val.Kind().String())
		}

		for i := 0; i < val.Type().NumField(); i++ {
			fieldName := val.Type().Field(i).Name

			shrinker := fieldShrinkers[fieldName]
			if shrinker == nil {
				continue
			}

			fieldVal, shrinker, err := shrinker(propertyFailed)
			if err != nil {
				return reflect.Value{}, nil, fmt.Errorf("failed to shrink struct field: %s. %w", fieldName, err)
			}

			val.FieldByName(fieldName).Set(fieldVal)
			fieldShrinkers[fieldName] = shrinker

			return val, Struct(val, fieldShrinkers), nil
		}

		return val, nil, nil
	}
}
