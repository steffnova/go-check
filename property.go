package check

import (
	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

type Details struct {
	NumberOfIterations int
	NumberOfShrinks    uint
	FailureReason      error
	FailureInput       arbitrary.Arbitraries
}

type property func(r arbitrary.Random, iterations int) (Details, error)

// Property defines a new property by specifing predicate and property generators.
// Predicate must be a function that can have any number of input values, and must
// have only one output value of error type. Number of predicate's input parameters
// must match number of generators.
func Property(in inputsGenerator, p predicate) property {
	return func(r arbitrary.Random, iterations int) (Details, error) {
		targets, runner := p()

		for i := 0; i < iterations; i++ {
			bias := constraints.Bias{
				Size:    iterations,
				Scaling: iterations - i,
			}

			arbs, shrinker, err := in(targets, bias, r)
			if err != nil {
				return Details{}, err
			}

			predicateErr := runner(arbs)
			if predicateErr == nil {
				continue
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
				FailureInput:       arbs,
				FailureReason:      predicateErr,
				NumberOfShrinks:    uint(numberOfShrinks),
				NumberOfIterations: i,
			}, nil
		}

		return Details{}, nil
	}
}
