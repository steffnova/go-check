package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

func ExampleMap() {
	// Streamer uses Map generator to generate map[int8]bool.
	check.Stream(check.Streamer(
		func(m map[int8]bool) {
			fmt.Printf("%#v\n", m)
		},
		// Map generator requires generator for it's keys
		// and for it's values. In this case function requires
		// int8 generator for map keys and bool generator for
		// map values
		generator.Map(
			generator.Int8(),
			generator.Bool(),
		),
	), check.Config{Seed: 0, Iterations: 1})

	// Output:
	// map[int8]bool{-128:false, -116:true, -100:false, -91:true, -86:true, -76:false, -68:false, -64:true, -62:true, -57:false, -54:true, -51:true, -40:false, -38:true, -21:true, -11:false, 0:true, 4:false, 7:false, 10:false, 19:true, 23:true, 25:false, 30:true, 67:false, 73:false, 76:false, 120:false, 122:true, 126:true, 127:false}
}

func ExampleMap_withConstraints() {
	// Streamer uses Map generator to generate map values.
	check.Stream(check.Streamer(
		func(m map[int8]uint8) {
			fmt.Printf("%#v\n", m)
		},
		generator.Map(
			generator.Int8(),
			generator.Uint8(),
			// constraints define map length. Maps generated will
			// have a length in range [0, 5]
			constraints.Length{Min: 0, Max: 5},
		),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// map[int8]uint8{-92:0x15, -81:0xfb, -16:0x1e, 40:0xbe, 122:0x4f}
	// map[int8]uint8{40:0x49, 67:0x24}
	// map[int8]uint8{-102:0xcd, -57:0x1e, 23:0x5d}
	// map[int8]uint8{-76:0x8b, 120:0x13, 127:0x56}
	// map[int8]uint8{-64:0x1b, -54:0x71, -36:0xaa, -4:0x84, 10:0x50}
	// map[int8]uint8{-126:0x91, 19:0xbd, 30:0x24}
	// map[int8]uint8{25:0xb8}
	// map[int8]uint8{-128:0xfe, -62:0x51, 76:0x16}
	// map[int8]uint8{}
	// map[int8]uint8{-116:0x71, 36:0xc5}
}
