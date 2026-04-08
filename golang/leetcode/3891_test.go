package main

import "testing"

func TestMinIncrease(t *testing.T) {
	tests := []struct {
		nums []int
		want int64
	}{
		{[]int{1, 2, 2}, 1},
		{[]int{2, 1, 1, 3}, 2},
		{[]int{5, 2, 1, 4, 3}, 4},
		{[]int{1, 1, 1, 1, 1, 1}, 2}, // Bonus: n=6 even case. Peaks {1, 3} or {1, 4} or {2, 4} all cost 2.
	}
	for i, tt := range tests {
		if got := minIncrease(tt.nums); got != tt.want {
			t.Errorf("case %d: minIncrease(%v) = %d; want %d", i, tt.nums, got, tt.want)
		}
	}
}
