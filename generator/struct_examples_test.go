package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates how to use Struct() generator for generation of struct values.
// Struct() generator requires map[string]arbitrary.Generator, where map's key-value pairs represent
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
		generator.Struct(map[string]arbitrary.Generator{
			"X": generator.Int16(),
			"Y": generator.Int16(),
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// generator_test.Point{X:12790, Y:-13142, Z:-30}
	// generator_test.Point{X:3323, Y:-16168, Z:50}
	// generator_test.Point{X:-4175, Y:5523, Z:116}
	// generator_test.Point{X:-29920, Y:-6926, Z:-23}
	// generator_test.Point{X:20130, Y:6440, Z:19}
	// generator_test.Point{X:20179, Y:-28489, Z:-10}
	// generator_test.Point{X:16896, Y:-743, Z:-95}
	// generator_test.Point{X:3072, Y:-7313, Z:-30}
	// generator_test.Point{X:20655, Y:29016, Z:81}
	// generator_test.Point{X:-7244, Y:12674, Z:7}
}
