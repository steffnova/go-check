package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Weighted() combinator and Nil() and PtrTo(Uint64())
// generators for generation of *uint64 values. Weighted() will use one of the generators
// passed to it based on generator's weight. Weights define how often a generator will be
// selected by Weighted(). Selection chance is calculated as generator's weight devided by
// summ of all weights and multiplied by 100. In this example Nil() generator will have 10%
// selection chance (1/10 * 100) and PtrTo(Uint64()) will have 90% selection chance (9/10 * 100)
func ExampleWeighted() {
	streamer := generator.Streamer(
		func(n *uint64) {
			if n == nil {
				fmt.Printf("%v\n", n)
			} else {
				fmt.Printf("%d\n", *n)
			}
		},
		generator.Weighted(
			[]uint64{1, 9},
			generator.Nil(),
			generator.Ptr(generator.Uint64(), constraints.Ptr{NilFrequency: 0}),
		),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// 4518808235179270133
	// <nil>
	// 7861855757602232086
	// 5254077479683016640
	// 11116474692239114024
	// 15398783846516204029
	// 14677457169740829639
	// 9472434474353809100
	// 2396012503939351775
	// 3877601997538530707
}
