package check

import (
	"fmt"
	"reflect"
	"strings"
)

type propertyError func() string

func (pe propertyError) Error() string {
	return pe()
}

func propertyFailed(inputs []reflect.Value) propertyError {
	return func() string {
		inputData := make([]string, len(inputs))
		for index, input := range inputs {
			inputData[index] = fmt.Sprintf("<%s> %#v", input.Type().String(), input.Interface())
		}

		return fmt.Sprintf("Property failed for inputs: [\n\t%s\n]", strings.Join(inputData, ",\n\t"))
	}
}
