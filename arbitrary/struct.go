package arbitrary

import "reflect"

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
