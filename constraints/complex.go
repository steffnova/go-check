package constraints

type Complex64 struct {
	Real      Float32
	Imaginary Float32
}

func Complex64Default() Complex64 {
	return Complex64{
		Real:      Float32Default(),
		Imaginary: Float32Default(),
	}
}

type Complex128 struct {
	Real      Float64
	Imaginary Float64
}

func Complex128Default() Complex128 {
	return Complex128{
		Real:      Float64Default(),
		Imaginary: Float64Default(),
	}
}
