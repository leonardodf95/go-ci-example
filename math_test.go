package main

import "testing"

func TestSum(t *testing.T) {
	x := 1
	y := 2
	if sum(x, y) != 3 {
		t.Errorf("Expected %d, got %d", 3, sum(x, y))
	}
}
