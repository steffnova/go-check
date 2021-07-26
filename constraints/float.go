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
