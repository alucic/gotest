// +build integration

package test

import "testing"

func TestIsNotEven(t *testing.T) {
	got := IsEven(3)
	if got != false {
		t.Fatalf("Failed")
	}
}
