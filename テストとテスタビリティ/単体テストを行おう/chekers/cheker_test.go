package checkers

import "testing"

var patterns = []struct {
	ptn    string
	input  int32
	expect bool
}{
	{"zero", 0, true},
	{"one", 1, false},
	{"two", 2, true},
	{"minus-one", -1, false},
	{"minus-two", -2, true},
	{"ten", 10, true},
}

func TestIsEven(t *testing.T) {
	for _, p := range patterns {
		t.Run(p.ptn, func(t *testing.T) {
			if r := isEven(p.input); r != p.expect {
				t.Errorf("test failed. expect=%t, but actual=%t", p.expect, r)
			}
		})
	}
}
