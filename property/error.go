package property

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	ErrorPredicate      = fmt.Errorf("predicate configuration error") // Property predicate is invalid
	ErrorInputs         = fmt.Errorf("inputs generator error")        // Property input generator is invalid
	ErrorPropertyConfig = fmt.Errorf("property configuration error")  // Property returned an error
)

func inputMissmatchError(targets []reflect.Type, predicate reflect.Type) error {
	expected := make([]string, len(targets))
	for index, target := range targets {
		expected[index] = target.String()
	}

	got := make([]string, predicate.NumIn())
	for index := 0; index < predicate.NumIn(); index++ {
		got[index] = predicate.In(index).String()
	}

	return fmt.Errorf(
		"%w. %s", ErrorInputs,
		fmt.Sprint(
			fmt.Sprintln("property.Generator.Filter: predicate inputs do not match property's predicate inputs"),
			fmt.Sprintf("Property Predicate: (%s)\n", strings.Join(expected, ", ")),
			fmt.Sprintf("Filter Predicate:   (%s)\n", strings.Join(got, ", ")),
		),
	)
}
