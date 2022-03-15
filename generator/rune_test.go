package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates how to use Rune() generator for generation of rune values.
func ExampleRune() {
	streamer := generator.Streamer(
		func(r rune) {
			fmt.Printf("%c\n", r)
		},
		generator.Rune(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
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
	streamer := generator.Streamer(
		func(r rune) {
			fmt.Printf("%c\n", r)
		},
		generator.Rune(constraints.Rune{
			MinCodePoint: 'a',
			MaxCodePoint: 'z',
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
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
