package generator

import (
	"errors"
	"testing"
)

func TestAny(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"InvalidTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(uintptr) {},
				Any(),
			))

			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidTarget)
			}
		},
		"Array": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func([10]int) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Bool": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(bool) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Complex64": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(complex64) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Complex128": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(complex128) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Chan": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(chan int) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Float32": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(float32) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Float64": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(float64) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Int": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(int) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Int8": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(int8) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Int16": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(int16) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Int32": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(int32) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Int64": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(int64) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Uint": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(uint) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Uint8": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(uint8) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Uint16": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(uint16) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Uint32": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(uint32) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Uint64": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(uint64) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Func": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(func(int) string) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Map": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(map[string]int) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Ptr": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(*int) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Struct": func(t *testing.T) {
			type st struct {
				a int
				b string
				c bool
			}
			err := Stream(0, 100, Streamer(
				func(st) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"Slice": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func([]int) {},
				Any(),
			))

			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		},
		"String": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(string) {},
				Any(),
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
