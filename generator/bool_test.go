package generator_test

import (
	"fmt"

	check "github.com/steffnova/go-check"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Bool() generator for generation of bool values.
func ExampleBool() {
	streamer := check.Streamer(
		func(b bool) {
			fmt.Printf("%v\n", b)
		},
		generator.Bool(),
	)

	if err := check.Stream(streamer, check.Config{Seed: 0, Iterations: 10}); err != nil {
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
