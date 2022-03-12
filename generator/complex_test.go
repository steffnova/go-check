package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

func ExampleComplex64() {
	// Streamer uses Complex64 generator to generate complex64 values
	check.Stream(check.Streamer(
		func(c complex64) {
			fmt.Printf("%#v\n", c)
		},
		generator.Complex64(),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// (6.5711266e-15-2.84718e-09i)
	// (1.841146e+06+3.9155332e-22i)
	// (1.0390683e+06-2.3398438e-37i)
	// (-5.3546978e+23+6.6470676e-23i)
	// (-1.5922673e+10+2.0516037e-32i)
	// (17.030838+1.0331481e+27i)
	// (-1.9262444e-32-7.02896e-34i)
	// (3.8865208e-32+0.42347318i)
	// (3.3288446e+30-1.1897855e+38i)
	// (0.0015746137+1.0883707e+20i)
}

func ExampleComplex64_withConstraints() {
	// Streamer uses Complex64 generator to generate complex64 values.
	check.Stream(check.Streamer(
		func(c complex64) {
			fmt.Printf("%#v\n", c)
		},
		// Passing constraint.Complex64 to Coplex64 generator defines
		// range of values for Real and Imaginary parts of complex number.
		generator.Complex64(constraints.Complex64{
			// Real part of complex64 will be from range [0,5]
			Real: constraints.Float32{
				Min: 0,
				Max: 5,
			},
			// Imaginary part of complex64 will be from range [-5,0]
			Imaginary: constraints.Float32{
				Min: -5,
				Max: 0,
			},
		}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// (6.5711266e-15-8.163024e-21i)
	// (2.84718e-09-2.999585e-16i)
	// (3.9155332e-22-3.9071593e-35i)
	// (2.3398438e-37-1.0598745e-06i)
	// (1.1053934e-10-6.6470676e-23i)
	// (1.2988068e-15-4.679253e-29i)
	// (3.644296e-16-2.0516037e-32i)
	// (3.823536e-36-5.0049135e-38i)
	// (1.9262444e-32-2.5963268e-11i)
	// (7.02896e-34-6.9491065e-26i)
}

func ExampleComplex128() {
	// Streamer uses Complex128 generator to generate complex128 values
	check.Stream(check.Streamer(
		func(c complex128) {
			fmt.Printf("%#v\n", c)
		},
		generator.Complex128(),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// (-1.194033741351553e-241-3.93536176617243e-117i)
	// (-2.978088836427668e+188-64716.894033756i)
	// (-1.6925393061358905e+61-3.290689787053985e-12i)
	// (1.8281685070362825e+43-4.579610072238924e+81i)
	// (2.5420973754048248e+146-2.9721426814134555e-286i)
	// (-4.07647343069254e-182-2.497850090921009e+236i)
	// (2.78249623188919e+86+4.120048535147697e+56i)
	// (2.33038011922896e-289+9.142337453567358e-167i)
	// (1.374619546296366e+49+1.5979905149546695e-148i)
	// (-7.908173283514067e+37-5.433575384625165e+75i)
}

func ExampleComplex128_withConstraints() {
	// Streamer uses Complex128 generator to generate complex128 values
	check.Stream(check.Streamer(
		func(c complex128) {
			fmt.Printf("%#v\n", c)
		},
		// Passing constraint.Complex128 to Coplex128 generator defines
		// range of values for Real and Imaginary parts of complex number.
		generator.Complex128(constraints.Complex128{
			// Real part of complex128 will be 0
			Real: constraints.Float64{
				Min: 0,
				Max: 0,
			},
			// Imaginary part of complex128 will be from range [-5,0]
			Imaginary: constraints.Float64{
				Min: -5,
				Max: 0,
			},
		}),
	), check.Config{Seed: 0, Iterations: 10})

	// Output:
	// (0-5.566863131315757e-260i)
	// (0-1.6566169045618332e-120i)
	// (0-3.5999967279572785e-304i)
	// (0-3.290689787053985e-12i)
	// (0-1.9905130526070908e-77i)
	// (0-1.6727313548317884e-205i)
	// (0-3.881227390501676e-204i)
	// (0-2.332451080513663e-43i)
	// (0-2.33038011922896e-289i)
	// (0-9.142337453567358e-167i)
}
