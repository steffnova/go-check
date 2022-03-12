package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

func ExampleChan() {
	// Streamer uses Chan generator to generate chan int.
	check.Stream(check.Streamer(
		func(ch chan int) {
			fmt.Printf("chan size: %d\n", cap(ch))
		},
		generator.Chan(),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// chan size: 31
	// chan size: 16
	// chan size: 80
	// chan size: 86
	// chan size: 69
	// chan size: 22
	// chan size: 84
	// chan size: 3
	// chan size: 64
	// chan size: 30
}

func ExampleChan_withConstraints() {
	// Streamer uses Chan generator to generate chan int.
	check.Stream(check.Streamer(
		func(ch chan int) {
			fmt.Printf("chan size: %d\n", cap(ch))
		},
		// Passing constraint.Length to Chan generator defines
		// generatable channel's size
		// In this example channel will have size in range [0, 10]
		generator.Chan(constraints.Length{
			Min: 0,
			Max: 10,
		}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// chan size: 5
	// chan size: 2
	// chan size: 6
	// chan size: 0
	// chan size: 0
	// chan size: 6
	// chan size: 1
	// chan size: 5
	// chan size: 6
	// chan size: 4
}
