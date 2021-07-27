package arbitrary

import "reflect"

type Chan struct {
	C reflect.Value
}

func (c Chan) Shrink() []Type {
	return nil
}

func (c Chan) Value() reflect.Value {
	return c.C
}
