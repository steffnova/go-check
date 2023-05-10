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
	// -31
	// -91
	// -40
	// 50
	// 79
	// -40
	// -116
	// -44
	// 102
	// -23
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
	// 12790
	// -13142
	// -5150
	// -3323
	// 3378
	// -4175
	// 5523
	// -5629
	// 31913
	// -20516
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
	// 826517535
	// 898515203
	// 649801091
	// -2023694845
	// -141136839
	// 1633988338
	// 2006390345
	// -16579444
	// -2004448480
	// -1121568663
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
	// 5
	// -4
	// -3
	// 0
	// 4
	// 1
	// 2
	// 3
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
	// -502
	// 854
	// 453
	// 259
	// 859
	// 81
	// 808
	// -616
	// 967
	// -185
}
