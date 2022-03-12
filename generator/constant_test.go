package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/generator"
)

func ExampleConstant() {
	// Streamer uses Constant generator to generate string values.
	// Constant will always return the same value that is passed to it.
	check.Stream(check.Streamer(
		func(s string) {
			fmt.Printf("%s\n", s)
		},
		generator.Constant("lorem ipsum"),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// lorem ipsum
	// lorem ipsum
	// lorem ipsum
	// lorem ipsum
	// lorem ipsum
	// lorem ipsum
	// lorem ipsum
	// lorem ipsum
	// lorem ipsum
	// lorem ipsum
}

func ExampleConstantFrom() {
	// Streamer uses ConstantFrom generator to generate string values.
	// ConstantFrom will generate one of the values passed to it.
	check.Stream(check.Streamer(
		func(s string) {
			fmt.Printf("%s\n", s)
		},
		generator.ConstantFrom("red", "green", "blue", "yellow", "black"),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// blue
	// red
	// red
	// green
	// black
	// yellow
	// red
	// yellow
	// yellow
	// green
}
