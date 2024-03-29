package arbitrary_test

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/generator"
)

func TestGeneratorMap(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"MapperIsNotAFunction": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(int) {},
				generator.Int().Map(nil),
			))

			if !errors.Is(err, arbitrary.ErrorMapper) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorMapper)
			}
		},
		"MapperInput": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(int) {},
				generator.Int().Map(func() int {
					return 0
				}),
			))

			if !errors.Is(err, arbitrary.ErrorMapper) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorMapper)
			}
		},
		"MapperOutput": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(int) {},
				generator.Int().Map(func(int) {}),
			))

			if !errors.Is(err, arbitrary.ErrorMapper) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorMapper)
			}
		},
		"MapperOutputType": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(int) {},
				generator.Int().Map(func(int) string {
					return ""
				}),
			))

			if !errors.Is(err, arbitrary.ErrorMapper) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorMapper)
			}
		},
		"InvalidTarget": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(string) {},
				generator.Uint().Map(func(int) string {
					return ""
				}),
			))
			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidTarget)
			}
		},
		"Mapping": func(t *testing.T) {
			err := generator.Stream(0, 100, generator.Streamer(
				func(n int) {
					if n%2 != 0 {
						t.Fatalf("Invalid mapped value: %d", n)
					}
				},
				generator.Int().Map(func(in int) int {
					return in * 2
				}),
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

func TestGeneratorFilter(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"PredicateIsNotAFunction": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(int) {},
				generator.Int().Filter(nil),
			))

			if !errors.Is(err, arbitrary.ErrorFilter) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorFilter)
			}
		},
		"PredicateInput": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(int) {},
				generator.Int().Filter(func() bool {
					return true
				}),
			))

			if !errors.Is(err, arbitrary.ErrorFilter) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorFilter)
			}
		},
		"PredicateOutput": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(int) {},
				generator.Int().Filter(func(int) {}),
			))

			if !errors.Is(err, arbitrary.ErrorFilter) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorFilter)
			}
		},
		"PredicateOutputType": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(int) {},
				generator.Int().Filter(func(in int) int {
					return 0
				}),
			))

			if !errors.Is(err, arbitrary.ErrorFilter) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorFilter)
			}
		},
		"InvalidTarget": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(string) {},
				generator.Int().Filter(func(int) bool {
					return false
				}),
			))
			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidTarget)
			}
		},
		"Filtering": func(t *testing.T) {
			err := generator.Stream(0, 100, generator.Streamer(
				func(n int) {
					if n%2 != 0 {
						t.Fatalf("Invalid filtered value: %d", n)
					}
				},
				generator.Int().Filter(func(in int) bool {
					return in%2 == 0
				}),
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

func TestGeneratorBind(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"BinderIsNotAFunction": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(int) {},
				generator.Int().Bind(0),
			))

			if !errors.Is(err, arbitrary.ErrorBinder) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorBinder)
			}
		},
		"BinderInput": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(int) {},
				generator.Int().Bind(func() arbitrary.Generator {
					return generator.Int()
				}),
			))

			if !errors.Is(err, arbitrary.ErrorBinder) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorBinder)
			}
		},
		"BinderOutput": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(int) {},
				generator.Int().Bind(func(int) {}),
			))

			if !errors.Is(err, arbitrary.ErrorBinder) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorBinder)
			}
		},
		"BinderOutputType": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(int) {},
				generator.Int().Bind(func(in int) int {
					return 0
				}),
			))

			if !errors.Is(err, arbitrary.ErrorBinder) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorBinder)
			}
		},
		"BounderTarget": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(string) {},
				generator.Int().Bind(func(int) arbitrary.Generator {
					return generator.Int()
				}),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidTarget)
			}
		},
		"InvalidTarget": func(t *testing.T) {
			err := generator.Stream(0, 10, generator.Streamer(
				func(int) {},
				generator.Int().Bind(func(uint) arbitrary.Generator {
					return generator.Int()
				}),
			))
			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidTarget)
			}
		},
		"Binding": func(t *testing.T) {
			err := generator.Stream(0, 100, generator.Streamer(
				func(n int) {
					if n%2 != 0 {
						t.Fatalf("Invalid filtered value: %d", n)
					}
				},
				generator.Int().Filter(func(in int) bool {
					return in%2 == 0
				}),
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
