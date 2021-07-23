package arbitrary

import (
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

type KeyValue struct {
	Key   Type
	Value Type
}

type Map struct {
	Constraint constraints.Length
	Key        reflect.Type
	Val        reflect.Type
	Pairs      []KeyValue
}

func (m Map) Shrink() []Type {
	return nil
}

func (m Map) Value() reflect.Value {
	val := reflect.MakeMapWithSize(reflect.MapOf(m.Key, m.Val), len(m.Pairs))
	for _, pair := range m.Pairs {
		val.SetMapIndex(pair.Key.Value(), pair.Value.Value())
	}
	return val
}
