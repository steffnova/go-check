package arbitrary

import "reflect"

type Type interface {
	Shrink() []Type
	Value() reflect.Value
}
