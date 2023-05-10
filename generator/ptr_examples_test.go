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
	// -4518808235179270133
	// -322463145728654719
	// <nil>
	// -2122761628320059770
	// <nil>
	// 5606570076237929230
	// -868800358632342511
	// <nil>
	// -315038161257240872
	// -8800248522230157011
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
	// -4518808235179270133
	// -5036824528102830934
	// 8071137008395949086
	// -2122761628320059770
	// 6925466992496964832
	// 5606570076237929230
	// -6485228379443441869
	// -1089963290385541773
	// -315038161257240872
	// -8800248522230157011
}
