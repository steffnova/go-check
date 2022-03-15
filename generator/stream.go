package generator

import (
	"fmt"
	"math/rand"
)

// Stream streams data using a streamer. Seed for random number generation is specified by
// seed paramere, Number of values that will be generated is specified by "count" parameter
// Error is returned if streamer returns an error.
func Stream(seed, count uint64, streamer streamer) error {
	random := RandomNumber{
		Rand: rand.New(rand.NewSource(int64(seed))),
	}

	for i := uint64(0); i < count; i++ {
		if err := streamer(random); err != nil {
			return fmt.Errorf("failed to run stream: %s", err)
		}
	}

	return nil
}
