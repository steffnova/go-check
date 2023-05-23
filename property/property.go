package property

import (
	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

type Details struct {
	NumberOfShrinks uint
	FailureReason   error
	FailureInput    arbitrary.Arbitraries
}

// Property is a function that takes [arbitrary.Random] and [constraints.Bias] parameters as inputs
// and returns [Details] and an error as output parameters. See [Define] for usage.
type Property func(r arbitrary.Random, bias constraints.Bias) (Details, error)

// Define creates a new property by specifying an input generator and a predicate.
// The generator is specified using [Inputs], and the predicate is specified using [Predicate].
// The following example demonstrates how to define a property:
//
//	// The number of generators used for inputs must match the number of predicate input
//	// parameters. The generator at index i must be able to generate a value for the predicate's
//	// parameter at index i. In this example, two generator.Int() generators are used
//	// because the predicate's input parameters x and y are of type int.
//	property.Define(
//		property.Inputs(
//			generator.Int(),
//			generator.Int(),
//		),
//		property.Predicate(func(x int, y int) error {
//			// ... rest of the predicate logic
//			return nil
//		}),
//	)
func Define(generator InputsGenerator, predicate predicate) Property {
	return func(r arbitrary.Random, bias constraints.Bias) (Details, error) {
		targets, runner := predicate()

		arbs, shrinker, err := generator(targets, bias, r)
		if err != nil {
			return Details{}, err
		}

		predicateErr := runner(arbs)
		if predicateErr == nil {
			return Details{}, nil
		}

		numberOfShrinks := uint(0)

		for shrinker != nil {
			var shrinkingErr error
			propertyFailed := predicateErr != nil
			arbs, shrinker, shrinkingErr = shrinker(arbs, propertyFailed)
			if shrinkingErr != nil {
				return Details{}, err
			}
			predicateErr = runner(arbs)
		}

		return Details{
			FailureInput:    arbs,
			FailureReason:   predicateErr,
			NumberOfShrinks: uint(numberOfShrinks),
		}, nil
	}
}
