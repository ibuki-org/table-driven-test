package hsd

import (
	"reflect"
	"testing"
)

func TestStringDistance(t *testing.T) {
	tests := []struct {
		name string
		lhs  string
		rhs  string
		want int
	}{
		{name: "abc/abc", lhs: "abc", rhs: "abc", want: 0},
		{name: "abc/abd", lhs: "abc", rhs: "abd", want: 1},
		{name: "abc/abcd", lhs: "abc", rhs: "abcd", want: -1},
		{name: "abc/ab", lhs: "abc", rhs: "ab", want: -1},
		{name: "abc/xyz", lhs: "abc", rhs: "xyz", want: 3},
		{name: "abc/xyz", lhs: "abc", rhs: "xyz", want: 3},
		{name: "abc/xyz", lhs: "abc", rhs: "xyz", want: 3},
		{name: "abc/xyz", lhs: "abc", rhs: "xyz", want: 3},
		{name: "abc/xyz", lhs: "abc", rhs: "xyz", want: 3},
		{name: "abc/xyz", lhs: "abc", rhs: "xyz", want: 3},
	}

	for _, tc := range tests {
		got := StringDistance(tc.lhs, tc.rhs)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("StringDistance(%q, %q) = %d; want %d", tc.lhs, tc.rhs, got, tc.want)
		}
	}
}
