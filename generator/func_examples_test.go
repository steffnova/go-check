package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Func() generator for generation of pure functions.
// Func() generator requires generators for it's output values to be passed to it. In this
// example generated function will return []int, thus Slice(Int()) generator is used.
func ExampleFunc() {
	streamer := generator.Streamer(
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

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// []int{4, 2, 10, 10, 3}
	// []int{4, 6, 1, 4, 9}
	// []int{7, 6, 5, 8}
	// []int{5, 8, 4, 10}
	// []int{4, 9, 0, 4}
	// []int{5, 4, 5, 3}
	// []int{10, 2, 8, 3}
	// []int{1, 9, 4, 3}
	// []int{4, 3}
	// []int{5, 5, 8, 10}
}
