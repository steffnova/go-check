package check

import (
	"math/rand"
	"testing"
	"time"
)

type Config struct {
	Seed       int64
	Iterations int64
}

func DefaultConfig() Config {
	return Config{
		Seed:       time.Now().UnixNano(),
		Iterations: 100,
	}
}

func Check(t *testing.T, property property, config ...Config) {
	configuration := DefaultConfig()
	if len(config) > 0 {
		configuration = config[0]
	}

	random := rng{
		Rand: rand.New(rand.NewSource(configuration.Seed)),
	}

	run, err := property(random)
	if err != nil {
		t.Fatalf("failed to run property. %s", err)
	}

	for i := int64(0); i < configuration.Iterations; i++ {
		if err := run(); err != nil {
			t.Fatalf("\nCheck failed with seed: %d. \n%s", configuration.Seed, err)
		}
	}
}
