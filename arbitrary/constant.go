package arbitrary

import "reflect"

type Constant struct {
	C reflect.Value
}

func (c Constant) Shrink() []Type {
	return nil
}

func (c Constant) Value() reflect.Value {
	return c.C
}
