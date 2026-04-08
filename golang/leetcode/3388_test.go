package main

import (
	"testing"
)

func TestBeautifulSplits(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 2, 1}, 2},
		{[]int{1, 2, 3, 4}, 0},
	}
	for _, tt := range tests {
		if got := beautifulSplits(tt.nums); got != tt.want {
			t.Errorf("beautifulSplits(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
