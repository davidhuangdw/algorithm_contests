package main

import (
	"testing"
)

func TestMaximumLength(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 1, 1, 3}, 2, 4},
		{[]int{2, 2, 3, 3}, 0, 2},
	}

	for _, tt := range tests {
		got := maximumLength(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("maximumLength(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
		}
	}
}
