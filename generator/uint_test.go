package generator

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/constraints"
)

func TestUint64(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			uintRange := constraints.Uint64{Min: 20, Max: 100}
			Stream(0, 100, Streamer(
				func(n uint64) {
					if n < uintRange.Min || n > uintRange.Max {
						t.Fatalf("arbitrary.Arbitraryd value is not withing given range: [%d, %d]", uintRange.Min, uintRange.Max)
					}
				},
				Uint64(uintRange),
			))
		},
		"InvalidRange": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n uint64) {},
				Uint64(constraints.Uint64{Min: 100, Max: 20}),
			))
			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidConstraints)
			}
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Uint64(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s' because constraints are invalid.", ErrorInvalidTarget)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType uint64
			err := Stream(0, 100, Streamer(
				func(n newType) {},
				Uint64(),
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

func TestUint32(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			uintRange := constraints.Uint32{Min: 50, Max: 100}
			Stream(0, 100, Streamer(
				func(n uint32) {
					if n < uintRange.Min || n > uintRange.Max {
						t.Fatalf("arbitrary.Arbitraryd value is not withing given range: [%d, %d]", uintRange.Min, uintRange.Max)
					}
				},
				Uint32(uintRange),
			))
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Uint32(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error because constraints are invalid: %s", err)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType uint32
			err := Stream(0, 10, Streamer(
				func(n newType) {},
				Uint32(),
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

func TestUint16(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			uintRange := constraints.Uint16{Min: 50, Max: 100}
			Stream(0, 100, Streamer(
				func(n uint16) {
					if n < uintRange.Min || n > uintRange.Max {
						t.Fatalf("arbitrary.Arbitraryd value is not withing given range: [%d, %d]", uintRange.Min, uintRange.Max)
					}
				},
				Uint16(uintRange),
			))
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Uint16(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error because constraints are invalid: %s", err)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType uint16
			err := Stream(0, 100, Streamer(
				func(n newType) {},
				Uint16(),
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

func TestUint8(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			uintRange := constraints.Uint8{Min: 50, Max: 100}
			Stream(0, 100, Streamer(
				func(n uint8) {
					if n < uintRange.Min || n > uintRange.Max {
						t.Fatalf("arbitrary.Arbitraryd value is not withing given range: [%d, %d]", uintRange.Min, uintRange.Max)
					}
				},
				Uint8(uintRange),
			))
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Uint8(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error because constraints are invalid: %s", err)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType uint8
			err := Stream(0, 100, Streamer(
				func(n newType) {},
				Uint8(),
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

func TestUint(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			uintRange := constraints.Uint{Min: 10, Max: 50}
			Stream(0, 100, Streamer(
				func(n uint) {
					if n < uintRange.Min || n > uintRange.Max {
						t.Fatalf("arbitrary.Arbitraryd value is not withing given range: [%d, %d]", uintRange.Min, uintRange.Max)
					}
				},
				Uint(uintRange),
			))
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Uint(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error because constraints are invalid: %s", err)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType uint
			err := Stream(0, 100, Streamer(
				func(n newType) {},
				Uint(),
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
