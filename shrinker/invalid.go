package shrinker

import "reflect"

// Invalid is shrinker that always returns an error. Error value is specified
// with err parameter.
func Invalid(err error) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		return reflect.Value{}, nil, err
	}
}
