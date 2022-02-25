package generator

import "github.com/steffnova/go-check/constraints"

// String is Arbitrary that creates string Generator. Range in which string's runes
// are generated, and string's length are defined by limits parameter. Even though
// limits is a variadic argument only the first value is used for defining. Error is
// returned if target's reflect.Kind is not String, if creation of underlaying slice
// fails.
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
