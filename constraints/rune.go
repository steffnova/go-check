package constraints

type Rune struct {
	MinCodePoint int32
	MaxCodePoint int32
}

func RuneDefault() Rune {
	return Rune{
		MinCodePoint: 0,
		MaxCodePoint: 0x10ffff,
	}
}
