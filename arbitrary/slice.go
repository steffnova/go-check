package arbitrary

import (
	"github.com/steffnova/go-check/constraints"
	"reflect"
)

type Slice struct {
	Constraint  constraints.Length
	ElementType reflect.Type
	Elements    []Type
}

func (s Slice) Shrink() []Type {
	return nil
}

func (s Slice) Value() reflect.Value {
	val := reflect.MakeSlice(reflect.SliceOf(s.ElementType), len(s.Elements), len(s.Elements))
	for index, element := range s.Elements {
		val.Index(index).Set(element.Value())
	}
	return val
}
