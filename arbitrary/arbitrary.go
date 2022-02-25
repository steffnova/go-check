package arbitrary

import "reflect"

// Arbitrary represents arbitrary type returned by shrinkers and generators.
type Arbitrary struct {
	Value      reflect.Value // Final value
	Elements   Arbitraries   // Arbitrary for each element in collection
	Precursors Arbitraries   // Precursor arbitraries from which this one is generated
}

type Arbitraries []Arbitrary

func (arbs Arbitraries) Values() []reflect.Value {
	out := make([]reflect.Value, len(arbs))
	for index, arb := range arbs {
		out[index] = arb.Value
	}
	return out
}
