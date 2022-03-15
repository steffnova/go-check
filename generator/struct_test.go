package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/generator"
)

// This example demonstrates how to use Struct() generator for generation of struct values.
// Struct() generator requires map[string]Generator, where map's key-value pairs represent
// struct's field generators. Provided field generator or Any() generator (if field generator
// is not provided) will be used to generate data for struct's field. In this example generators
// for fields X and Y are provided while for Z is ommited.
func ExampleStruct() {
	// Point struct will be used as struct example
	type Point struct {
		X int16
		Y int16
		Z int8
	}

	streamer := generator.Streamer(
		func(p Point) {
			fmt.Printf("%#v\n", p)
		},
		generator.Struct(map[string]generator.Generator{
			"X": generator.Int16(),
			"Y": generator.Int16(),
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// generator_test.Point{X:-12790, Y:13142, Z:-91}
	// generator_test.Point{X:-3323, Y:16168, Z:104}
	// generator_test.Point{X:-24597, Y:-4175, Z:40}
	// generator_test.Point{X:-5523, Y:5629, Z:-116}
	// generator_test.Point{X:-31913, Y:20516, Z:-57}
	// generator_test.Point{X:6926, Y:-17391, Z:23}
	// generator_test.Point{X:-27956, Y:13619, Z:127}
	// generator_test.Point{X:6440, Y:21880, Z:-12}
	// generator_test.Point{X:-20179, Y:28489, Z:10}
	// generator_test.Point{X:26139, Y:-3364, Z:7}
}
