package check

import (
	"math"
	"math/big"
	"math/rand"

	"github.com/steffnova/go-check/generator"
)

type rng struct {
	Rand *rand.Rand
}

func (r rng) Int64(minInt64, maxInt64 int64) int64 {
	max := big.NewInt(maxInt64)
	min := big.NewInt(minInt64)

	val := big.NewInt(0).Sub(max, min)
	val = big.NewInt(0).Add(val, big.NewInt(1))
	val = val.Rand(r.Rand, val)
	val = val.Add(val, min)

	return val.Int64()
}

func (r rng) Uint64(minUint64, maxUint64 uint64) uint64 {
	max := big.NewInt(math.MaxInt64)
	max = max.Mul(max, big.NewInt(int64(maxUint64/uint64(math.MaxInt64))))
	max = max.Add(max, big.NewInt(int64(maxUint64%uint64(math.MaxInt64))))

	min := big.NewInt(math.MaxInt64)
	min = min.Mul(min, big.NewInt(int64(minUint64/uint64(math.MaxInt64))))
	min = min.Add(min, big.NewInt(int64(minUint64%uint64(math.MaxInt64))))

	diff := big.NewInt(0).Sub(max, min)
	diff = diff.Add(diff, big.NewInt(1))

	n := diff.Rand(r.Rand, diff)
	n = n.Add(diff, min)

	return n.Uint64()
}

func (r rng) Float64(minFloat64, maxFloat64 float64) float64 {
	deviation := maxFloat64/2 - minFloat64/2
	mean := deviation + minFloat64

	for {
		random := r.Rand.NormFloat64()*deviation + mean
		if minFloat64 <= random && random <= maxFloat64 {
			return random
		}
	}
}

func (r rng) Seed(seed int64) {
	r.Rand.Seed(seed)
}

func (r rng) Split() generator.Random {
	newSeed := r.Int64(math.MinInt64, math.MaxInt64)
	return &rng{
		Rand: rand.New(rand.NewSource(newSeed)),
	}
}
