package generator

import (
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

// Invalid returns Generator that always returns an error. "err" parameter
// specifies error returned by generator
func Invalid(err error) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		return nil, err
	}
}
