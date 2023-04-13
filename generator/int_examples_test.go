package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Int() generator for generation of int values.
func ExampleInt() {
	streamer := generator.Streamer(
		func(n int) {
			fmt.Printf("%d\n", n)
		},
		generator.Int(),
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

// This example demonstrates usage of Int() generator with constraints for generation of int values.
// Constraints defines range of generatable int values.
func ExampleInt_constraints() {
	streamer := generator.Streamer(
		func(n int) {
			fmt.Printf("%d\n", n)
		},
		generator.Int(constraints.Int{
			Min: 0,
			Max: 10,
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Int8() generator for generation of int8 values.
func ExampleInt8() {
	streamer := generator.Streamer(
		func(n int8) {
			fmt.Printf("%d\n", n)
		},
		generator.Int8(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Int8() generator with constraints for generation of int8 values.
// Constraints defines range of generatable int8 values.
func ExampleInt8_constraints() {
	streamer := generator.Streamer(
		func(n int8) {
			fmt.Printf("%d\n", n)
		},
		generator.Int8(constraints.Int8{
			Min: 100,
			Max: 127,
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Int16() generator for generation of int16 values.
func ExampleInt16() {
	streamer := generator.Streamer(
		func(n int16) {
			fmt.Printf("%d\n", n)
		},
		generator.Int16(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Int16() generator with constraints for generation of int16 values.
// Constraints defines range of generatable int16 values.
func ExampleInt16_constraints() {
	streamer := generator.Streamer(
		func(n int16) {
			fmt.Printf("%d\n", n)
		},
		generator.Int16(constraints.Int16{
			Min: -200,
			Max: -100,
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Int32() generator for generation of int32 values.
func ExampleInt32() {
	streamer := generator.Streamer(
		func(n int32) {
			fmt.Printf("%d\n", n)
		},
		generator.Int32(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Int32() generator with constraints for generation of int32 values.
// Constraints defines range of generatable int32 values.
func ExampleInt32_constraints() {
	streamer := generator.Streamer(
		func(n int32) {
			fmt.Printf("%d\n", n)
		},
		generator.Int32(constraints.Int32{
			Min: -5,
			Max: 5,
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Int64() generator for generation of int64 values.
func ExampleInt64() {
	streamer := generator.Streamer(
		func(n int64) {
			fmt.Printf("%d\n", n)
		},
		generator.Int64(),
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

// This example demonstrates usage of Int64() generator with constraints for generation of int64 values.
// Constraints defines range of generatable int64 values.
func ExampleInt64_constraints() {
	streamer := generator.Streamer(
		func(n int64) {
			fmt.Printf("%d\n", n)
		},
		generator.Int64(constraints.Int64{
			Min: -1000,
			Max: 1000,
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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
