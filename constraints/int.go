package constraints

import (
	"math"
	"strconv"
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
// For 32bit architecture Int{Min: math.MinInt32, Max: math.MaxInt32}
//
// For 64bit architecture Int{Min: math.MinInt64, Max: math.MaxInt64}
func IntDefault() Int {
	if strconv.IntSize == 32 {
		return Int{
			Min: math.MinInt32,
			Max: math.MaxInt32,
		}
	}
	return Int{
		// Using bitwise operations to calculate maxInt64 and minInt64 to avoid
		// compiler errors on 32 bit platforms, as int64 cannot be assigned to int
		Min: math.MinInt32 << 1,
		Max: (int(math.MaxInt32) << 32) | ((int(math.MaxInt32 + 1)) | int(math.MaxInt32)),
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
		Min: math.MinInt16,
		Max: math.MaxInt16,
	}
}
