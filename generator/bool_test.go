package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/generator"
)

func ExampleBool() {
	// Streamer uses Bool generator to generate bool values
	check.Stream(check.Streamer(
		func(b bool) {
			fmt.Printf("%#v\n", b)
		},
		generator.Bool(),
	), check.Config{Seed: 0, Iterations: 10})

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
