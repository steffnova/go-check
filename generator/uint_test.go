package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Unt() generator for generation of uint values.
func ExampleUint() {
	streamer := generator.Streamer(
		func(n uint) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Int() generator with constraints for generation of int values.
// Constraints defines range of generatable int values.
func ExampleUint_constraints() {
	streamer := generator.Streamer(
		func(n uint) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint(constraints.Uint{
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

// This example demonstrates usage of Unt8() generator for generation of uint8 values.
func ExampleUint8() {
	streamer := generator.Streamer(
		func(n uint8) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint8(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Uint8() generator with constraints for generation of uint8 values.
// Constraints defines range of generatable uint8 values.
func ExampleUint8_constraints() {
	streamer := generator.Streamer(
		func(n uint8) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint8(constraints.Uint8{
			Min: 20,
			Max: 50,
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Unt16() generator for generation of uint16 values.
func ExampleUint16() {
	streamer := generator.Streamer(
		func(n uint16) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint16(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Uint16() generator with constraints for generation of int16 values.
// Constraints defines range of generatable int values.
func ExampleUint16_constraints() {
	streamer := generator.Streamer(
		func(n uint16) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint16(constraints.Uint16{
			Min: 100,
			Max: 500,
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Unt32() generator for generation of uint32 values.
func ExampleUint32() {
	streamer := generator.Streamer(
		func(n uint32) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint32(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Uint32() generator with constraints for generation of uint32 values.
// Constraints defines range of generatable int32 values.
func ExampleUint32_constraints() {
	streamer := generator.Streamer(
		func(n uint32) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint32(constraints.Uint32{
			Min: 10000,
			Max: 20000,
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Unt64() generator for generation of uint64 values.
func ExampleUint64() {
	streamer := generator.Streamer(
		func(n uint64) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint64(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Uint64() generator with constraints for generation of uint64 values.
// Constraints defines range of generatable int64 values.
func ExampleUint64_constraints() {
	streamer := generator.Streamer(
		func(n uint64) {
			fmt.Printf("%d\n", n)
		},
		generator.Uint64(constraints.Uint64{
			Min: 0,
			Max: 100,
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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
