package arbitrary

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type stringEncoder func(val reflect.Value) string

func (s stringEncoder) Nil() stringEncoder {
	return func(val reflect.Value) string {
		switch val.Kind() {
		case reflect.Slice, reflect.Chan, reflect.Map, reflect.Ptr:
			if val.IsZero() {
				return "(nil)"
			}
			fallthrough
		default:
			return s(val)
		}
	}
}

func (s stringEncoder) Type() stringEncoder {
	return func(val reflect.Value) string {
		return fmt.Sprintf("<%s> %s", val.Type().String(), s(val))
	}
}

func (s stringEncoder) ArrayOrSlice() stringEncoder {
	return func(val reflect.Value) string {
		switch val.Kind() {
		case reflect.Slice, reflect.Array:
			data := make([]string, val.Len())
			for index := 0; index < val.Len(); index++ {
				data[index] = s(val.Index(index))
			}
			return fmt.Sprintf("[%s]", strings.Join(data, ", "))
		default:
			return s(val)
		}
	}
}

func (s stringEncoder) Map() stringEncoder {
	return func(val reflect.Value) string {
		switch val.Kind() {
		case reflect.Map:
			data := make([]string, val.Len())

			// In order to present Map data in predictable way every time
			// it's keys are sorted by their encoded string value
			keys := val.MapKeys()
			sort.SliceStable(keys, func(i1, i2 int) bool {
				return s(keys[i1]) > s(keys[i2])
			})
			for index, key := range keys {
				data[index] = fmt.Sprintf("%s: %s", s(key), s(val.MapIndex(key)))
			}
			return fmt.Sprintf("{%s}", strings.Join(data, ", "))
		default:
			return s(val)
		}
	}
}

func (s stringEncoder) Struct() stringEncoder {
	return func(val reflect.Value) string {
		switch val.Kind() {
		case reflect.Struct:
			data := make([]string, val.NumField())
			for index := 0; index < val.NumField(); index++ {
				data[index] = fmt.Sprintf("\"%s\": %s", val.Type().Field(index).Name, s.Type()(val.Field(index)))
			}
			return fmt.Sprintf("{%s}", strings.Join(data, ", "))
		default:
			return s(val)
		}
	}
}

func encodeToString() stringEncoder {
	return func(val reflect.Value) string {
		s := encodeToString()
		switch val.Kind() {
		case reflect.Slice, reflect.Array:
			return s.Nil().ArrayOrSlice()(val)
		case reflect.Map:
			return s.Nil().Map()(val)
		case reflect.Struct:
			return s.Struct()(val)
		case reflect.Ptr:
			return s.Nil().Type()(val.Elem())
		case reflect.Func, reflect.Chan:
			return fmt.Sprintf("(%#x)", val.Pointer())
		default:
			return fmt.Sprintf("%v", val.Interface())
		}
	}
}

// EncodeToString encodes val to it's string representation.
func EncodeToString(val reflect.Value) string {
	return encodeToString().Nil().Type()(val)
}
