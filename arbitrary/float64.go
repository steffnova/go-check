package arbitrary

import (
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

type Float64 struct {
	Constraint constraints.Float64
	N          float64
}

func (n Float64) Shrink() []Type {
	return nil
}

func (n Float64) Value() reflect.Value {
	return reflect.ValueOf(n.N)
}
