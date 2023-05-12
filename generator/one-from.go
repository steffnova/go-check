package generator

import "github.com/steffnova/go-check/arbitrary"

// OneFrom returns one of the provided generators. Error is returned if
// number of generators is 0, or chosen generator returns an error.
func OneFrom(generator arbitrary.Generator, generators ...arbitrary.Generator) arbitrary.Generator {
	generators = append([]arbitrary.Generator{generator}, generators...)
	weights := make([]uint64, len(generators))

	for index := range generators {
		weights[index] = 1
	}

	return Weighted(weights, generators...)
}
