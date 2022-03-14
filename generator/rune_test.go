package generator_test

import (
	"fmt"

	check "github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates how to use Rune() generator for generation of rune values.
func ExampleRune() {
	streamer := check.Streamer(
		func(r rune) {
			fmt.Printf("%c\n", r)
		},
		generator.Rune(),
	)

	if err := check.Stream(streamer, check.Config{Seed: 0, Iterations: 10}); err != nil {
		panic(err)
	}
	// Output:
	// 󋿲
	// 𺠟
	// 󎨐
	// 뿐
	// 򳍖
	// 󊷱
	// 󻵿
	// 󤄃
	// 𬟀
	// 띛
}

// This example demonstrates how to use Rune() generator with constraints for generation of rune values.
// Constraints define range of generatble rune values.
func ExampleRune_constraints() {
	streamer := check.Streamer(
		func(r rune) {
			fmt.Printf("%c\n", r)
		},
		generator.Rune(constraints.Rune{
			MinCodePoint: 'a',
			MaxCodePoint: 'z',
		}),
	)

	if err := check.Stream(streamer, check.Config{Seed: 0, Iterations: 10}); err != nil {
		panic(err)
	}
	// Output:
	// v
	// s
	// w
	// q
	// q
	// w
	// r
	// f
	// w
	// u
}
