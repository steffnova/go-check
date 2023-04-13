package generator

// OneFrom returns one of the provided generators. Error is returned if
// number of generators is 0, or chosen generator returns an error.
func OneFrom(generator Generator, generators ...Generator) Generator {
	generators = append([]Generator{generator}, generators...)
	weights := make([]uint64, len(generators))

	for index := range generators {
		weights[index] = 1
	}

	return Weighted(weights, generators...)
}
