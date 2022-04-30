package sum

import "testing"

func TestSum(t *testing.T) {
	xs := []int{1, -2, 3, 0}
	got := Sum(xs...)
	want := 2

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, xs)
	}
}
