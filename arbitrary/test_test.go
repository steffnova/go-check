package arbitrary

import (
	"testing"
)

func TestA(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, len(a))
	copy(b, a)
	b[0] = 3

	if a[0] == b[0] {
		t.Error(a[0])
	}
}
