package main

import "testing"

func TestHexString(t *testing.T) {
	h := Hex(10)
	expect := "a"
	actual := h.string()
	if expect != actual {
		t.Errorf("expect=%s actual=%s", expect, actual)
	}
}
