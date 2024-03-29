package addbinary

import "testing"

func TestAddBinary(t *testing.T) {
	input1 := []string{
		"11",
		"1010",
	}
	input2 := []string{
		"1",
		"1011",
	}
	want := []string{
		"100",
		"10101",
	}

	for i := range input1 {
		got := addBinary(input1[i], input2[i])
		if got != want[i] {
			t.Fatalf(`addBinary function failed, want %s, got %s .`, want[i], got)
		}
	}
}
