package generator

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/constraints"
)

func TestComplex128(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			complexRange := constraints.Complex128{
				Real:      constraints.Float64{Min: -50, Max: 50},
				Imaginary: constraints.Float64{Min: -50, Max: 50},
			}
			Stream(0, 100, Streamer(
				func(n complex128) {
					switch {
					case real(n) < complexRange.Real.Min:
						fallthrough
					case real(n) > complexRange.Real.Max:
						fallthrough
					case imag(n) < complexRange.Imaginary.Min:
						fallthrough
					case imag(n) > complexRange.Imaginary.Max:
						t.Fatalf("Generated value is not withing given range: [%#v]", complexRange)
					}
				},
				Complex128(complexRange),
			))
		},
		"InvalidRealRange": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n complex128) {},
				Complex128(constraints.Complex128{
					Real:      constraints.Float64{Min: 50, Max: -50},
					Imaginary: constraints.Float64Default(),
				}),
			))
			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"InvalidImaginaryRange": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n complex128) {},
				Complex128(constraints.Complex128{
					Real:      constraints.Float64Default(),
					Imaginary: constraints.Float64{Min: 50, Max: -50},
				}),
			))
			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Complex128(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s':", ErrorInvalidTarget)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType complex128
			err := Stream(0, 100, Streamer(
				func(n newType) {},
				Complex128(),
			))
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}

func TestComplex64(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			complexRange := constraints.Complex64{
				Real:      constraints.Float32{Min: -50, Max: 50},
				Imaginary: constraints.Float32{Min: -50, Max: 50},
			}
			Stream(0, 100, Streamer(
				func(n complex64) {
					switch {
					case real(n) < complexRange.Real.Min:
						fallthrough
					case real(n) > complexRange.Real.Max:
						fallthrough
					case imag(n) < complexRange.Imaginary.Min:
						fallthrough
					case imag(n) > complexRange.Imaginary.Max:
						t.Fatalf("Generated value is not withing given range: [%#v]", complexRange)
					}
				},
				Complex64(complexRange),
			))
		},
		"InvalidRealRange": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n complex64) {},
				Complex64(constraints.Complex64{
					Real:      constraints.Float32{Min: 50, Max: -50},
					Imaginary: constraints.Float32Default(),
				}),
			))
			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"InvalidImaginaryRange": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n complex64) {},
				Complex64(constraints.Complex64{
					Real:      constraints.Float32Default(),
					Imaginary: constraints.Float32{Min: 50, Max: -50},
				}),
			))
			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Complex64(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s':", ErrorInvalidTarget)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType complex64
			err := Stream(0, 100, Streamer(
				func(n newType) {},
				Complex64(),
			))
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
