package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

func ExampleInt() {
	// Streamer uses Int generator to generate int values.
	check.Stream(check.Streamer(
		func(n int) {
			fmt.Printf("%d\n", n)
		},
		generator.Int(),
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

func ExampleInt_withConstraints() {
	// Streamer uses Int generator to generate int values.
	check.Stream(check.Streamer(
		func(n int) {
			fmt.Printf("%d\n", n)
		},
		// Passing constraint.Int to Int generator defines
		// range of generatable values [0, 10]
		generator.Int(constraints.Int{
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

func ExampleInt8() {
	// Streamer uses Int8 generator to generate int values.
	check.Stream(check.Streamer(
		func(n int8) {
			fmt.Printf("%d\n", n)
		},
		generator.Int8(),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// -16
	// -91
	// -81
	// 40
	// 104
	// -21
	// 122
	// -40
	// -116
	// 67
}

func ExampleInt8_withConstraints() {
	// Streamer uses Int8 generator to generate int8 values.
	check.Stream(check.Streamer(
		func(n int8) {
			fmt.Printf("%d\n", n)
		},
		// Passing constraint.Int8 to Int8 generator defines
		// range of generatable values [100, 127]
		generator.Int8(constraints.Int8{
			Min: 100,
			Max: 127,
		}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 121
	// 118
	// 122
	// 116
	// 116
	// 122
	// 117
	// 105
	// 122
	// 120
}

func ExampleInt16() {
	// Streamer uses Int16 generator to generate int values.
	check.Stream(check.Streamer(
		func(n int16) {
			fmt.Printf("%d\n", n)
		},
		generator.Int16(),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// -12790
	// 13142
	// 15828
	// -5150
	// 11651
	// -3323
	// 16168
	// -24597
	// -4175
	// 1247
}

func ExampleInt16_withConstraints() {
	// Streamer uses Int16 generator to generate int16 values.
	check.Stream(check.Streamer(
		func(n int16) {
			fmt.Printf("%d\n", n)
		},
		// Passing constraint.Int16 to Int16 generator defines
		// range of generatable values [-200, -100]
		generator.Int16(constraints.Int16{
			Min: -200,
			Max: -100,
		}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// -131
	// -116
	// -180
	// -186
	// -169
	// -122
	// -184
	// -103
	// -164
	// -130
}

func ExampleInt32() {
	// Streamer uses Int32 generator to generate int values.
	check.Stream(check.Streamer(
		func(n int32) {
			fmt.Printf("%d\n", n)
		},
		generator.Int32(),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 1530763030
	// -1726138304
	// 446740315
	// -1349338157
	// 2023694845
	// -141136839
	// 189385913
	// -1915228239
	// -2125661407
	// 986604357
}

func ExampleInt32_withConstraints() {
	// Streamer uses Int32 generator to generate int32 values.
	check.Stream(check.Streamer(
		func(n int32) {
			fmt.Printf("%d\n", n)
		},
		// Passing constraint.Int32 to Int32 generator defines
		// range of generatable values [-5, 5]
		generator.Int32(constraints.Int32{
			Min: -5,
			Max: 5,
		}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 2
	// 0
	// -1
	// 4
	// 0
	// -5
	// -3
	// 0
	// 2
	// 4
}

func ExampleInt64() {
	// Streamer uses Int64 generator to generate int values.
	check.Stream(check.Streamer(
		func(n int64) {
			fmt.Printf("%d\n", n)
		},
		generator.Int64(),
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

func ExampleInt64_withConstraints() {
	// Streamer uses Int64 generator to generate int64 values.
	check.Stream(check.Streamer(
		func(n int64) {
			fmt.Printf("%d\n", n)
		},
		// Passing constraint.Int64 to Int64 generator defines
		// range of generatable values [-1000, 1000]
		generator.Int64(constraints.Int64{
			Min: -1000,
			Max: 1000,
		}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// -31
	// -976
	// -497
	// 453
	// 468
	// -960
	// 859
	// 45
	// 251
	// -808
}
