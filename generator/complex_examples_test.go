package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Complex64() generator for generation of complex64 values.
func ExampleComplex64() {
	streamer := generator.Streamer(
		func(c complex64) {
			fmt.Printf("%g\n", c)
		},
		generator.Complex64(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Complex64() generator with constraints for generation of complex64 values.
// Constraints define range of generatable float32 values for real and imaginary parts of complex64 number.
func ExampleComplex64_constraints() {
	streamer := generator.Streamer(
		func(c complex64) {
			fmt.Printf("%g\n", c)
		},
		generator.Complex64(constraints.Complex64{
			Real: constraints.Float32{
				Min: -3,
				Max: -1,
			},
			Imaginary: constraints.Float32{
				Min: 3,
				Max: 5,
			},
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// (-1.2046497+4.1142726i)
	// (-2.701786+3.1750083i)
	// (-1.6229705+3.9623466i)
	// (-2.4881487+3.222718i)
	// (-2.5111911+3.9246528i)
	// (-1.8536431+4.56498i)
	// (-2.5411756+3.1288548i)
	// (-1.6691408+3.4864495i)
	// (-1.8248223+3.6883054i)
	// (-1.5765601+3.0912352i)
}

// This example demonstrates usage of Complex128() generator for generation of complex128 values.
func ExampleComplex128() {
	streamer := generator.Streamer(
		func(c complex128) {
			fmt.Printf("%g\n", c)
		},
		generator.Complex128(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
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

// This example demonstrates usage of Complex128() generator with constraints for generation of complex128 values.
// Constraints define range of generatable float64 values for real and imaginary parts of complex128 number.
func ExampleComplex128_constraints() {
	streamer := generator.Streamer(
		func(c complex128) {
			fmt.Printf("%g\n", c)
		},
		generator.Complex128(constraints.Complex128{
			Real: constraints.Float64{
				Min: 10,
				Max: 20,
			},
			Imaginary: constraints.Float64{
				Min: -200,
				Max: -100,
			},
		}),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// (14.738934753084298-115.39130813223936i)
	// (13.197699354402783-196.80036731935937i)
	// (12.532697960464143-175.5617077932535i)
	// (17.038394305448502-153.8634523801138i)
	// (11.219426600620917-113.24664941496796i)
	// (13.31424363947598-140.03814918293662i)
	// (12.825849578744581-158.8143128221763i)
	// (11.81996096177609-107.61529798063697i)
	// (11.084114411123993-115.55906892458002i)
	// (11.405528609648313-101.38558928949668i)
}
