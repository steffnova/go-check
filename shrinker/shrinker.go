package shrinker

import "reflect"

// Shrinker returns shrinked value and next shrinker. Returned values
// are usually affected by propertyFailed parameter. It helps shrinker
// to decide how to properly shrink the value. If returned Shrinker
// is nil, it indicates that value can no longer be shrinked
type Shrinker func(propertyFailed bool) (reflect.Value, Shrinker)

// Map maps the shrinked value to a new one using the mapper. Mapper
// must be a function with one input and one output parameter. Input
// parameter must match shrinked type, otherwise panic occurs.
func (shrinker Shrinker) Map(mapper interface{}) Shrinker {
	return func(propertyFailed bool) (reflect.Value, Shrinker) {
		shrink, shrinker := shrinker(propertyFailed)
		shrink = reflect.ValueOf(mapper).Call([]reflect.Value{shrink})[0]

		if shrinker == nil {
			return shrink, nil
		}

		return shrink, shrinker.Map(mapper)
	}
}
