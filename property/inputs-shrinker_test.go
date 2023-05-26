package property

import "testing"

func TestInputsShrinker(t *testing.T) {

	testCases := map[string]func(*testing.T){}

	for name, testCase := range testCases {
		t.Run(name, testCase)
	}

}
