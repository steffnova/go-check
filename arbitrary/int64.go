package arbitrary

import (
	"github.com/steffnova/go-check/constraints"
	"reflect"
)

type Int64 struct {
	Constraint constraints.Int64
	N          int64
}

func (n Int64) Shrink() []Type {
	switch {
	// Shrink towards Maximum if both Maximum and Minimum are negative
	case n.Constraint.Min < 0 && n.Constraint.Max < 0:
		diff := n.Constraint.Max - n.N
		current := n.N
		shrinks := []Type{}
		for {
			diff = diff / int64(2)
			if diff == 0 {
				break
			}
			current = current + diff
			shrinks = append(shrinks, Int64{
				Constraint: n.Constraint,
				N:          current,
			})
		}
		return shrinks
	// Shrink towards minimum if both Maximum and Minimum are positive
	case n.Constraint.Min > 0 && n.Constraint.Max > 0:
		return nil
	// Shrink towards 0
	default:
		diff := n.Constraint.Max - n.N
		current := n.N
		shrinks := []Type{}
		for {
			diff = diff / int64(2)
			if diff == 0 {
				break
			}
			current = current + diff
			shrinks = append(shrinks, Int64{
				Constraint: n.Constraint,
				N:          current,
			})
		}
		return shrinks
	}
}

func (n Int64) Value() reflect.Value {
	return reflect.ValueOf(n.N)
}
