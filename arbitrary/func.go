package arbitrary

import "reflect"

type Func struct {
	Fn reflect.Value
}

func (f Func) Shrink() []Type {
	return nil
}

func (f Func) Value() reflect.Value {
	return f.Fn
}
