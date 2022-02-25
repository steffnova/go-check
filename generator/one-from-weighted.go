package generator

import (
	"fmt"
	"math"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

type Weighted struct {
	Weight uint
	Gen    Generator
}

func OneFromWeighted(first Weighted, other ...Weighted) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		totalWeight := uint(0)

		arbs := append([]Weighted{first}, other...)
		weights := make([]uint, len(arbs))
		generators := make([]Generate, len(arbs))

		for index, weightedGen := range arbs {
			if weightedGen.Weight < 1 {
				return nil, fmt.Errorf("weight can't be less than 1: %d", weightedGen.Weight)
			}
			gen, err := weightedGen.Gen(target, bias, r)
			if err != nil {
				return nil, fmt.Errorf("faile to instantiate generator with index: %d. %s", index, err)
			}

			prevWeight := totalWeight
			totalWeight += weightedGen.Weight
			if index == 0 {
				totalWeight -= 1
			}
			if prevWeight > totalWeight {
				return nil, fmt.Errorf("total weght overflow. (sum of all weights can't exceed %d)", uint(math.MaxUint64))
			}
			weights[index] = totalWeight
			generators[index] = gen
		}

		return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
			x := r.Uint64(constraints.Uint64{Min: 0, Max: uint64(totalWeight)})

			generator := generators[0]
			for index, weight := range weights {
				if weight >= uint(x) {
					generator = generators[index]
					break
				}
			}

			return generator()
		}, nil
	}
}
