package check

import (
	"flag"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var (
	seedFlag       = flag.CommandLine.Int64("seed", time.Now().UnixNano(), "seed value used for generating property inputs")
	iterationsFlag = flag.CommandLine.Int64("iterations", 100, "number of iterations run for the property")
)

type Config struct {
	Seed       int64
	Iterations int64
}

func DefaultConfig() Config {
	config := Config{
		Seed:       *seedFlag,
		Iterations: *iterationsFlag,
	}

	return config
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
			t.Fatal(
				fmt.Sprintf("\nCheck failed after %d tests with seed: %d. \n%s", i+1, configuration.Seed, err),
				fmt.Sprintf("\n\nRe-run:\ngo test -run=%s -seed=%d -iterations=%d", t.Name(), configuration.Seed, configuration.Iterations),
			)
		}
	}
}
