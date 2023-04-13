package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Chan() generator for generation of chan int values.
func ExampleChan() {
	streamer := generator.Streamer(
		func(ch chan int) {
			fmt.Printf("chan size: %d\n", cap(ch))
		},
		generator.Chan(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Chan() generator with constraints for generation of
// chan int values. Constraints define capacity of generated channel and in this example
// generated channel will have capacity in range [0, 10]
func ExampleChan_constraints() {
	streamer := generator.Streamer(
		func(ch chan int) {
			fmt.Printf("chan size: %d\n", cap(ch))
		},
		generator.Chan(constraints.Length{
			Min: 0,
			Max: 10,
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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
