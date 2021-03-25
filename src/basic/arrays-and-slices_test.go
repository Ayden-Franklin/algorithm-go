package main
import "testing"

func TestSumArray(t *testing.T) {

	numbers := [5]int{1,2,3,4,5}

	got := SumArray(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
func TestSumSlice(t *testing.T) {

	numbers := []int{1,2,3,4,5}

	length, got := SumSlice(numbers)
	want := 15

	if length != 5 || got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}