package arbitrary

import (
	"github.com/steffnova/go-check/constraints"
	"reflect"
)

type Uint64 struct {
	Constraint constraints.Uint64
	N          uint64
}

func (n Uint64) Shrink() []Type {
	return nil
}

func (n Uint64) Value() reflect.Value {
	return reflect.ValueOf(n.N)
}
