package generator

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"reflect"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

// Uint64 returns Arbitrary generator that can be used to create Uint64 generator. Range in which
// uint64 value is generated is defined by constraints [min, max] (min and max are included in range).
//
// If no constraints are provided default range for int64 is used [math.MinUint64, math.MaxUint64]. Even
// though limits is a variadic argument only the first value is used for defining constraints
func Uint64(limits ...constraints.Uint64) Arbitrary {
	constraint := constraints.Uint64Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return func(target reflect.Type) (Type, error) {
		if target.Kind() != reflect.Uint64 {
			return Type{}, fmt.Errorf("target arbitrary's kind must be Uint64. Got: %s", target.Kind())
		}
		return Type{
			Type: reflect.TypeOf(uint64(0)),
			Generate: func(rand *rand.Rand) arbitrary.Type {
				max := big.NewInt(math.MaxInt64)
				max = max.Mul(max, big.NewInt(int64(constraint.Max/uint64(math.MaxInt64))))
				max = max.Add(max, big.NewInt(int64(constraint.Max%uint64(math.MaxInt64))))

				min := big.NewInt(math.MaxInt64)
				min = min.Mul(min, big.NewInt(int64(constraint.Min/uint64(math.MaxInt64))))
				min = min.Add(min, big.NewInt(int64(constraint.Min%uint64(math.MaxInt64))))

				diff := big.NewInt(0).Sub(max, min)
				diff = diff.Add(diff, big.NewInt(1))

				n := diff.Rand(rand, diff)
				n = n.Add(diff, min)

				return arbitrary.Uint64{
					Constraint: constraint,
					N:          n.Uint64(),
				}
			},
		}, nil
	}
}

// Uint32 returns Arbitrary generator that can be used to create Uint32 generator. Range in which
// uint32 value is generated is defined by constraints [min, max] (min and max are included in range).
//
// If no constraints are provided default range for uint32 is used [math.MinUint32, math.MaxUint32]. Even
// though limits is a variadic argument only the first value is used for defining constraints
func Uint32(limits ...constraints.Uint32) Arbitrary {
	constraint := constraints.Uint32Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return Uint64(constraints.Uint64{
		Max: uint64(constraint.Max),
		Min: uint64(constraint.Min),
	}).Map(func(n uint64) uint32 {
		return uint32(n)
	})
}

// Uint16 returns Arbitrary generator that can be used to create Uint16 generator. Range in which
// uint16 value is generated is defined by constraints [min, max] (min and max are included in range).
//
// If no constraints are provided default range for uint16 is used [math.MinUint16, math.MaxUint16]. Even
// though limits is a variadic argument only the first value is used for defining constraints
func Uint16(limits ...constraints.Uint16) Arbitrary {
	constraint := constraints.Uint16Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return Uint64(constraints.Uint64{
		Max: uint64(constraint.Max),
		Min: uint64(constraint.Min),
	}).Map(func(n uint64) uint16 {
		return uint16(n)
	})
}

// Uint8 returns Arbitrary generator that can be used to create Uint8 generator. Range in which
// uint8 value is generated is defined by constraints [min, max] (min and max are included in range).
//
// If no constraints are provided default range for int8 is used [math.MinUint8, math.MaxUint8]. Even
// though limits is a variadic argument only the first value is used for defining constraints
func Uint8(limits ...constraints.Uint8) Arbitrary {
	constraint := constraints.Uint8Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return Uint64(constraints.Uint64{
		Max: uint64(constraint.Max),
		Min: uint64(constraint.Min),
	}).Map(func(n uint64) uint8 {
		return uint8(n)
	})
}

// Int returns Arbitrary generator that can be used to create Int generator. Range in which
// int value is generated is defined by constraints [min, max] (min and max are included in range).
//
// If no constraints are provided default range for uint is used [math.MinUint32, math.MaxUint32] for
// 32bit architecture or [math.MinUint64, math.MaxUint64] for 64bit architecture. Even though limits
// is a variadic argument only the first value is used for defining constraints
func Uint(limits ...constraints.Uint) Arbitrary {
	constraint := constraints.UintDefault()
	if len(limits) > 0 {
		constraint = limits[0]
	}

	return Uint64(constraints.Uint64{
		Max: uint64(constraint.Max),
		Min: uint64(constraint.Min),
	}).Map(func(n uint64) uint {
		return uint(n)
	})
}
