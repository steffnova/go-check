package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Ptr is Arbitrary that creates pointer Generator. Generator will return
// either valid or invalid (nil) pointer for target's type. Error is returned
// if target's reflect.Kind is not Ptr, or creation of arb's Generator fails.
func Ptr(arb Generator) Generator {
	return OneFrom(Nil(), PtrTo(arb))
}

// PtrTo is Arbitrary that creates pointer Generator. Generator will
// always return non-nil pointer for target's type. Error is returned
// if target's reflect.Kind is not Ptr, or creation of arb's Generator
// fails.
func PtrTo(arb Generator) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		if target.Kind() != reflect.Ptr {
			return nil, fmt.Errorf("can't use Ptr generator for %s type", target)
		}

		mapper := arbitrary.Mapper(target.Elem(), target, func(in reflect.Value) reflect.Value {
			out := reflect.New(target.Elem())
			out.Elem().Set(in)
			return out
		})

		return arb.Map(mapper)(target, bias, r)
	}
}
