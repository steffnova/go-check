package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

func ExampleFloat32() {
	// Streamer uses Float32 generator to generate float32 values.
	check.Stream(check.Streamer(
		func(f float32) {
			fmt.Printf("%#v\n", f)
		},
		generator.Float32(),
	), check.Config{Seed: 0, Iterations: 10})

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

func ExampleFloat32_withConstraints() {
	// Streamer uses Float32 generator to generate float32 values.
	check.Stream(check.Streamer(
		func(f float32) {
			fmt.Printf("%#v\n", f)
		},
		// Passing constraint.Float32 to Float32 generator defines
		// range of generatable values [5, 10]
		generator.Float32(constraints.Float32{Min: 5, Max: 10}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 5.8185987
	// 7.1142726
	// 7.491882
	// 9.854924
	// 6.924693
	// 8.8291445
	// 7.6765633
	// 8.598578
	// 7.3062406
	// 5.1824703
}

func ExampleFloat64() {
	check.Stream(check.Streamer(
		func(f float64) {
			fmt.Printf("%#v\n", f)
		},
		generator.Float64(),
	), check.Config{Seed: 0, Iterations: 10})

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

func ExampleFloat64_withConstraints() {
	// Streamer uses Float64 generator to generate float64 values.
	check.Stream(check.Streamer(
		func(f float64) {
			fmt.Printf("%#v\n", f)
		},
		// Passing constraint.Float64 to Float64 generator defines
		// range of generatable values [5, 10]
		generator.Float64(constraints.Float64{Min: 5, Max: 10}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// 7.369467376542149
	// 5.96195675826496
	// 6.598849677201391
	// 9.80002295745996
	// 6.266348980232071
	// 8.472606737078344
	// 8.519197152724251
	// 7.558232886878557
	// 5.6097133003104585
	// 5.827915588435498
}
