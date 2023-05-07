package generator

import (
	"fmt"
	"math"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

type weightedGenerator struct {
	weight    uint64
	generator arbitrary.Generator
}

// func Weighted2(weight uint64, generator arbitrary.Generator) weightedarbitrary.Generator {
// 	return weightedarbitrary.Generator{
// 		weight:    weight,
// 		generator: generator,
// 	}
// }

// Weighted returns one of the generators based on their weight. Weights and
// generators are specified by "weights" and "generators" parameters respectively.
// Number of weights and generators must be the same and greater than 0. Total sum
// of all weights can't exceed math.Uint64. Error is returned if number of "weights"
// and "generators" is invalid, sum of all weights exceed math.Uint64, weight value
// is lower than 1, or weighted generators returns an error.
func Weighted(weights []uint64, generators ...arbitrary.Generator) arbitrary.Generator {
	switch {
	case len(weights) == 0:
		return Invalid(fmt.Errorf("%w. Number of weights can't be 0", ErrorInvalidConfig))
	case len(generators) == 0:
		return Invalid(fmt.Errorf("%w. Number of generators can't be 0", ErrorInvalidConfig))
	case len(weights) != len(generators):
		return Invalid(fmt.Errorf("%w. Number of weights and generators must be the same", ErrorInvalidConfig))
	default:
		return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
			totalWeight := uint64(0)

			weightsIndex := make([]uint64, len(generators))
			arbitraries := make([]arbitrary.Arbitrary, len(generators))

			for index, generator := range generators {
				if weights[index] < 1 {
					return arbitrary.Arbitrary{}, fmt.Errorf("%w. Weight can't be less than 1: weights[%d] %d", ErrorInvalidConfig, index, weights[index])
				}
				gen, err := generator(target, bias, r)
				if err != nil {
					return arbitrary.Arbitrary{}, fmt.Errorf("failed to instantiate generator with index: %d. %w", index, err)
				}

				prevWeight := totalWeight
				totalWeight += weights[index]
				if index == 0 {
					totalWeight -= 1
				}
				if prevWeight > totalWeight {
					return arbitrary.Arbitrary{}, fmt.Errorf("%w. Total weght overflow. (sum of all weights can't exceed %d)", ErrorInvalidConfig, uint(math.MaxUint64))
				}
				weightsIndex[index] = totalWeight
				arbitraries[index] = gen
			}

			x := r.Uint64(constraints.Uint64{Min: 0, Max: uint64(totalWeight)})
			arb := arbitraries[0]
			for index, weight := range weightsIndex {
				if weight >= x {
					arb = arbitraries[index]
					break
				}
			}

			return arb, nil
		}
	}
}
