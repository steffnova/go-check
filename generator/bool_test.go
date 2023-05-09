package generator

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
)

func TestBool(t *testing.T) {
	err := Stream(0, 1, Streamer(
		func(n int) {},
		Bool(),
	))

	if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
		t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidTarget)
	}
}
