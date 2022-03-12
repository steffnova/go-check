package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

func ExampleOneFrom() {
	// Streamer uses OneOf generator to generate int values.
	check.Stream(check.Streamer(
		func(n int) {
			fmt.Printf("%d\n", n)
		},
		// OneFrom will select one of the generators passed to it
		// and use it to generate int values.
		generator.OneFrom(
			generator.Int(constraints.Int{Min: 500, Max: 1000}),
			generator.Int(constraints.Int{Min: 5, Max: 10}),
			generator.Int(constraints.Int{Min: -10, Max: -5}),
			generator.Int(constraints.Int{Min: -1000, Max: -500}),
		),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// -5
	// 842
	// 10
	// -948
	// -887
	// 8
	// 690
	// 592
	// -6
	// -10
}
