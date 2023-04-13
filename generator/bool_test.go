package generator

import (
	"errors"
	"testing"
)

func TestBool(t *testing.T) {
	err := Stream(0, 1, Streamer(
		func(n int) {},
		Bool(),
	))

	if !errors.Is(err, ErrorInvalidTarget) {
		t.Fatalf("Expected error: '%s'", ErrorInvalidTarget)
	}
}
