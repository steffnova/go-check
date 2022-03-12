package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

func ExampleFunc() {
	// Streamer uses Func generator to generate pure functions.
	check.Stream(check.Streamer(
		func(f func(x, y int) []int) {
			fmt.Printf("%#v\n", f(1, 2))
		},
		// Function generators are defined by their output parameters.
		// For each output parameter a generator needs to be provided.
		// In this example, function f return []int so corresponding
		// generator needs to be provided.
		generator.Func(generator.Slice(
			generator.Int(constraints.Int{Min: 0, Max: 10}),
			constraints.Length{
				Min: 2,
				Max: 5,
			},
		)),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// []int{3, 3, 4, 9, 3}
	// []int{10, 4}
	// []int{8, 3}
	// []int{9, 2}
	// []int{2, 6, 8, 2, 8}
	// []int{7, 7}
	// []int{1, 6}
	// []int{9, 6}
	// []int{5, 5, 9, 1}
	// []int{10, 1, 9, 6}
}
