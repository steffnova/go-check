package arbitrary

import (
	"reflect"
)

type Ptr struct {
	Type        reflect.Type
	ElementType Type
}

func (ptr Ptr) Shrink() []Type {
	if ptr.ElementType == nil {
		return nil
	}

	shrinked := ptr.ElementType.Shrink()
	result := make([]Type, len(shrinked))
	for index, t := range shrinked {
		result[index] = Ptr{
			ElementType: t,
		}
	}
	return append(result, Ptr{
		ElementType: nil,
	})
}

func (ptr Ptr) Value() reflect.Value {
	if ptr.ElementType == nil {
		return reflect.Zero(ptr.Type)
	}
	val := reflect.New(ptr.Type.Elem())
	val.Elem().Set(ptr.ElementType.Value())
	return val
}
