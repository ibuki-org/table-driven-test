package hsd

import "testing"

func TestStringDistance(t *testing.T) {
	got := StringDistance("abc", "abd")
	want := 1
	if got != want {
		t.Fatalf("StringDistance(\"abc\", \"abd\") = %d; want %d", got, want)
	}
}
