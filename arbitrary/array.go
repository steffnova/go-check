package arbitrary

import "reflect"

type Array struct {
	Elements []Type
	Val      reflect.Value
}

func (a Array) Shrink() []Type {
	return nil
}

func (a Array) Value() reflect.Value {
	return a.Val
}
