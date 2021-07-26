package constraints

import "math"

type Float64 struct {
	Min float64
	Max float64
}

func Float64Default() Float64 {
	return Float64{
		Min: -math.MaxFloat64,
		Max: math.MaxFloat64,
	}
}

type Float32 struct {
	Min float32
	Max float32
}

func Float32Default() Float32 {
	return Float32{
		Min: -math.MaxFloat32,
		Max: math.MaxFloat32,
	}
}
