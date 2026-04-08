package main

import "testing"

func TestGoodSubtreeSum(t *testing.T) {
	tests := []struct {
		vals []int
		par  []int
		want int
	}{
		{
			vals: []int{1, 2, 3},
			par:  []int{-1, 0, 0},
			want: 11,
		},
		{
			vals: []int{12, 3},
			par:  []int{-1, 0},
			want: 18,
		},
		{
			vals: []int{11, 2},
			par:  []int{-1, 0},
			want: 4,
		},
		{
			vals: []int{135867, 418763},
			par:  []int{-1, 0},
			want: 837526,
		},
	}
	for i, tt := range tests {
		if got := goodSubtreeSum(tt.vals, tt.par); got != tt.want {
			t.Errorf("case %d: got %d, want %d", i, got, tt.want)
		}
	}
}
