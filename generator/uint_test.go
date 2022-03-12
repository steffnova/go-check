package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

func ExampleUint() {
	// Streamer uses Uint generator to generate int values.
	check.Stream(check.Streamer(
		func(n uint) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint(),
	), check.Config{Seed: 0, Iterations: 10})

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

func ExampleUint_withConstraints() {
	// Streamer uses Uint generator to generate uint values.
	check.Stream(check.Streamer(
		func(n uint) {
			fmt.Printf("%d\n", n)
		},
		// Passing constraint.Unt to Uint generator defines
		// range of generatable values [0, 10]
		generator.Uint(constraints.Uint{
			Min: 0,
			Max: 10,
		}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 5
	// 2
	// 6
	// 0
	// 0
	// 6
	// 1
	// 5
	// 6
	// 4
}

func ExampleUint8() {
	// Streamer uses Uint8 generator to generate int8 values.
	check.Stream(check.Streamer(
		func(n uint8) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint8(),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 31
	// 16
	// 30
	// 45
	// 81
	// 251
	// 159
	// 190
	// 104
	// 92
}

func ExampleUint8_withConstraints() {
	// Streamer uses Uint8 generator to generate uint8 values.
	check.Stream(check.Streamer(
		func(n uint8) {
			fmt.Printf("%d\n", n)
		},
		// Passing constraint.Unt8 to Uint8 generator defines
		// range of generatable values [20, 50]
		generator.Uint8(constraints.Uint8{
			Min: 20,
			Max: 50,
		}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 41
	// 38
	// 42
	// 36
	// 36
	// 42
	// 37
	// 25
	// 42
	// 40
}

func ExampleUint16() {
	// Streamer uses Int16 generator to generate int16 values.
	check.Stream(check.Streamer(
		func(n uint16) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint16(),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 24565
	// 49138
	// 12790
	// 59920
	// 49104
	// 44529
	// 16643
	// 51136
	// 46939
	// 5201
}

func ExampleUint16_withConstraints() {
	// Streamer uses Uint16 generator to generate uint16 values.
	check.Stream(check.Streamer(
		func(n uint16) {
			fmt.Printf("%d\n", n)
		},
		// Passing constraint.Unt16 to Uint16 generator defines
		// range of generatable values [100, 500]
		generator.Uint16(constraints.Uint16{
			Min: 100,
			Max: 500,
		}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 131
	// 116
	// 442
	// 483
	// 378
	// 359
	// 130
	// 447
	// 487
	// 145
}

func ExampleUint32() {
	// Streamer uses Int32 generator to generate int32 values.
	check.Stream(check.Streamer(
		func(n uint32) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint32(),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 3649778518
	// 2615979505
	// 1530763030
	// 898515203
	// 2935165982
	// 649801091
	// 2329218299
	// 2225250975
	// 3653467838
	// 2023694845
}

func ExampleUint32_withConstraints() {
	// Streamer uses Uint32 generator to generate uint32 values.
	check.Stream(check.Streamer(
		func(n uint32) {
			fmt.Printf("%d\n", n)
		},
		// Passing constraint.Unt32 to Uint32 generator defines
		// range of generatable values [100, 500]
		generator.Uint32(constraints.Uint32{
			Min: 10000,
			Max: 20000,
		}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 18181
	// 15910
	// 10259
	// 11984
	// 15150
	// 11069
	// 15201
	// 13323
	// 18863
	// 19832
}

func ExampleUint64() {
	// Streamer uses Int64 generator to generate int64 values.
	check.Stream(check.Streamer(
		func(n uint64) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint64(),
	), check.Config{Seed: 0, Iterations: 10})

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

func ExampleUint64_withConstraints() {
	// Streamer uses Uint64 generator to generate uint64 values.
	check.Stream(check.Streamer(
		func(n uint64) {
			fmt.Printf("%d\n", n)
		},
		// Passing constraint.Unt64 to Uint64 generator defines
		// range of generatable values [100, 500]
		generator.Uint64(constraints.Uint64{
			Min: 0,
			Max: 100,
		}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 31
	// 16
	// 80
	// 86
	// 69
	// 22
	// 84
	// 3
	// 64
	// 30
}
