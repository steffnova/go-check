package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("Fizz", func(tb *testing.T) {
		check.Check(tb, check.Property(
			func(in uint) error {
				result := fizzBuzz(in)
				if result != "Fizz" {
					return fmt.Errorf("%d did not return Fizz. Got: %s", in, result)
				}
				return nil
			},
			generator.Uint(constraints.Uint{Max: math.MaxUint / 3}).
				Filter(func(n uint) bool {
					return n%5 != 0
				}).
				Map(func(n uint) uint {
					return n * 3
				}),
		))
	})

	t.Run("Buzz", func(tb *testing.T) {
		check.Check(tb, check.Property(
			func(in uint) error {
				result := fizzBuzz(in)
				if fizzBuzz(in) != "Buzz" {
					return fmt.Errorf("%d did not return Buzz. Got: %s", in, result)
				}
				return nil
			},
			generator.Uint(constraints.Uint{Max: math.MaxUint / 5}).
				Filter(func(n uint) bool {
					return n%3 != 0
				}).
				Map(func(n uint) uint {
					return n * 5
				}),
		))
	})

	t.Run("Fizz Buzz", func(tb *testing.T) {
		check.Check(tb, check.Property(
			func(in uint) error {
				result := fizzBuzz(in)
				if fizzBuzz(in) != "Fizz Buzz" {
					return fmt.Errorf("%d did not return Fizz Buzz. Got %s", in, result)
				}
				return nil
			},
			generator.Uint(constraints.Uint{Max: math.MaxUint / 3 / 5}).
				Map(func(n uint) uint {
					return n * 3
				}).
				Map(func(n uint) uint {
					return n * 5
				}),
		))
	})

	t.Run("Regular Number", func(tb *testing.T) {
		check.Check(tb, check.Property(
			func(in uint) error {
				if fizzBuzz(in) != fmt.Sprintf("%d", in) {
					return fmt.Errorf("%d did not return %d", in, in)
				}
				return nil
			},
			generator.Uint().Filter(func(in uint) bool {
				return in%3 != 0 && in%5 != 0
			}),
		))
	})

}
