package constraints

type String struct {
	Rune   Rune
	Length Length
}

func StringDefault() String {
	return String{
		Rune:   RuneDefault(),
		Length: LengthDefault(),
	}
}
