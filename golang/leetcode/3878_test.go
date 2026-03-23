package main

import "testing"

func TestCountGoodSubarrays(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{
			name:     "Example 1",
			nums:     []int{4, 2, 3},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 3, 1},
			expected: 6,
		},
		{
			name:     "Example 3",
			nums:     []int{10, 6, 4, 2},
			expected: 6,
		},
		{
			name:     "Example 4",
			nums:     []int{6, 6},
			expected: 3,
		},
		{
			name:     "Single element",
			nums:     []int{100},
			expected: 1,
		},
		{
			name:     "All same elements",
			nums:     []int{7, 7, 7},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodSubarrays(tt.nums); got != tt.expected {
				t.Errorf("countGoodSubarrays() = %v, want %v", got, tt.expected)
			}
		})
	}
}
