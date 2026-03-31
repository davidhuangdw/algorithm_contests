package main

import (
	"testing"
)

func TestMaxIncreasingCells(t *testing.T) {
	tests := []struct {
		name     string
		mat      [][]int
		expected int
	}{
		{
			name:     "Example 1",
			mat:      [][]int{{3, 1}, {3, 4}},
			expected: 2,
		},
		{
			name:     "Example 2",
			mat:      [][]int{{1, 1}, {1, 1}},
			expected: 1,
		},
		{
			name:     "Example 3",
			mat:      [][]int{{3, 1, 6}, {-9, 5, 7}},
			expected: 4,
		},
		{
			name:     "Single row",
			mat:      [][]int{{1, 2, 3, 4}},
			expected: 4,
		},
		{
			name:     "Single column",
			mat:      [][]int{{1}, {2}, {3}, {4}},
			expected: 4,
		},
		{
			name:     "Descending values",
			mat:      [][]int{{4, 3}, {2, 1}},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIncreasingCells(tt.mat); got != tt.expected {
				t.Errorf("maxIncreasingCells(%v) = %v, want %v", tt.name, got, tt.expected)
			}
			// maxIncreasingCells1 is also available, checking if it passes as well
			if got := maxIncreasingCells1(tt.mat); got != tt.expected {
				t.Errorf("maxIncreasingCells1(%v) = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}
