package main

import "testing"

func TestTimes(t *testing.T) {

	result := times(2, 3)

	if result != 6 {
		t.Error("The result must be 6")
	}
}
