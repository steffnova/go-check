package generator

import (
	"errors"
	"testing"
)

func TestGeneratorMap(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"MapperIsNotAFunction": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Int().Map(nil),
			))

			if !errors.Is(err, ErrorMapper) {
				t.Fatalf("Expected error: '%s'", ErrorMapper)
			}
		},
		"MapperInput": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Int().Map(func() int {
					return 0
				}),
			))

			if !errors.Is(err, ErrorMapper) {
				t.Fatalf("Expected error: '%s'", ErrorMapper)
			}
		},
		"MapperOutput": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Int().Map(func(int) {}),
			))

			if !errors.Is(err, ErrorMapper) {
				t.Fatalf("Expected error: '%s'", ErrorMapper)
			}
		},
		"MapperOutputType": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Int().Map(func(int) string {
					return ""
				}),
			))

			if !errors.Is(err, ErrorMapper) {
				t.Fatalf("Expected error: '%s'", ErrorMapper)
			}
		},
		"InvalidTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(string) {},
				Uint().Map(func(int) string {
					return ""
				}),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidTarget)
			}
		},
		"Mapping": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(n int) {
					if n%2 != 0 {
						t.Fatalf("Invalid mapped value: %d", n)
					}
				},
				Int().Map(func(in int) int {
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
			err := Stream(0, 10, Streamer(
				func(int) {},
				Int().Filter(nil),
			))

			if !errors.Is(err, ErrorFilter) {
				t.Fatalf("Expected error: '%s'", ErrorFilter)
			}
		},
		"PredicateInput": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Int().Filter(func() bool {
					return true
				}),
			))

			if !errors.Is(err, ErrorFilter) {
				t.Fatalf("Expected error: '%s'", ErrorFilter)
			}
		},
		"PredicateOutput": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Int().Filter(func(int) {}),
			))

			if !errors.Is(err, ErrorFilter) {
				t.Fatalf("Expected error: '%s'", ErrorFilter)
			}
		},
		"PredicateOutputType": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Int().Filter(func(in int) int {
					return 0
				}),
			))

			if !errors.Is(err, ErrorFilter) {
				t.Fatalf("Expected error: '%s'", ErrorFilter)
			}
		},
		"InvalidTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(string) {},
				Int().Filter(func(int) bool {
					return false
				}),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidTarget)
			}
		},
		"Filtering": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(n int) {
					if n%2 != 0 {
						t.Fatalf("Invalid filtered value: %d", n)
					}
				},
				Int().Filter(func(in int) bool {
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
			err := Stream(0, 10, Streamer(
				func(int) {},
				Int().Bind(0),
			))

			if !errors.Is(err, ErrorBinder) {
				t.Fatalf("Expected error: '%s'", ErrorBinder)
			}
		},
		"BinderInput": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Int().Bind(func() Generator {
					return Int()
				}),
			))

			if !errors.Is(err, ErrorBinder) {
				t.Fatalf("Expected error: '%s'", ErrorBinder)
			}
		},
		"BinderOutput": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Int().Bind(func(int) {}),
			))

			if !errors.Is(err, ErrorBinder) {
				t.Fatalf("Expected error: '%s'", ErrorBinder)
			}
		},
		"BinderOutputType": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Int().Bind(func(in int) int {
					return 0
				}),
			))

			if !errors.Is(err, ErrorBinder) {
				t.Fatalf("Expected error: '%s'", ErrorBinder)
			}
		},
		"BounderTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(string) {},
				Int().Bind(func(int) Generator {
					return Int()
				}),
			))

			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidTarget)
			}
		},
		"InvalidTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Int().Bind(func(uint) Generator {
					return Int()
				}),
			))
			if !errors.Is(err, ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", ErrorInvalidTarget)
			}
		},
		"Binding": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(n int) {
					if n%2 != 0 {
						t.Fatalf("Invalid filtered value: %d", n)
					}
				},
				Int().Filter(func(in int) bool {
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
