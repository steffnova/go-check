package shrinker

import "reflect"

type Shrink struct {
	Value    reflect.Value
	Shrinker Shrinker
}

type SliceShrink struct {
	Type     reflect.Type
	Elements []Shrink
}

func (ss SliceShrink) Value() reflect.Value {
	val := reflect.MakeSlice(ss.Type, len(ss.Elements), len(ss.Elements))
	for index, shrink := range ss.Elements {
		val.Index(index).Set(shrink.Value)
	}
	return val
}
