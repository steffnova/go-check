package generator

import (
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Invalid returns arbitrary.Generator that always returns an error. "err" parameter
// specifies error returned by generator
func Invalid(err error) arbitrary.Generator {
	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
		return arbitrary.Arbitrary{}, err
	}
}
