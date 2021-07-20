package arbitrary

import "reflect"

type Mapped struct {
	Base   Type
	Mapper interface{}
	Val    reflect.Value
}

func (mapped Mapped) Shrink() []Type {
	baseShrinks := mapped.Base.Shrink()
	shrinks := make([]Type, len(baseShrinks))
	for index, shrink := range baseShrinks {
		inputs := []reflect.Value{shrink.Value()}
		outputs := reflect.ValueOf(mapped.Mapper).Call(inputs)
		shrinks[index] = outputs[0].Interface().(Type)
	}
	return shrinks
}

func (derived Mapped) Value() reflect.Value {
	return derived.Val
}
