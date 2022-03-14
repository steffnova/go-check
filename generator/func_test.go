package generator_test

import (
	"fmt"

	check "github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Func() generator for generation of pure functions.
// Func() generator requires generators for it's output values to be passed to it. In this
// example generated function will return []int, thus Slice(Int()) generator is used.
func ExampleFunc() {
	streamer := check.Streamer(
		func(f func(x, y int) []int) {
			fmt.Printf("%#v\n", f(1, 2))
		},
		generator.Func(generator.Slice(
			generator.Int(constraints.Int{
				Min: 0,
				Max: 10,
			}),
			constraints.Length{
				Min: 2,
				Max: 5,
			},
		)),
	)

	if err := check.Stream(streamer, check.Config{Seed: 0, Iterations: 10}); err != nil {
		panic(err)
	}
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
