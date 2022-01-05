package main

import "testing"

func TestDiff(t *testing.T) {

	result := diff(3, 2)

	if result != 1 {
		t.Error("The result must be 1")
	}
}
