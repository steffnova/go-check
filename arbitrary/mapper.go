package arbitrary

import "reflect"

type mapFn func(reflect.Value) reflect.Value

func Mapper(in, out reflect.Type, mapFn mapFn) interface{} {
	mapperSignature := reflect.FuncOf([]reflect.Type{in}, []reflect.Type{out}, false)

	mapper := reflect.MakeFunc(mapperSignature, func(arg []reflect.Value) []reflect.Value {
		return []reflect.Value{mapFn(arg[0])}
	})

	return mapper.Interface()
}
