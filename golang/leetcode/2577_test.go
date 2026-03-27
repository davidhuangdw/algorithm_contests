package main

import (
	"testing"
)

func TestMinimumTime(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{0, 1, 3, 2},
				{5, 1, 2, 5},
				{4, 3, 8, 6},
			},
			expected: 7,
		},
		{
			name: "Example 2",
			grid: [][]int{
				{0, 2, 4},
				{3, 2, 1},
				{1, 0, 4},
			},
			expected: -1,
		},
		{
			name: "Small grid reachable",
			grid: [][]int{
				{0, 1},
				{1, 1},
			},
			expected: 2,
		},
		{
			name: "Need to wait",
			grid: [][]int{
				{0, 1, 10},
				{1, 1, 10},
				{1, 1, 1},
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTime(tt.grid); got != tt.expected {
				t.Errorf("minimumTime() = %v, want %v", got, tt.expected)
			}
		})
	}
}
