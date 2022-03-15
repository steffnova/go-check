package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Bool() generator for generation of bool values.
func ExampleBool() {
	streamer := generator.Streamer(
		func(b bool) {
			fmt.Printf("%v\n", b)
		},
		generator.Bool(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// true
	// false
	// false
	// true
	// true
	// false
	// false
	// true
	// true
	// false
}
