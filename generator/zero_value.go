package generator

import (
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func zeroValue() arbitrary.Generator {
	return func(target reflect.Type, _ constraints.Bias, _ arbitrary.Random) (arbitrary.Arbitrary, error) {
		return arbitrary.Arbitrary{
			Value: reflect.Zero(target),
		}, nil
	}
}
