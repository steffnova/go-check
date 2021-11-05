package shrinker

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/constraints"
)

type Value struct {
	Value    reflect.Value
	Shrinker Shrinker
}

type MapElement struct {
	Key   Value
	Value Value
}

func Map(val reflect.Value, mapElements []MapElement, limits constraints.Length) Shrinker {
	return MapSize(val, mapElements, limits).
		Compose(MapValues(val, mapElements, limits)).
		Compose(MapKeys(val, mapElements, limits))
}

func MapSize(val reflect.Value, mapElements []MapElement, limits constraints.Length) Shrinker {
	mapperSignature := reflect.FuncOf(
		[]reflect.Type{reflect.TypeOf(mapElements)},
		[]reflect.Type{val.Type()},
		false,
	)
	mapper := reflect.MakeFunc(mapperSignature, func(args []reflect.Value) (results []reflect.Value) {
		elements := args[0].Interface().([]MapElement)

		for _, key := range val.MapKeys() {
			val.SetMapIndex(key, reflect.Value{})
		}
		for _, element := range elements {
			val.SetMapIndex(element.Key.Value, element.Value.Value)
		}
		return []reflect.Value{val}
	})

	return sliceSize(reflect.ValueOf(mapElements), 0, limits).
		Map(mapper.Interface())
}

func MapValue(val reflect.Value, element MapElement) Shrinker {
	if element.Value.Shrinker == nil {
		return nil
	}

	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		value, shrinker, err := element.Value.Shrinker(propertyFailed)
		switch {
		case err != nil:
			return reflect.Value{}, nil, fmt.Errorf("failed to shrink map's value: %v. %w", element.Value.Value.Interface(), err)
		case !val.MapIndex(element.Key.Value).IsValid():
			return val, nil, nil
		default:
			element.Value.Value, element.Value.Shrinker = value, shrinker
			val.SetMapIndex(element.Key.Value, element.Value.Value)
			return val, MapValue(val, element), nil
		}
	}
}

func MapValues(val reflect.Value, elements []MapElement, limits constraints.Length) Shrinker {
	var shrinker Shrinker
	for _, tempElement := range elements {
		shrinker = shrinker.Compose(MapValue(val, tempElement))
	}

	return shrinker
}

func MapKey(val reflect.Value, element MapElement) Shrinker {

	if element.Key.Shrinker == nil {
		return nil
	}
	element.Value.Value = val.MapIndex(element.Key.Value)

	return func(propertyFailed bool) (reflect.Value, Shrinker, error) {
		value, shrinker, err := element.Key.Shrinker(propertyFailed)
		switch {
		case err != nil:
			return reflect.Value{}, nil, fmt.Errorf("failed to shrink map's key: %v. %w", element.Key.Value.Interface(), err)
		case !val.MapIndex(element.Key.Value).IsValid():
			return val, nil, nil
		case val.MapIndex(value).IsValid() || reflect.DeepEqual(value, element.Key.Value):
			element.Key.Shrinker = shrinker
			return val, MapKey(val, element), nil
		default:
			currentValue := val.MapIndex(element.Key.Value)
			val.SetMapIndex(element.Key.Value, reflect.Value{})
			val.SetMapIndex(value, currentValue)
			element.Key.Value, element.Key.Shrinker = value, shrinker
			return val, MapKey(val, element), nil
		}
	}
}

func MapKeys(val reflect.Value, elements []MapElement, limits constraints.Length) Shrinker {
	var shrinker Shrinker
	for _, tempElement := range elements {
		shrinker = shrinker.Compose(MapKey(val, tempElement))
	}

	return shrinker
}
