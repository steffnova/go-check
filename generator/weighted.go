package generator

import (
	"fmt"
	"math"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

// Weighted returns one of the generators based on their weight. Weights and
// generators are specified by "weights" and "generators" parameters respectively.
// Number of weights and generators must be the same and greater than 0. Total sum
// of all weights can't exceed math.Uint64. Error is returned if number of "weights"
// and "generators" is invalid, sum of all weights exceed math.Uint64, weight value
// is lower than 1, or weighted generators returns an error.
func Weighted(weights []uint64, generators ...Generator) Generator {
	switch {
	case len(weights) == 0:
		return Invalid(fmt.Errorf("number of weights can't be 0"))
	case len(generators) == 0:
		return Invalid(fmt.Errorf("number of generators can't be 0"))
	case len(weights) != len(generators):
		return Invalid(fmt.Errorf("number of weights and generators must be the same"))
	default:
		return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
			totalWeight := uint64(0)

			weightsIndex := make([]uint64, len(generators))
			generates := make([]Generate, len(generators))

			for index, generator := range generators {
				if weights[index] < 1 {
					return nil, fmt.Errorf("weight can't be less than 1: %d", weights[index])
				}
				gen, err := generator(target, bias, r)
				if err != nil {
					return nil, fmt.Errorf("faile to instantiate generator with index: %d. %s", index, err)
				}

				prevWeight := totalWeight
				totalWeight += weights[index]
				if index == 0 {
					totalWeight -= 1
				}
				if prevWeight > totalWeight {
					return nil, fmt.Errorf("total weght overflow. (sum of all weights can't exceed %d)", uint(math.MaxUint64))
				}
				weightsIndex[index] = totalWeight
				generates[index] = gen
			}

			return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
				x := r.Uint64(constraints.Uint64{Min: 0, Max: uint64(totalWeight)})
				generator := generates[0]
				for index, weight := range weightsIndex {
					if weight >= x {
						generator = generates[index]
						break
					}
				}

				return generator()
			}, nil
		}
	}
}
