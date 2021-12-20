package constraints

import (
	"math"
	"math/bits"
)

// Int constraints
type Int struct {
	Min int // Min int value
	Max int // Max int value
}

// IntDefault returns default int constraints.
// Underlying architecture defines whether int is int32 or int64
// and because of that default min and max constraint values are
// different.
//
// Spec definition: https://golang.org/ref/spec#Numeric_types
//
// - For 32bit architecture Int{Min: math.MinInt32, Max: math.MaxInt32}
//
// - For 64bit architecture Int{Min: math.MinInt64, Max: math.MaxInt64}
func IntDefault() Int {
	if bits.UintSize == 32 {
		return Int{
			Min: math.MinInt32,
			Max: math.MaxInt32,
		}
	}
	return Int{
		Min: int(math.MinInt64),
		Max: int(math.MaxInt64),
	}
}

type Int8 struct {
	Min int8
	Max int8
}

func Int8Default() Int8 {
	return Int8{
		Min: math.MinInt8,
		Max: math.MaxInt8,
	}
}

type Int16 struct {
	Min int16
	Max int16
}

func Int16Default() Int16 {
	return Int16{
		Min: math.MinInt16,
		Max: math.MaxInt16,
	}
}

type Int32 struct {
	Min int32
	Max int32
}

func Int32Default() Int32 {
	return Int32{
		Min: math.MinInt32,
		Max: math.MaxInt32,
	}
}

type Int64 struct {
	Min int64
	Max int64
}

func Int64Default() Int64 {
	return Int64{
		Min: math.MinInt64,
		Max: math.MaxInt64,
	}
}

func (i Int64) Biased(bias Bias) Int64 {
	switch {
	case i.Min <= 0 && i.Max <= 0:
		ui := Uint64{Min: uint64(-i.Max), Max: uint64(-i.Min)}.Baised(bias)
		return Int64{
			Min: int64(-ui.Max),
			Max: int64(-ui.Min),
		}
	case i.Min >= 0 && i.Max >= 0:
		ui := Uint64{Min: uint64(i.Min), Max: uint64(i.Max)}.Baised(bias)
		return Int64{
			Min: int64(ui.Min),
			Max: int64(ui.Max),
		}
	default:
		// In order to perserve symmetry ratio for generated numbers (50:50 ration between
		// negative and positive numbers when range is [-n, n]), 0 is considered a positive
		// number and negative numbers start from -1. This way negative numbers can be
		// represented with same number of bits as positive, and will ensure that range of
		// positive and negative values is expanded at the same rate.
		ui := Uint64{Min: 0, Max: uint64(-i.Min) - 1}.Baised(bias)
		return Int64{
			Min: int64(-ui.Max) - 1,
			Max: Int64{Min: 0, Max: i.Max}.Biased(bias).Max,
		}
	}

}
