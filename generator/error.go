package generator

import (
	"fmt"
	"reflect"
)

var (
	ErrorInvalidTarget         = fmt.Errorf("Invalid generator target")
	ErrorInvalidConstraints    = fmt.Errorf("Invalid generator constraints")
	ErrorInvalidCollectionSize = fmt.Errorf("Invalid number of generators for collection")
	ErrorInvalidConfig         = fmt.Errorf("Invalid generator configuration")
	ErrorMapper                = fmt.Errorf("Generator mapper is invalid")
	ErrorFilter                = fmt.Errorf("Generator filter's predicate is invalid")
	ErrorBinder                = fmt.Errorf("Geneartor binder is invalid")
	ErrorStream                = fmt.Errorf("Failed to stream data")
)

func NewErrorInvalidTarget(target reflect.Type, generatorType string) error {
	return fmt.Errorf("%w. Can't use %s generator for %s type", ErrorInvalidTarget, generatorType, target)
}

func NewErrorInvalidCollectionSize(collectionSize int, numberOfGenerators int) error {
	return fmt.Errorf(
		"%w. Collection size is %d, number of collection element generators: %d",
		ErrorInvalidCollectionSize,
		collectionSize,
		numberOfGenerators,
	)
}
