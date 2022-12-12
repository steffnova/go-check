package constraints

import "math"

// Bias contains Size and Scaling properties which are used to scale constraint limits.
// Scaling factor is calculated by diving Size with Scaling
type Bias struct {
	Size    int // Total size of value that should be scaled
	Scaling int // Defines how much the size is scaled. Range of values: [1, Size]
}

func (b Bias) Speed(x int) Bias {
	if b.Scaling*x < 0 {
		x = int(math.MaxInt64) / b.Scaling
	}
	if (b.Scaling*x)%b.Size == 0 {
		return Bias{
			Size:    b.Size,
			Scaling: b.Size,
		}
	} else {
		return Bias{
			Size:    b.Size,
			Scaling: (b.Scaling * x) % b.Size,
		}
	}

}

func biasedFactor(bias Bias, size int) int {
	switch {
	case bias.Size >= size+1:
		return (bias.Scaling * size) / bias.Size
	case bias.Size == 1:
		return 0
	case bias.Scaling == bias.Size:
		return size
	default:
		return (bias.Scaling - 1) * size / (bias.Size - 1)
	}
}
