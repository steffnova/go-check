package constraints

// Bias contains Size and Scaling properties which are used to scale constraint limits.
// Scaling factor is calculated by diving Size with Scaling
type Bias struct {
	Size    int // Total size of value that should be scaled
	Scaling int // Defines how much the size is scaled. Range of values: [1, Size]
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
