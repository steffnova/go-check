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
	// (-2.4597852e+30-2.84718e-09i)
	// (3.9155332e-22+5.336489e+16i)
	// (5.3546978e+23-6.6470676e-23i)
	// (2.0516037e-32+6.878372e+15i)
	// (-2.5820768e+34-7.02896e-34i)
	// (0.42347318-3.3288446e+30i)
	// (0.0015746137-1.0883707e+20i)
	// (-7812.692+5.0595798e+33i)
	// (-1.0232254e+22-1.7621445e-22i)
	// (-6.7646325e-35-1.2282892e-32i)
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
	// (-1.3131993411626801e-06+2.6499394281062435e-102i)
	// (-5.544832151478103e+28-1.6925393061358905e+61i)
	// (-1.3148118989728537e-70-4.579610072238924e+81i)
	// (-3.1950849453106405e+135+1.6727313548317884e-205i)
	// (2.78249623188919e+86+5.04589615041297e+124i)
	// (9.142337453567358e-167-1.374619546296366e+49i)
	// (-4.1757521047374676e+224+4.010186513222081e+152i)
	// (6.556580829938433e+219+9.452945816540518e+154i)
	// (4.151349376347527e-207+4.786664746486726e+126i)
	// (-2.872024231170062e+70+1.8272776570354314e-65i)
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
