package generator

import "github.com/steffnova/go-check/constraints"

// String returns generator for string types. Range of slice size is defined by
// "limits" parameter. If "limits" parameter is not specified default [0, 100]
// range is used instead. Error is returned if generator's target is not a
// string type, or limits.Min > limits.Max
func String(limits ...constraints.String) Generator {
	constraint := constraints.StringDefault()
	if len(limits) != 0 {
		constraint = limits[0]
	}

	return Slice(
		Rune(constraint.Rune),
		constraint.Length,
	).Map(func(runes []rune) string {
		return string(runes)
	})
}
