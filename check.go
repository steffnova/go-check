package check

import (
	"flag"
	"fmt"
	"hash/maphash"
	"math/rand"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/property"
)

var (
	seedFlag       = flag.CommandLine.Int64("seed", int64(new(maphash.Hash).Sum64()), "seed value used for generating property inputs")
	iterationsFlag = flag.CommandLine.Int64("iterations", 100, "number of iterations run for the property")
)

// Config is configuration used by Check.
type Config struct {
	Seed       int64 // Seed used by random number generator
	Iterations int64 // Number of times property will be checked
}

// Check checks if property holds. First parameter is *testing.T that will report
// error if property returns an error (property doesn't hold). The property paramter
// should be defined using [property.Define]. The config parameter, even though it is
// a variadice parameter it uses only the first instance of [Config] passed to it. If
// config is not specified default configuration is used (random seed and 100 iterations).
// Config will run property number of times equal to Iterations values, specified by config
// parameter or (100 if default value is used).
// Following example demonstrates how to use Check in tests:
//
//	package main_test
//
//	import (
//	    "fmt"
//	    "testing"
//
//	    "github.com/steffnova/go-check"
//	    "github.com/steffnova/go-check/generator"
//	    "github.com/steffnova/go-check/property"
//	)
//
//	func TestSubtractionCommutativity(t *testing.T) {
//	    check.Check(t, property.Define(
//	        property.Inputs(
//	            generator.Int()
//	            generator.Int()
//	        )
//	        property.Predicate(func(x, y int) error {
//	            if x-y != y-x {
//	                return fmt.Errorf("commutativity does not hold for subtraction.")
//	            }
//	            return nil
//	        })
//	    ))
//	}
//
// Commutativity does not hold for subtraction operation for every x and y as it is defined in predicate.
// Because of this an error will be thrown that will report value for which property failed. If shrinking
// is possible, [property.Define] will report smallest possible value for which predicate didn't hold.
//
//	--- FAIL: TestSubtractionCommutativity (0.00s)
//	    Check failed after 1 tests with seed: 1646421732271105000.
//	    Property failed for inputs: [
//	        <int> 0,
//	        <int> 1
//	    ]
//	    Shrunk 123 time(s)
//	    Failure reason: commutativity does not hold for subtraction.
//
//	    Re-run:
//	    go test -run=TestSubtractionCommutativity -seed=1646421732271105000 -iterations=100
func Check(t *testing.T, property property.Property, config ...Config) {
	t.Helper()
	if property == nil {
		t.Fatalf("property can't be nil")
	}
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

	for i := int64(0); i < configuration.Iterations; i++ {
		bias := constraints.Bias{
			Size:    int(configuration.Iterations),
			Scaling: int(configuration.Iterations) - int(i),
		}

		details, err := property(random, bias)
		if err != nil {
			t.Fatal(err)
		}

		if details.FailureReason != nil {
			t.Fatal(
				fmt.Sprintf("\nCheck failed after %d test(s) with seed: %d.", i, configuration.Seed),
				fmt.Sprintf("\n%s", (propertyFailed(details.FailureInput.Values()))),
				fmt.Sprintf("\nShrunk %d time(s)", details.NumberOfShrinks),
				fmt.Sprintf("\nFailure reason: %s", details.FailureReason),
				fmt.Sprintf("\n\nRe-run:\ngo test -run=%s -seed=%d -iterations=%d", t.Name(), configuration.Seed, configuration.Iterations),
			)
		}
	}
}
