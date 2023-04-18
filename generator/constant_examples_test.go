package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Constant() generator for generating string values.
func ExampleConstant() {
	streamer := generator.Streamer(
		func(s string) {
			fmt.Printf("%s\n", s)
		},
		generator.Constant("I won't write test cases"),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// I won't write test cases
	// I won't write test cases
	// I won't write test cases
	// I won't write test cases
	// I won't write test cases
	// I won't write test cases
	// I won't write test cases
	// I won't write test cases
	// I won't write test cases
	// I won't write test cases
}

// This example demonstrates usage of ConstantFrom() generator for generating string values.
func ExampleConstantFrom() {
	streamer := generator.Streamer(
		func(in interface{}) {
			fmt.Printf("%v\n", in)
		},
		generator.ConstantFrom("red", "green", "blue", "yellow", "black"),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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
