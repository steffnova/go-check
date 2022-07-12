package main

import (
	"fmt"
)

func fizzBuzz(n uint) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "Fizz Buzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}
