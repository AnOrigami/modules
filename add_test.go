package modules

import "testing"

func TestIntAdd(t *testing.T) {
	c := IntAdd(1, 2)
	if c != 3 {
		t.Error("add failed")
	}
}
