package main

import (
	"testing"
)

func TestClosestToTarget(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{
			name:   "Example 1",
			arr:    []int{9, 12, 3, 7, 15},
			target: 5,
			want:   2,
		},
		{
			name:   "Example 2",
			arr:    []int{1000000, 1000000, 1000000},
			target: 1,
			want:   999999,
		},
		{
			name:   "Example 3",
			arr:    []int{1, 2, 4, 8, 16},
			target: 0,
			want:   0,
		},
		{
			name:   "Single element",
			arr:    []int{10},
			target: 5,
			want:   5,
		},
		{
			name:   "Multiple elements same result",
			arr:    []int{7, 7, 7},
			target: 7,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestToTarget(tt.arr, tt.target); got != tt.want {
				t.Errorf("closestToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
