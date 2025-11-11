package main

import "testing"

func TestAdd(t *testing.T) {
	r := Plus(2, 2)
	exp := 4

	if r != exp {
		t.Errorf("Add(2, 3) = %d; expected %d", r, exp)
	}
}
