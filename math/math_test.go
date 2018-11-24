package math

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		desc string
		in   []int
		want int
	}{
		{
			desc: "returns the sum of a single int",
			in:   []int{20},
			want: 20,
		},
		{
			desc: "returns the sum of the int slices",
			in:   []int{1, 2, 3, 4, 5},
			want: 15,
		},
	}

	for _, test := range tests {
		got := Sum(test.in...)

		if got != test.want {
			t.Errorf("Sum(%v) = %d, want %d", test.in, got, test.want)
		}
	}
}
