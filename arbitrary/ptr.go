package arbitrary

import (
	"reflect"
)

type Ptr struct {
	IsNull bool
	Type   Type
}

func (ptr Ptr) Shrink() []Type {
	if ptr.IsNull {
		return nil
	}

	shrinked := ptr.Type.Shrink()
	result := make([]Type, len(shrinked), len(shrinked))
	for index, t := range shrinked {
		result[index] = Ptr{
			IsNull: false,
			Type:   t,
		}
	}
	return append(result, Ptr{
		IsNull: true,
	})
}

func (ptr Ptr) Value() reflect.Value {
	if ptr.IsNull {
		return reflect.Zero(reflect.PtrTo(ptr.Type.Value().Type()))
	}
	val := reflect.New(ptr.Type.Value().Type())
	val.Elem().Set(ptr.Type.Value())
	return val
}
