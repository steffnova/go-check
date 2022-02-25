package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

// Any is Arbitrary that returns default Generator for Generator's target parameter. Any will
// return Generator for all Go's types except interface{} types. If User defined type is passed,
// default generator for it's reflect.Kind will be returned. Default generator for any type is
// a generator with default constraints. Error will be thrown if target type is not supported.
func Any() Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		var generator Generator
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
			outputs := make([]Generator, target.NumOut())
			for index := range outputs {
				outputs[index] = Any()
			}
			generator = Func(outputs...)
		case reflect.Map:
			generator = Map(Any(), Any())
		// case reflect.Ptr:
		// 	generator = Ptr(Any())
		case reflect.Struct:
			generator = Struct()
		case reflect.Slice:
			generator = Slice(Any())
		case reflect.String:
			generator = String()
		default:
			return nil, fmt.Errorf("no support for generating values for kind: %s", target.Kind())
		}

		return generator(target, bias, r)
	}
}
