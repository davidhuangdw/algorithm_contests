package main

import (
	"testing"
)

func Test_minStable(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		maxC int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 5, 10},
			maxC: 1,
			want: 1,
		},
		{
			name: "Example 2",
			nums: []int{2, 6, 8},
			maxC: 2,
			want: 1,
		},
		{
			name: "Example 3",
			nums: []int{2, 4, 9, 6},
			maxC: 1,
			want: 2,
		},
		{
			name: "No modifications needed",
			nums: []int{1, 1, 1},
			maxC: 1,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStable(tt.nums, tt.maxC); got != tt.want {
				t.Errorf("minStable() = %v, want %v", got, tt.want)
			}
		})
	}
}
