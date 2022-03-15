package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Float32() generator for generation of float32 values.
func ExampleFloat32() {
	streamer := generator.Streamer(
		func(f float32) {
			fmt.Printf("%g\n", f)
		},
		generator.Float32(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// 6.5711266e-15
	// -2.84718e-09
	// 1.841146e+06
	// 3.9155332e-22
	// 1.0390683e+06
	// -2.3398438e-37
	// -5.3546978e+23
	// 6.6470676e-23
	// -1.5922673e+10
	// 2.0516037e-32
}

// This example demonstrates usage of Float32() generator with constraints for generation of float32 values.
// Constraints defines range of generatable float32 values.
func ExampleFloat32_constraints() {
	streamer := generator.Streamer(
		func(f float32) {
			fmt.Printf("%g\n", f)
		},
		generator.Float32(constraints.Float32{
			Min: -2,
			Max: -1,
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// -1.2046497
	// -1.5285681
	// -1.6229705
	// -1.9818655
	// -1.4811733
	// -1.8536431
	// -1.6691408
	// -1.8248223
	// -1.5765601
	// -1.0456176
}

// This example demonstrates usage of Float64() generator for generation of float64 values.
func ExampleFloat64() {
	streamer := generator.Streamer(
		func(f float64) {
			fmt.Printf("%#v\n", f)
		},
		generator.Float64(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// -1.194033741351553e-241
	// -3.93536176617243e-117
	// -2.978088836427668e+188
	// -64716.894033756
	// -1.6925393061358905e+61
	// -3.290689787053985e-12
	// 1.8281685070362825e+43
	// -4.579610072238924e+81
	// 2.5420973754048248e+146
	// -2.9721426814134555e-286
}

// This example demonstrates usage of Float64() generator with constraints for generation of float64 values.
// Constraints defines range of generatable float64 values.
func ExampleFloat64_constraints() {
	streamer := generator.Streamer(
		func(f float64) {
			fmt.Printf("%#v\n", f)
		},
		generator.Float64(constraints.Float64{
			Min: 1,
			Max: 5,
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// 3.4237086631530884
	// 1.24048918956624
	// 2.967593950595308
	// 2.138899399920753
	// 3.365311348740657
	// 1.809075842134793
	// 1.1524283250776146
	// 4.827915588435498
	// 2.5447450908036644
	// 2.197932525132764
}
