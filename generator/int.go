package generator

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
)

func Int64(limits ...constraints.Int64) Generator {
	constraint := constraints.Int64Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}

	switch {
	case constraint.Min > constraint.Max:
		return InvalidGen(fmt.Errorf("lower limit: %d cannot be greater than upper limit: %d", constraint.Min, constraint.Max))
	case constraint.Max < 0:
		return Uint64(constraints.Uint64{Min: uint64(-constraint.Max), Max: uint64(-constraint.Min)}).
			Map(func(x uint64) int64 {
				return int64(-x)
			})
	case constraint.Min >= 0:
		return Uint64(constraints.Uint64{Min: uint64(constraint.Min), Max: uint64(constraint.Max)}).
			Map(func(x uint64) int64 {
				return int64(x)
			})
	default:
		return OneFromWeighted(
			Weighted{
				Weight: uint(-(constraint.Min)),
				Gen: Uint64(constraints.Uint64{Min: 0, Max: uint64(-constraint.Min)}).
					Map(func(x uint64) int64 {
						return int64(-x)
					}),
			},
			Weighted{
				Weight: uint(constraint.Max) + 1,
				Gen: Uint64(constraints.Uint64{Min: 0, Max: uint64(constraint.Max)}).
					Map(func(x uint64) int64 {
						return int64(x)
					}),
			},
		)
	}
}

func Int32(limits ...constraints.Int32) Generator {
	constraint := constraints.Int32Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return Int64(constraints.Int64{Min: int64(constraint.Min), Max: int64(constraint.Max)}).
		Map(func(x int64) int32 {
			return int32(x)
		})
}

func Int16(limits ...constraints.Int16) Generator {
	constraint := constraints.Int16Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return Int64(constraints.Int64{Min: int64(constraint.Min), Max: int64(constraint.Max)}).
		Map(func(x int64) int16 {
			return int16(x)
		})
}

func Int8(limits ...constraints.Int8) Generator {
	constraint := constraints.Int8Default()
	if len(limits) > 0 {
		constraint = limits[0]
	}
	return Int64(constraints.Int64{Min: int64(constraint.Min), Max: int64(constraint.Max)}).
		Map(func(x int64) int8 {
			return int8(x)
		})
}

func Int(limits ...constraints.Int) Generator {
	constraint := constraints.IntDefault()
	if len(limits) > 0 {
		constraint.Min, constraint.Max = limits[0].Min, limits[0].Max
	}

	return Int64(constraints.Int64{
		Max: int64(constraint.Max),
		Min: int64(constraint.Min),
	}).Map(func(n int64) int {
		return int(n)
	})
}
