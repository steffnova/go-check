package generator

func OneFrom(first Generator, other ...Generator) Generator {
	arbs := append([]Generator{first}, other...)
	weightedArbs := make([]Weighted, len(arbs))

	for index, arb := range arbs {
		weightedArbs[index] = Weighted{
			Weight: 1,
			Gen:    arb,
		}
	}

	return OneFromWeighted(weightedArbs[0], weightedArbs[1:]...)
}
