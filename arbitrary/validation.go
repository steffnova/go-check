package arbitrary

import (
	"fmt"
	"reflect"
)

type Validate func(Arbitrary) error

func ValidateSlice() Validate {
	return func(arb Arbitrary) error {
		switch {
		case arb.Value.Kind() != reflect.Slice:
			return fmt.Errorf("arb is not a slice")
		case arb.Value.Len() != len(arb.Elements):
			return fmt.Errorf("number of elements %d must match size of the slice %d", len(arb.Elements), arb.Value.Len())
		default:
			return nil
		}
	}
}

func ValidateMap() Validate {
	return func(arb Arbitrary) error {
		switch {
		case arb.Value.Kind() != reflect.Map:
			return fmt.Errorf("arb is not a map")
		default:
			return nil
		}
	}
}

func ValidateArray() Validate {
	return func(arb Arbitrary) error {
		switch {
		case arb.Value.Kind() != reflect.Array:
			return fmt.Errorf("arb is not a array")
		case arb.Value.Len() != len(arb.Elements):
			return fmt.Errorf("number of elements %d must match size of the array %d", len(arb.Elements), arb.Value.Len())
		default:
			return nil
		}
	}
}

func ValidateStruct() Validate {
	return func(arb Arbitrary) error {
		switch {
		case arb.Value.Kind() != reflect.Struct:
			return fmt.Errorf("arb is not a struct")
		case arb.Value.NumField() != len(arb.Elements):
			return fmt.Errorf("number of elements %d must match number of the struct fields %d", len(arb.Elements), arb.Value.Len())
		default:
			return nil
		}
	}
}
