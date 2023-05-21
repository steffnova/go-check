package check

import (
	"flag"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/steffnova/go-check/arbitrary"
)

var (
	seedFlag       = flag.CommandLine.Int64("seed", time.Now().UnixNano(), "seed value used for generating property inputs")
	iterationsFlag = flag.CommandLine.Int64("iterations", 100, "number of iterations run for the property")
)

// Config is configuration used by Check.
type Config struct {
	Seed       int64 // Seed used by random number generator
	Iterations int64 // Number of times property will be checked
}

// Check checks if property holds. First parameter is *testing.T that will report
// error if property doesn't hold. Even though config is a variadice parameter it
// uses only the first value. If config is not specified default configuration is
// used (random seed and 100 iterations).
func Check(t *testing.T, property property, config ...Config) {
	t.Helper()
	configuration := Config{
		Seed:       *seedFlag,
		Iterations: *iterationsFlag,
	}

	if len(config) > 0 {
		configuration = config[0]
	}

	random := arbitrary.RandomNumber{
		Rand: rand.New(rand.NewSource(configuration.Seed)),
	}

	details, err := property(random, int(configuration.Iterations))
	if err != nil {
		t.Fatal(err)
	}
	if details.FailureReason != nil {
		t.Fatal(
			fmt.Sprintf("\nCheck failed after %d test(s) with seed: %d.", details.NumberOfIterations, configuration.Seed),
			fmt.Sprintf("\n%s", (propertyFailed(details.FailureInput.Values()))),
			fmt.Sprintf("\nShrunk %d time(s)", details.NumberOfShrinks),
			fmt.Sprintf("\nFailure reason: %s", details.FailureReason),
			fmt.Sprintf("\n\nRe-run:\ngo test -run=%s -seed=%d -iterations=%d", t.Name(), configuration.Seed, configuration.Iterations),
		)
	}
}
