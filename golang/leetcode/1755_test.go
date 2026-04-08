package main

import (
	"testing"
)

func TestMinAbsDifference(t *testing.T) {
	tests := []struct {
		nums []int
		goal int
		want int
	}{
		{[]int{5, -7, 3, 5}, 6, 0},
		{[]int{7, -9, 15, -2}, -5, 1},
		{[]int{1, 2, 3}, -7, 7},
	}
	for i, tt := range tests {
		if got := minAbsDifference(tt.nums, tt.goal); got != tt.want {
			t.Errorf("case %d: minAbsDifference(%v, %d) = %d; want %d", i, tt.nums, tt.goal, got, tt.want)
		}
	}
}
