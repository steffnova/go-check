package arbitrary

import (
	"reflect"
)

type StructField struct {
	Name string
	Type Type
}

type Struct struct {
	Fields map[string]Type
	Type   reflect.Type
}

func (s Struct) Shrink() []Type {
	return nil
}

func (s Struct) Value() reflect.Value {
	v := reflect.New(s.Type).Elem()
	for name, t := range s.Fields {
		v.FieldByName(name).Set(t.Value())
	}
	return v
}
