package property

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
	"github.com/steffnova/go-check/shrinker"
)

func TestInputsGeneratorFilter(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"PredicateInputsNumberAndTargetsMissmathch": func(t *testing.T) {
			r := arbitrary.RandomNumber{Rand: rand.New(rand.NewSource(0))}
			b := constraints.Bias{}

			generator := InputsGenerator(func(targets []reflect.Type, bias constraints.Bias, random arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
				return nil, nil, nil
			}).Filter(func(x int) bool {
				return false
			})

			targets := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)}

			if _, _, err := generator(targets, b, r); !errors.Is(err, ErrorInputs) {
				t.Fatalf("Expected error: %s", ErrorInputs)
			}
		},
		"PredicateInputsTypesAndTargetsMissmatch": func(t *testing.T) {
			r := arbitrary.RandomNumber{Rand: rand.New(rand.NewSource(0))}
			b := constraints.Bias{}

			generator := InputsGenerator(func(targets []reflect.Type, bias constraints.Bias, random arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
				return nil, nil, nil
			}).Filter(func(x int, y uint) bool {
				return false
			})

			targets := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)}

			if _, _, err := generator(targets, b, r); !errors.Is(err, ErrorInputs) {
				t.Fatalf("Expected error: %s", ErrorInputs)
			}
		},
		"GeneratorError": func(t *testing.T) {
			r := arbitrary.RandomNumber{Rand: rand.New(rand.NewSource(0))}
			b := constraints.Bias{}
			generatorError := fmt.Errorf("generator error")

			generator := InputsGenerator(func(targets []reflect.Type, bias constraints.Bias, random arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
				return nil, nil, generatorError
			}).Filter(func(x int, y int) bool {
				return false
			})

			targets := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)}

			if _, _, err := generator(targets, b, r); !errors.Is(err, generatorError) {
				t.Fatalf("Expected error: %s", generatorError)
			}
		},
		"GeneratedValueSatisfiesPredicate": func(t *testing.T) {
			r := arbitrary.RandomNumber{Rand: rand.New(rand.NewSource(0))}
			b := constraints.Bias{}
			arbs1 := arbitrary.Arbitraries{
				{Value: reflect.ValueOf(100)},
				{Value: reflect.ValueOf(200)},
			}

			generator := InputsGenerator(func(targets []reflect.Type, bias constraints.Bias, random arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
				return arbs1, nil, nil
			}).Filter(func(x int, y int) bool {
				return true
			})

			targets := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)}

			arbs2, _, err := generator(targets, b, r)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			for index := range arbs1 {
				parameter1 := arbs1[index].Value.Int()
				parameter2 := arbs2[index].Value.Int()
				if parameter1 != parameter2 {
					t.Fatalf("generated value %d at index %d doesn't match expected value: %d", index, parameter2, parameter1)
				}
			}

		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}

func TestInputsGeneratorLog(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"GeneratorError": func(t *testing.T) {
			r := arbitrary.RandomNumber{Rand: rand.New(rand.NewSource(0))}
			b := constraints.Bias{}
			generatorError := fmt.Errorf("generator error")

			generator := InputsGenerator(func(targets []reflect.Type, bias constraints.Bias, random arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
				return nil, nil, generatorError
			}).Log()

			targets := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)}

			if _, _, err := generator(targets, b, r); !errors.Is(err, generatorError) {
				t.Fatalf("Expected error: %s", generatorError)
			}
		},
		"Passthrough": func(t *testing.T) {
			r := arbitrary.RandomNumber{Rand: rand.New(rand.NewSource(0))}
			b := constraints.Bias{}
			arbs1 := arbitrary.Arbitraries{
				{Value: reflect.ValueOf(100)},
				{Value: reflect.ValueOf(200)},
			}

			generator := InputsGenerator(func(targets []reflect.Type, bias constraints.Bias, random arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
				return arbs1, nil, nil
			}).Log()

			targets := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)}

			arbs2, _, err := generator(targets, b, r)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			for index := range arbs1 {
				parameter1 := arbs1[index].Value.Int()
				parameter2 := arbs2[index].Value.Int()
				if parameter1 != parameter2 {
					t.Fatalf("generated value %d at index %d doesn't match expected value: %d", index, parameter2, parameter1)
				}
			}

		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}

func TestInputsGeneratorNoSrink(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"GeneratorError": func(t *testing.T) {
			r := arbitrary.RandomNumber{Rand: rand.New(rand.NewSource(0))}
			b := constraints.Bias{}
			generatorError := fmt.Errorf("generator error")

			generator := InputsGenerator(func(targets []reflect.Type, bias constraints.Bias, random arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
				return nil, nil, generatorError
			}).NoShrink()

			targets := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)}

			if _, _, err := generator(targets, b, r); !errors.Is(err, generatorError) {
				t.Fatalf("Expected error: %s", generatorError)
			}
		},
		"GeneratedArbitrariesDoNotHaveShrinker": func(t *testing.T) {
			r := arbitrary.RandomNumber{Rand: rand.New(rand.NewSource(0))}
			b := constraints.Bias{}
			arbs1 := arbitrary.Arbitraries{
				{
					Value:    reflect.ValueOf(uint64(100)),
					Shrinker: shrinker.Uint64(constraints.Uint64Default()),
				},
				{
					Value:    reflect.ValueOf(uint64(200)),
					Shrinker: shrinker.Uint64(constraints.Uint64Default()),
				},
			}

			generator := InputsGenerator(func(targets []reflect.Type, bias constraints.Bias, random arbitrary.Random) (arbitrary.Arbitraries, inputShrinker, error) {
				return arbs1, nil, nil
			}).NoShrink()

			targets := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)}

			arbs2, _, err := generator(targets, b, r)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			for index := range arbs1 {
				parameter1 := arbs1[index].Value.Uint()
				parameter2 := arbs2[index].Value.Uint()
				if parameter1 != parameter2 {
					t.Fatalf("generated value %d at index %d doesn't match expected value: %d", index, parameter2, parameter1)
				}
				if arbs2[index].Shrinker != nil {
					t.Fatalf("Generated arb at index %d has non-nil shrinker", index)
				}
			}

		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}

func TestInputs(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"InsufficientGenerators": func(t *testing.T) {
			r := arbitrary.RandomNumber{Rand: rand.New(rand.NewSource(0))}
			b := constraints.Bias{}
			targets := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)}

			generator := Inputs(generator.Int())
			if _, _, err := generator(targets, b, r); !errors.Is(err, ErrorInputs) {
				t.Fatalf("Expected error: %s", ErrorInputs)
			}
		},
		"GeneratorError": func(t *testing.T) {
			r := arbitrary.RandomNumber{Rand: rand.New(rand.NewSource(0))}
			b := constraints.Bias{}
			targets := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)}
			generatorError := fmt.Errorf("generator error")

			gen := Inputs(
				generator.Int(),
				arbitrary.Generator(func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
					return arbitrary.Arbitrary{}, generatorError
				}),
			)
			if _, _, err := gen(targets, b, r); !errors.Is(err, ErrorInputs) {
				t.Fatalf("Expected error: %s", ErrorInputs)
			}
		},
	}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}
}
