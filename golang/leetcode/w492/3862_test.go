package main

import "testing"

func Test_smallestBalancedIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "sum equals product",
			nums: []int{2, 1, 1, 2},
			want: 1,
		},
		{
			name: "no balance found",
			nums: []int{3, 2, 1, 6},
			want: -1,
		},
		{
			name: "no balance",
			nums: []int{1, 2, 3},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestBalancedIndex(tt.nums)
			if got != tt.want {
				t.Errorf("smallestBalancedIndex(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
