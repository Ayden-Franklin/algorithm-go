package main

import "testing"

func TestAdd(t *testing.T) {
	var i, k int = 2, 3
	result := Add(i,k)
	expected := 5
	if result != expected {
		t.Errorf("Failt to calculate integers %v + %v, want: %v", i, k, expected)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}