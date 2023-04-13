package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates how to use OneFrom() combinator and Int() generator to
// generate int values. OneFrom() will select one of the generators passed to it to
// generate new int value. In this example there are 4 Int() generators with different
// constraints passed to OneFrom() generator, and every time a int value needs to be
// generated one of them will be selected randomly.
func ExampleOneFrom() {
	streamer := generator.Streamer(
		func(n int) {
			fmt.Printf("%d\n", n)
		},
		generator.OneFrom(
			generator.Int(constraints.Int{Min: 500, Max: 1000}),
			generator.Int(constraints.Int{Min: 5, Max: 10}),
			generator.Int(constraints.Int{Min: -10, Max: -5}),
			generator.Int(constraints.Int{Min: -1000, Max: -500}),
		),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// -5
	// 842
	// 10
	// -948
	// -887
	// 8
	// 690
	// 592
	// -6
	// -10
}
