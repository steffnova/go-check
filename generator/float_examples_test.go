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
	// -2.4597852e+30
	// -2.84718e-09
	// 3.9155332e-22
	// 5.336489e+16
	// 5.3546978e+23
	// -6.6470676e-23
	// 2.0516037e-32
	// 6.878372e+15
	// -2.5820768e+34
	// -7.02896e-34
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
	// -1.3131993411626801e-06
	// 2.6499394281062435e-102
	// -5.544832151478103e+28
	// -1.6925393061358905e+61
	// -1.3148118989728537e-70
	// -4.579610072238924e+81
	// -3.1950849453106405e+135
	// 1.6727313548317884e-205
	// 2.78249623188919e+86
	// 5.04589615041297e+124
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
