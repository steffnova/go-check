package generator

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
)

// Ptr is Arbitrary that creates pointer Generator. Type to which pointer
// points is defined by arb paramter. Arbitrary fails to create pointer
// Generator if target's reflect.Kind is not Ptr, or arb fails to generate
// Generator for value pointer points to.
func Ptr(arb Arbitrary) Arbitrary {
	return func(target reflect.Type) (Generator, error) {
		if target.Kind() != reflect.Ptr {
			return nil, fmt.Errorf("target's kind must be Ptr. Got: %s", target.Kind())
		}
		generateValue, err := arb(target.Elem())
		if err != nil {
			return nil, fmt.Errorf("failed to create base generator. %s", err)
		}
		generateBool, err := Bool()(reflect.TypeOf(false))
		if err != nil {
			return nil, fmt.Errorf("failed to generate bool generator. %s", err)
		}

		return func(rand *rand.Rand) arbitrary.Type {
			isNull := generateBool(rand).Value().Bool()
			t := arbitrary.Type(nil)
			if !isNull {
				t = generateValue(rand)
			}
			return arbitrary.Ptr{
				IsNull: isNull,
				Type:   t,
			}
		}, nil
	}
}
