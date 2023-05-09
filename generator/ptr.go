package generator

import (
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Ptr is Arbitrary that creates pointer arbitrary.Generator. arbitrary.Generator will return
// either valid or invalid (nil) pointer for target's type. Error is returned
// if target's reflect.Kind is not Ptr, or creation of arb's arbitrary.Generator fails.
func Ptr(arb arbitrary.Generator) arbitrary.Generator {
	return OneFrom(Nil(), PtrTo(arb))
}

// PtrTo is Arbitrary that creates pointer arbitrary.Generator. arbitrary.Generator will
// always return non-nil pointer for target's type. Error is returned
// if target's reflect.Kind is not Ptr, or creation of arb's arbitrary.Generator
// fails.
func PtrTo(arb arbitrary.Generator) arbitrary.Generator {
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		if target.Kind() != reflect.Ptr {
			return arbitrary.Arbitrary{}, arbitrary.NewErrorInvalidTarget(target, "PtrTo")
		}

		mapper := arbitrary.Mapper(target.Elem(), target, func(in reflect.Value) reflect.Value {
			out := reflect.New(target.Elem())
			out.Elem().Set(in)
			return out
		})

		return arb.Map(mapper)(target, bias, r)
	}
}
