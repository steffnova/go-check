package arbitrary

import "reflect"

// Mapper returns mapper function that can be used with Map combinator for generators and
// shrinkers. Parameter "in' defines mapper's input type, parameter "out" defines mapper's
// output type while parameter "mapFn" implements mapping.
func Mapper(in, out reflect.Type, mapFn func(reflect.Value) reflect.Value) interface{} {
	mapperSignature := reflect.FuncOf([]reflect.Type{in}, []reflect.Type{out}, false)

	mapper := reflect.MakeFunc(mapperSignature, func(arg []reflect.Value) []reflect.Value {
		return []reflect.Value{mapFn(arg[0])}
	})

	return mapper.Interface()
}
