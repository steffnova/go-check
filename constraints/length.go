package constraints

type Length struct {
	Min int
	Max int
}

func LengthDefault() Length {
	return Length{
		Min: 0,
		Max: 100,
	}
}
