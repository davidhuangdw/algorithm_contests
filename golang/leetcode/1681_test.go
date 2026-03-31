package main

import (
	"testing"
)

func TestMinimumIncompatibility(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 1, 4},
			k:        2,
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{6, 3, 8, 1, 3, 1, 2, 2},
			k:        4,
			expected: 6,
		},
		{
			name:     "Example 3",
			nums:     []int{5, 3, 3, 6, 3, 3},
			k:        3,
			expected: -1,
		},
		{
			name:     "Single element subsets",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 0,
		},
		{
			name:     "Impossible due to too many duplicates",
			nums:     []int{1, 1, 1, 2},
			k:        2,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumIncompatibility(tt.nums, tt.k); got != tt.expected {
				t.Errorf("minimumIncompatibility(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}
