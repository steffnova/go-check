package generator

import (
	"github.com/steffnova/go-check/arbitrary"
)

// Recurse is a type that is provided by the [Recursive] generator when
// defining a generator with recursion.
type Recurse func() arbitrary.Generator

// Recursive can be used to define recursive types (structures and functions). The 'genFunc' parameter
// is a function that provides a [Recurse] function, which can be used to specify a recursive call of
// the generator returned by 'genFunc'. The 'depth' parameter controls how deep the recursion goes
// (a value of n will cause n recursions).
func Recursive(genFunc func(Recurse) arbitrary.Generator, depth uint) arbitrary.Generator {
	return genFunc(func() arbitrary.Generator {
		if depth == 0 {
			return zeroValue()
		}
		return Recursive(genFunc, depth-1)
	})
}
