package generator

// import (
// 	"math"
// 	"math/big"
// 	"math/rand"

// 	"github.com/steffnova/go-check/constraints"
// )

// // arbitrary.Random is an interface for random number generation
// type Random interface {
// 	// Uint64 generates random uint64 in specified range [min, max] (inclusive)
// 	Uint64(constraints.Uint64) uint64

// 	// Split returns new arbitrary.Random that can be used idenpendently of original. arbitrary.Random
// 	// returned by Split can have it's seed changed without affecting the original
// 	Split() arbitrary.Random

// 	// Seed seeds arbitrary.Random with seed value
// 	Seed(seed int64)
// }

// // arbitrary.RandomNumber is implementation of arbitrary.Random interface
// type arbitrary.RandomNumber struct {
// 	Rand *rand.Rand
// }

// // Uint64 is implementation of arbitrary.Random.Uint64
// func (r arbitrary.RandomNumber) Uint64(limit constraints.Uint64) uint64 {
// 	max := big.NewInt(math.MaxInt64)
// 	max = max.Mul(max, big.NewInt(int64(limit.Max/uint64(math.MaxInt64))))
// 	max = max.Add(max, big.NewInt(int64(limit.Max%uint64(math.MaxInt64))))

// 	min := big.NewInt(math.MaxInt64)
// 	min = min.Mul(min, big.NewInt(int64(limit.Min/uint64(math.MaxInt64))))
// 	min = min.Add(min, big.NewInt(int64(limit.Min%uint64(math.MaxInt64))))

// 	diff := big.NewInt(0).Sub(max, min)
// 	diff = diff.Add(diff, big.NewInt(1))

// 	n := diff.Rand(r.Rand, diff)
// 	n = n.Add(diff, min)

// 	return n.Uint64()
// }

// // Seed is implementation of arbitrary.Random.Seed
// func (r arbitrary.RandomNumber) Seed(seed int64) {
// 	r.Rand.Seed(seed)
// }

// // Split is implementation of arbitrary.Random.Split
// func (r arbitrary.RandomNumber) Split() arbitrary.Random {
// 	newSeed := int64(r.Uint64(constraints.Uint64Default()))
// 	return &arbitrary.RandomNumber{
// 		Rand: rand.New(rand.NewSource(newSeed)),
// 	}
// }
