package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

func ExampleRune() {
	// Streamer uses Rune generator to generate rune values.
	check.Stream(check.Streamer(
		func(r rune) {
			fmt.Printf("%c\n", r)
		},
		generator.Rune(),
	), check.Config{Seed: 0, Iterations: 10})

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

func ExampleRune_withConstraints() {
	// Streamer uses Rune generator to generate rune values.
	check.Stream(check.Streamer(
		func(r rune) {
			fmt.Printf("%c\n", r)
		},
		// Passing constraint.Rune to Rune generator defines minimal and maximal
		// unicode code point for generated rune value.
		// In this example all rune value will be in range [a-z]
		generator.Rune(constraints.Rune{
			MinCodePoint: 'a',
			MaxCodePoint: 'z',
		}),
	), check.Config{Seed: 0, Iterations: 10})

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
