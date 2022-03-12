package arbitrary

import "reflect"

// FilterPredicate returns predicate function that can be used in Filter methods for generators
// and shrinkers. First parameter in is used to define predicate function signature as input
// value. Output value is always bool. Second parameter predicateFn is a arbitrary function that
// defines the behaviour of predicate.
func FilterPredicate(in reflect.Type, predicate func(reflect.Value) bool) interface{} {
	boolType := reflect.TypeOf(bool(false))
	filterSignature := reflect.FuncOf([]reflect.Type{in}, []reflect.Type{boolType}, false)

	filter := reflect.MakeFunc(filterSignature, func(arg []reflect.Value) []reflect.Value {
		out := predicate(arg[0])
		return []reflect.Value{reflect.ValueOf(out)}
	})

	return filter.Interface()
}
