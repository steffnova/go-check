package generator

import (
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Invalid returns Generator that always returns an error. "err" parameter
// specifies error returned by generator
func Invalid(err error) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (arbitrary.Arbitrary, shrinker.Shrinker, error) {
		return arbitrary.Arbitrary{}, nil, err
	}
}
