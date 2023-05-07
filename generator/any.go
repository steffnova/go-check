package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Any returns generator with default constraints for a type specified by generator's target.
// Unsupported target: interface{}
func Any() arbitrary.Generator {
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		var generator arbitrary.Generator
		switch target.Kind() {
		case reflect.Array:
			generator = Array(Any())
		case reflect.Bool:
			generator = Bool()
		case reflect.Complex64:
			generator = Complex64()
		case reflect.Complex128:
			generator = Complex128()
		case reflect.Chan:
			generator = Chan()
		case reflect.Float32:
			generator = Float32()
		case reflect.Float64:
			generator = Float64()
		case reflect.Int:
			generator = Int()
		case reflect.Int8:
			generator = Int8()
		case reflect.Int16:
			generator = Int16()
		case reflect.Int32:
			generator = Int32()
		case reflect.Int64:
			generator = Int64()
		case reflect.Uint:
			generator = Uint()
		case reflect.Uint8:
			generator = Uint8()
		case reflect.Uint16:
			generator = Uint16()
		case reflect.Uint32:
			generator = Uint32()
		case reflect.Uint64:
			generator = Uint64()
		case reflect.Func:
			outputs := make([]arbitrary.Generator, target.NumOut())
			for index := range outputs {
				outputs[index] = Any()
			}
			generator = Func(outputs...)
		case reflect.Map:
			generator = Map(Any(), Any())
		case reflect.Ptr:
			generator = Ptr(Any())
		case reflect.Struct:
			generator = Struct()
		case reflect.Slice:
			generator = Slice(Any())
		case reflect.String:
			generator = String()
		default:
			return arbitrary.Arbitrary{}, fmt.Errorf("%w. Any generator does not support values of kind: %s", ErrorInvalidTarget, target.Kind())
		}

		return generator(target, bias, r)
	}
}
