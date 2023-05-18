package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates how to use the [Ptr] and [Int] generators to generate *int values.
// The Ptr generator requires a generator for the type that the pointer points to, so the Int
// generator is used in this case.
func ExamplePtr() {
	streamer := generator.Streamer(
		func(n *int) {
			if n != nil {
				fmt.Printf("%v\n", *n)
			} else {
				fmt.Printf("%v\n", n)
			}
		},
		generator.Ptr(generator.Int()),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// 3087144572626463248
	// -322463145728654719
	// 6634247955539956817
	// -2122761628320059770
	// 6925466992496964832
	// <nil>
	// 5606570076237929230
	// 5259823212710600989
	// -1089963290385541773
	// <nil>
}

// This example demonstrates how to use the [Ptr] and [Uint] generators in conjunction with
// [constraints.Ptr] to generate non-nil *uint values. By setting NilFrequency to 0, it ensures
// that nil pointers will never be generated. The [Ptr] generator requires a generator of the
// type to which the pointer points, so the Uint generator is used in this case.

func ExamplePtr_noNil() {
	streamer := generator.Streamer(
		func(n *uint) {
			fmt.Printf("%v\n", *n)
		},
		generator.Ptr(generator.Uint(), constraints.Ptr{NilFrequency: 0}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// 4518808235179270133
	// 3087144572626463248
	// 7861855757602232086
	// 12784885724210938115
	// 2119085704421221023
	// 1543285579645681342
	// 15398783846516204029
	// 9472434474353809100
	// 3877601997538530707
	// 16172318933975836041
}
