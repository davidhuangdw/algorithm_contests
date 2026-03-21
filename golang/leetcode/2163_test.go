package main

import (
	"testing"
)

func TestMinimumDifference(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 1, 2},
			expected: -1,
		},
		{
			name:     "Example 2",
			nums:     []int{7, 9, 5, 8, 1, 3},
			expected: 1,
		},
		{
			name:     "Increasing numbers",
			nums:     []int{1, 2, 3, 4, 5, 6},
			expected: -8,
		},
		{
			name:     "Decreasing numbers",
			nums:     []int{6, 5, 4, 3, 2, 1},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDifference(tt.nums); got != tt.expected {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.expected)
			}
		})
	}
}
