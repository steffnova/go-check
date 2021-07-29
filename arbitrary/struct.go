package arbitrary

import (
	"fmt"
	"reflect"
)

type StructField struct {
	Name string
	Type Type
}

type Struct struct {
	Fields []StructField
}

func (s Struct) Shrink() []Type {
	return nil
}

func (s Struct) Value() reflect.Value {
	fields := make([]reflect.StructField, len(s.Fields), len(s.Fields))
	for index, field := range s.Fields {
		fmt.Println(field.Type.Value().Type())
		if field.Type.Value().Type().PkgPath() != "" {
			continue
		}
		fields[index] = reflect.StructField{
			Name: field.Name,
			Type: field.Type.Value().Type(),
		}
	}

	v := reflect.New(reflect.StructOf(fields)).Elem()
	for index, field := range s.Fields {
		v.Field(index).Set(field.Type.Value())
	}
	return v
}
