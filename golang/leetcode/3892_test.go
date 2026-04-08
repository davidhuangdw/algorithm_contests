package main

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 2}, 1, 1},
		{[]int{1, 2, 2}, 2, -1}, // Cannot have 2 peaks in length 3.
		{[]int{1, 1, 1, 1}, 2, 2}, // cost [1,1,1,1]. solve([1,1,1], 2) = 2.
		{[]int{1, 1, 1, 1}, 3, -1},
	}
	for i, tt := range tests {
		if got := minOperations(tt.nums, tt.k); got != tt.want {
			t.Errorf("case %d: minOperations(%v, %d) = %d; want %d", i, tt.nums, tt.k, got, tt.want)
		}
	}
}
