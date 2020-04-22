package test

import "testing"

func TestIsEven(t *testing.T) {
	got := IsEven(2)
	if got != true {
		t.Fatalf("Failed")
	}
}
