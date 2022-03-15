package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/generator"
)

// This example demonstrates how to use Nil() generator for generating nil values
// for channels, pointers, slices, maps and interfaces.
func ExampleNil() {
	streamer := generator.Streamer(
		func(p1 chan int, p2 *int, p3 []int, p4 map[string]int, p5 interface{}) {
			fmt.Printf("%#v, %#v, %#v, %#v, %#v\n", p1, p2, p3, p4, p5)
		},
		generator.Nil(),
		generator.Nil(),
		generator.Nil(),
		generator.Nil(),
		generator.Nil(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// (chan int)(nil), (*int)(nil), []int(nil), map[string]int(nil), <nil>
	// (chan int)(nil), (*int)(nil), []int(nil), map[string]int(nil), <nil>
	// (chan int)(nil), (*int)(nil), []int(nil), map[string]int(nil), <nil>
	// (chan int)(nil), (*int)(nil), []int(nil), map[string]int(nil), <nil>
	// (chan int)(nil), (*int)(nil), []int(nil), map[string]int(nil), <nil>
	// (chan int)(nil), (*int)(nil), []int(nil), map[string]int(nil), <nil>
	// (chan int)(nil), (*int)(nil), []int(nil), map[string]int(nil), <nil>
	// (chan int)(nil), (*int)(nil), []int(nil), map[string]int(nil), <nil>
	// (chan int)(nil), (*int)(nil), []int(nil), map[string]int(nil), <nil>
	// (chan int)(nil), (*int)(nil), []int(nil), map[string]int(nil), <nil>
}
