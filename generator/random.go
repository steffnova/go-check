package generator

import (
	"math"
	"math/big"
	"math/rand"

	"github.com/steffnova/go-check/constraints"
)

// Random is an interface for random number generation
type Random interface {
	// Uint64 generates random uint64 in specified range [min, max] (inclusive)
	Uint64(constraints.Uint64) uint64

	// Split returns new Random that can be used idenpendently of original. Random
	// returned by Split can have it's seed changed without affecting the original
	Split() Random

	// Seed seeds Random with seed value
	Seed(seed int64)
}

// RandomNumber is implementation of Random interface
type RandomNumber struct {
	Rand *rand.Rand
}

// Uint64 is implementation of Random.Uint64
func (r RandomNumber) Uint64(limit constraints.Uint64) uint64 {
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

// Seed is implementation of Random.Seed
func (r RandomNumber) Seed(seed int64) {
	r.Rand.Seed(seed)
}

// Split is implementation of Random.Split
func (r RandomNumber) Split() Random {
	newSeed := int64(r.Uint64(constraints.Uint64Default()))
	return &RandomNumber{
		Rand: rand.New(rand.NewSource(newSeed)),
	}
}
