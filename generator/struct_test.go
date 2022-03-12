package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/generator"
)

func ExampleStruct() {
	// Point struct will be used as an example
	type Point struct {
		X int16
		Y int16
		Z int8
	}

	// Streamer uses Struct generator to generate Point values.
	check.Stream(check.Streamer(
		func(p Point) {
			fmt.Printf("%#v\n", p)
		},
		// Struct generator accepts map where key is a string and value is
		// generator.Generator, where map's keys represent struct fields. If
		// generator for a field is not specified Any() generator is used.
		// In this example generators for fields X and Y are defined, while
		// for Z field it is ommited.
		generator.Struct(map[string]generator.Generator{
			"X": generator.Int16(),
			"Y": generator.Int16(),
		}),
	), check.Config{Seed: 0, Iterations: 10})

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
