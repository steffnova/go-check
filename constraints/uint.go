package constraints

import (
	"math"
	"strconv"
)

type Uint struct {
	Min uint
	Max uint
}

func UintDefault() Uint {
	if strconv.IntSize == 32 {
		return Uint{
			Min: 0,
			Max: math.MaxUint32,
		}
	}
	return Uint{
		Min: 0,
		Max: uint(math.MaxUint32)<<32 | uint(math.MaxUint32),
	}
}

type Uint8 struct {
	Min uint8
	Max uint8
}

func Uint8Default() Uint8 {
	return Uint8{
		Min: 0,
		Max: math.MaxUint8,
	}
}

type Uint16 struct {
	Min uint16
	Max uint16
}

func Uint16Default() Uint16 {
	return Uint16{
		Min: 0,
		Max: math.MaxUint16,
	}
}

type Uint32 struct {
	Min uint32
	Max uint32
}

func Uint32Default() Uint32 {
	return Uint32{
		Min: 0,
		Max: math.MaxUint32,
	}
}

type Uint64 struct {
	Min uint64
	Max uint64
}

func Uint64Default() Uint64 {
	return Uint64{
		Min: 0,
		Max: math.MaxUint64,
	}
}
