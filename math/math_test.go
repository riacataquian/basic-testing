package math

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		desc string
		in   []int
		want int
	}{
		{
			desc: "returns the sum of ints",
			in:   []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			desc: "returns the sum of negative ints",
			in:   []int{1, 2, 3, 4, -5},
			want: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got := Sum(test.in...)

			if got != test.want {
				t.Errorf("Sum(%v) = %d, want %d", test.in, got, test.want)
			}
		})
	}
}
