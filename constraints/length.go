package constraints

type Length struct {
	Min uint64
	Max uint64
}

func LengthDefault() Length {
	return Length{
		Min: 0,
		Max: 100,
	}
}
