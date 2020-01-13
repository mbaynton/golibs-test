package adder

import "testing"

func TestAdder(t *testing.T) {
	want := 10
	got := Add(7, 3)
	if want != got {
		t.Errorf("Want 10, got %d", got)
	}
}
