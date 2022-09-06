package generator

import (
	"reflect"
)

// Invalid returns Generator that always returns an error. "err" parameter
// specifies error returned by generator
func Invalid(err error) Generator {
	return func(target reflect.Type, r Random) (Generate, error) {
		return nil, err
	}
}
