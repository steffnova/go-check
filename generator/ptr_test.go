package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/generator"
)

func ExamplePtr() {
	// Streamer uses Ptr generator to generate *int values.
	check.Stream(check.Streamer(
		func(n *int) {
			if n != nil {
				fmt.Printf("%v\n", *n)
			} else {
				fmt.Printf("%v\n", n)
			}
		},
		// Ptr generator requires a generator for type pointer
		// points to. In this case generator for int values is
		// used. Ptr will create nil or valid pointer
		generator.Ptr(generator.Int()),
	), check.Config{Seed: 0, Iterations: 10})

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

func ExamplePtrTo() {
	// Streamer uses Ptr generator to generate *int values.
	check.Stream(check.Streamer(
		func(n *int) {
			fmt.Printf("%v\n", *n)
		},
		// Unlike Ptr generator PtrTo will always create a non-nil pointer
		generator.PtrTo(generator.Int()),
	), check.Config{Seed: 0, Iterations: 10})

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
