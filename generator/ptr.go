package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/shrinker"
)

// Ptr is Arbitrary that creates pointer Generator. Generator will return
// either valid or invalid (nil) pointer for target's type. Error is returned
// if target's reflect.Kind is not Ptr, or creation of arb's Generator fails.
func Ptr(arb Arbitrary) Arbitrary {
	return OneFrom(Nil(), PtrValid(arb))
}

// PtrValid is Arbitrary that creates pointer Generator. Generator will
// always return non-nil pointer for target's type. Error is returned
// if target's reflect.Kind is not Ptr, or creation of arb's Generator
// fails.
func PtrValid(arb Arbitrary) Arbitrary {
	return func(target reflect.Type, r Random) (Generator, error) {
		if target.Kind() != reflect.Ptr {
			return nil, fmt.Errorf("target's kind must be Ptr. Got: %s", target.Kind())
		}

		generateValue, err := arb(target.Elem(), r)
		if err != nil {
			return nil, fmt.Errorf("failed to create base generator. %s", err)
		}

		return func() (reflect.Value, shrinker.Shrinker) {
			val, valShrinker := generateValue()
			ptr := reflect.New(target.Elem())
			ptr.Elem().Set(val)
			return ptr, shrinker.Ptr(ptr, valShrinker)
		}, nil
	}
}
