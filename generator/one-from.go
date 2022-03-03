package generator

import "fmt"

// OneFrom returns one of the provided generators. Error is returned if
// number of generators is 0, or chosen generator returns an error.
func OneFrom(generators ...Generator) Generator {
	if len(generators) == 0 {
		return Invalid(fmt.Errorf("number of generators must be greater than 0"))
	}

	weights := make([]uint64, len(generators))
	for index := range generators {
		weights[index] = 1
	}

	return Weighted(weights, generators...)
}
