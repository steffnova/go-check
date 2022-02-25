package generator

import (
	"fmt"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/shrinker"
)

func Slice(element Generator, limits ...constraints.Length) Generator {
	return func(target reflect.Type, bias constraints.Bias, r Random) (Generate, error) {
		constraint := constraints.LengthDefault()
		if len(limits) != 0 {
			constraint = limits[0]
		}
		if target.Kind() != reflect.Slice {
			return nil, fmt.Errorf("targets kind must be Slice. Got: %s", target.Kind())
		}

		generator, err := element(target.Elem(), bias, r)
		if err != nil {
			return nil, fmt.Errorf("failed to create generator for slice elements: %s", err)
		}

		return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
			biasedConstraints := constraints.Int64{
				Min: int64(constraint.Min),
				Max: int64(constraint.Max),
			}.Biased(bias)
			size := r.Int64(biasedConstraints)

			shrinkers := make([]shrinker.Shrinker, size)

			arb := arbitrary.Arbitrary{
				Value:    reflect.MakeSlice(target, int(size), int(size)),
				Elements: make([]arbitrary.Arbitrary, int(size)),
			}

			for index := range shrinkers {
				arb.Elements[index], shrinkers[index] = generator()
				arb.Value.Index(index).Set(arb.Elements[index].Value)
			}

			return arb, shrinker.Slice(shrinker.CollectionSize(arb.Elements, shrinkers, 0, constraints.Length{
				Min: int(biasedConstraints.Min),
				Max: int(biasedConstraints.Max),
			}))
		}, nil
	}

}
