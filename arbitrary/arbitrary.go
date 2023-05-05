package arbitrary

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Arbitrary represents arbitrary type returned by shrinkers and generators.
type Arbitrary struct {
	Value      reflect.Value // Final value
	Elements   Arbitraries   // Arbitrary for each element in collection
	Precursors Arbitraries   // Precursor arbitraries from which this one is generated
}

func (arb Arbitrary) CompareType(target Arbitrary) error {
	if arb.Value.Type() != target.Value.Type() {
		return fmt.Errorf("invalid type")
	}
	if len(arb.Precursors) != len(target.Precursors) {
		return fmt.Errorf("invalid type")
	}
	for index := range arb.Precursors {
		if err := arb.Precursors[index].CompareType(target.Precursors[index]); err != nil {
			return err
		}
	}

	return nil
}

type Arbitraries []Arbitrary

func (arbs Arbitraries) Values() []reflect.Value {
	out := make([]reflect.Value, len(arbs))
	for index, arb := range arbs {
		out[index] = arb.Value
	}
	return out
}

func NewSlice(t reflect.Type) func(Arbitrary) Arbitrary {
	return func(arb Arbitrary) Arbitrary {
		arb.Value = reflect.MakeSlice(t, len(arb.Elements), len(arb.Elements))
		for index, element := range arb.Elements {
			arb.Value.Index(index).Set(element.Value)
		}

		return arb
	}
}

func NewArray(t reflect.Type) func(Arbitrary) Arbitrary {
	return func(arb Arbitrary) Arbitrary {
		arb.Value = reflect.New(t).Elem()
		for index, element := range arb.Elements {
			arb.Value.Index(index).Set(element.Value)
		}

		return arb
	}
}

func NewMap(t reflect.Type) func(Arbitrary) Arbitrary {
	return func(arb Arbitrary) Arbitrary {
		arb.Value = reflect.MakeMap(t)
		for _, node := range arb.Elements {
			key, value := node.Elements[0], node.Elements[1]
			arb.Value.SetMapIndex(key.Value, value.Value)
		}
		return arb
	}
}

func NewStruct(t reflect.Type) func(Arbitrary) Arbitrary {
	return func(arb Arbitrary) Arbitrary {
		arb.Value = reflect.New(t).Elem()
		for index, element := range arb.Elements {
			reflect.NewAt(
				arb.Value.Field(index).Type(),
				unsafe.Pointer(arb.Value.Field(index).UnsafeAddr()),
			).Elem().Set(element.Value)
		}
		return arb
	}
}
