package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

// Ptr is Arbitrary that creates pointer Generator. Generator will return
// either valid or invalid (nil) pointer for targetn's type. Error is returned
// if target's reflect.Kind is not Ptr, or creation of arb's Generator fails.
func Ptr(arb Arbitrary) Arbitrary {
	return OneOf(PtrInvalid(), PtrValid(arb))
}

// PtrInvalid is Arbitrary that creates pointer Generator. Generator will
// always return nil pointer for target's type. Error is returned if target's
// reflect.Kind is not Ptr.
func PtrInvalid() Arbitrary {
	return func(target reflect.Type, r Random) (Generator, error) {
		if target.Kind() != reflect.Ptr {
			return nil, fmt.Errorf("target's kind must be Ptr. Got: %s", target.Kind())
		}

		return func() arbitrary.Type {
			return arbitrary.Ptr{
				ElementType: nil,
				Type:        target,
			}
		}, nil
	}
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

		return func() arbitrary.Type {
			return arbitrary.Ptr{
				Type:        target,
				ElementType: generateValue(),
			}
		}, nil
	}
}
