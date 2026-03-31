package main

import (
	"testing"
)

func TestFindKthSmallest(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		k        int
		expected int64
	}{
		{
			name:     "Example 1",
			coins:    []int{3, 6, 9},
			k:        3,
			expected: 9,
		},
		{
			name:     "Example 2",
			coins:    []int{5, 2},
			k:        7,
			expected: 12,
		},
		{
			name:     "Single coin",
			coins:    []int{2},
			k:        5,
			expected: 10,
		},
		{
			name:     "Multiple coins with overlaps",
			coins:    []int{2, 3, 5},
			k:        10,
			expected: 14, // 2, 3, 4, 5, 6, 8, 9, 10, 12, 14
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthSmallest(tt.coins, tt.k); got != tt.expected {
				t.Errorf("findKthSmallest() = %v, want %v", got, tt.expected)
			}
		})
	}
}
