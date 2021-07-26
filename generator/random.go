package generator

// Random is an interface for random number generation
type Random interface {
	// Int64 generates random int64 in specifed range [min, max] (inclusive)
	Int64(min, max int64) int64

	// Uint64 generate random uint64 in specified range [min, max] (inclusive)
	Uint64(min, max uint64) uint64

	// Split returns new Random that can be used idenpendently of original. Random
	// returned by Split can have it's seed changed without affecting the original
	Split() Random

	// Seed seeds Random with seed value
	Seed(seed int64)
}
