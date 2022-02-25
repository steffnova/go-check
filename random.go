package check

import (
	"math"
	"math/big"
	"math/rand"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

type Rng struct {
	Rand *rand.Rand
}

func (r Rng) Int64(limit constraints.Int64) int64 {
	max := big.NewInt(limit.Max)
	min := big.NewInt(limit.Min)

	val := big.NewInt(0).Sub(max, min)
	val = big.NewInt(0).Add(val, big.NewInt(1))
	val = val.Rand(r.Rand, val)
	val = val.Add(val, min)

	return val.Int64()
}

func (r Rng) Uint64(limit constraints.Uint64) uint64 {
	max := big.NewInt(math.MaxInt64)
	max = max.Mul(max, big.NewInt(int64(limit.Max/uint64(math.MaxInt64))))
	max = max.Add(max, big.NewInt(int64(limit.Max%uint64(math.MaxInt64))))

	min := big.NewInt(math.MaxInt64)
	min = min.Mul(min, big.NewInt(int64(limit.Min/uint64(math.MaxInt64))))
	min = min.Add(min, big.NewInt(int64(limit.Min%uint64(math.MaxInt64))))

	diff := big.NewInt(0).Sub(max, min)
	diff = diff.Add(diff, big.NewInt(1))

	n := diff.Rand(r.Rand, diff)
	n = n.Add(diff, min)

	return n.Uint64()
}

func (r Rng) Float64(limit constraints.Float64) float64 {
	deviation := limit.Max/2 - limit.Min/2
	mean := deviation + limit.Min

	for {
		random := r.Rand.NormFloat64()*deviation + mean
		if limit.Min <= random && random <= limit.Max {
			return random
		}
	}
}

func (r Rng) Seed(seed int64) {
	r.Rand.Seed(seed)
}

func (r Rng) Split() generator.Random {
	newSeed := r.Int64(constraints.Int64{
		Min: math.MinInt64,
		Max: math.MaxInt64,
	})
	return &Rng{
		Rand: rand.New(rand.NewSource(newSeed)),
	}
}
