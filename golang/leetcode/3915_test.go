package main

import (
	"testing"
)

func TestMaxAlternatingSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int64
	}{
		{
			name:     "Example 1",
			nums:     []int{5, 4, 2},
			k:        2,
			expected: 7, // indices [0, 2] -> values [5, 2] sum 7
		},
		{
			name:     "Example 2",
			nums:     []int{3, 5, 4, 2, 4},
			k:        1,
			expected: 14,
		},
		{
			name:     "Custom 1",
			nums:     []int{1, 2, 3, 4, 5},
			k:        1,
			expected: 9, // [4, 5]
		},
		{
			name:     "Custom 2",
			nums:     []int{5, 1, 4, 2, 6},
			k:        1,
			expected: 18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxAlternatingSum(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("maxAlternatingSum(%v, %v) = %v; expected %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
