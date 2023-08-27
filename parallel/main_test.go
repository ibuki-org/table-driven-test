package parallel

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		lhs  int
		rhs  int
		want int
	}{
		{name: "1+2", lhs: 1, rhs: 2, want: 3},
		{name: "2+3", lhs: 2, rhs: 3, want: 5},
		{name: "3+4", lhs: 3, rhs: 4, want: 7},
	}

	for _, test := range tests {
		test := test // ここをしないとタイミングによって、最後の項目のみ参照されてしまう。
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Add(test.lhs, test.rhs)
			if got != test.want {
				t.Errorf("Add(%d, %d) = %d, want %d", test.lhs, test.rhs, got, test.want)
			}
		})
	}
}
