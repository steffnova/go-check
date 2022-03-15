package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/generator"
)

// This example demonstrates how to use Ptr(Int()) generator for generation of *int values.
// PtrTo requires a generator for type (int in this example) pointer points to, thus Int()
// generator is used.
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
	// -3087144572626463248
	// <nil>
	// -5254077479683016640
	// -1543285579645681342
	// <nil>
	// 2122761628320059770
	// -2396012503939351775
	// <nil>
	// -5365688832259816617
	// <nil>
}

// This example demonstrates how to use PtrTo(Int()) generator for generation of *int values.
// PtrTo requires a generator for type (int in this example) pointer points to, thus Int()
// generator is used.
func ExamplePtrTo() {
	streamer := generator.Streamer(
		func(n *int) {
			fmt.Printf("%v\n", *n)
		},
		generator.PtrTo(generator.Int()),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// -5339971465336467958
	// 5036824528102830934
	// 4435185786993720788
	// 8071137008395949086
	// 2122761628320059770
	// -5365688832259816617
	// -300681375570251064
	// -6485228379443441869
	// -8468275846115330281
	// -1089963290385541773
}
