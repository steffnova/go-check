package generator

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/constraints"
)

func TestInt64(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			intRange := constraints.Int64{Min: -50, Max: 50}
			Stream(0, 100, Streamer(
				func(n int64) {
					if n < intRange.Min || n > intRange.Max {
						t.Fatalf("Generated value is not withing given range: [%d, %d]", intRange.Min, intRange.Max)
					}
				},
				Int64(intRange),
			))
		},
		"InvalidRange": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n int64) {},
				Int64(constraints.Int64{Min: 50, Max: -50}),
			))
			if !errors.Is(err, ErrorInvalidConstraints) {
				t.Fatalf("Expected error: %s", ErrorInvalidConstraints)
			}
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Int64(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error because constraints are invalid: %s", err)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType int64
			err := Stream(0, 100, Streamer(
				func(n newType) {},
				Int64(),
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

func TestInt32(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			intRange := constraints.Int32{Min: -50, Max: 50}
			Stream(0, 100, Streamer(
				func(n int32) {
					if n < intRange.Min || n > intRange.Max {
						t.Fatalf("Generated value is not withing given range: [%d, %d]", intRange.Min, intRange.Max)
					}
				},
				Int32(intRange),
			))
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Int32(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error because constraints are invalid: %s", err)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType int32
			err := Stream(0, 10, Streamer(
				func(n newType) {},
				Int32(),
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

func TestInt16(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			intRange := constraints.Int16{Min: -50, Max: 50}
			Stream(0, 100, Streamer(
				func(n int16) {
					if n < intRange.Min || n > intRange.Max {
						t.Fatalf("Generated value is not withing given range: [%d, %d]", intRange.Min, intRange.Max)
					}
				},
				Int16(intRange),
			))
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Int16(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error because constraints are invalid: %s", err)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType int16
			err := Stream(0, 100, Streamer(
				func(n newType) {},
				Int16(),
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

func TestInt8(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			intRange := constraints.Int8{Min: -50, Max: 50}
			Stream(0, 100, Streamer(
				func(n int8) {
					if n < intRange.Min || n > intRange.Max {
						t.Fatalf("Generated value is not withing given range: [%d, %d]", intRange.Min, intRange.Max)
					}
				},
				Int8(intRange),
			))
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Int8(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error because constraints are invalid: %s", err)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType int8
			err := Stream(0, 100, Streamer(
				func(n newType) {},
				Int8(),
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

func TestInt(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"WithinRange": func(t *testing.T) {
			intRange := constraints.Int{Min: -50, Max: 50}
			Stream(0, 100, Streamer(
				func(n int) {
					if n < intRange.Min || n > intRange.Max {
						t.Fatalf("Generated value is not withing given range: [%d, %d]", intRange.Min, intRange.Max)
					}
				},
				Int(intRange),
			))
		},
		"InvalidType": func(t *testing.T) {
			err := Stream(0, 1, Streamer(
				func(n string) {},
				Int(),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error because constraints are invalid: %s", err)
			}
		},
		"UnderlyingType": func(t *testing.T) {
			type newType int
			err := Stream(0, 100, Streamer(
				func(n newType) {},
				Int(),
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
