package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/generator"
)

func ExampleWeighted() {
	// Streamer uses Weighted generator to generate *uint64 values.
	check.Stream(check.Streamer(
		func(n *uint64) {
			if n == nil {
				fmt.Printf("%v\n", n)
			} else {
				fmt.Printf("%d\n", *n)
			}
		},
		// Weighted generator as a first parameter accepts slice of uint64
		// values that define weights for each of generators passed to it.
		// In this case weight 1 is assigned to generator.Nil(), and 9
		// to generator.PtrTo(). Weights define frequency to determine
		// which generator will be picked based on it's weight. To calculate
		// the chance, all weights are summed together, and individiual chance
		// for each generator is calculated as it's weight devided by sum of
		// all weight and multiplied by 100. In this case generator.Nil() has
		// 10% chance (1/10 * 100) and generator.PtrTo has 90% chance (9/10 * 100)
		generator.Weighted(
			[]uint64{1, 9},
			generator.Nil(),
			generator.PtrTo(generator.Uint64()),
		),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 1002466900374765554
	// <nil>
	// <nil>
	// 14746210962209877445
	// 12784885724210938115
	// 11116474692239114024
	// 15398783846516204029
	// 14677457169740829639
	// 9472434474353809100
	// 2396012503939351775
}
