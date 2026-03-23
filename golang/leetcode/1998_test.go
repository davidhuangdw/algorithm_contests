package main

import "testing"

func TestGcdSort(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1: Sortable",
			nums:     []int{7, 21, 3},
			expected: true,
		},
		{
			name:     "Example 2: Not sortable",
			nums:     []int{5, 2, 6, 2},
			expected: false,
		},
		{
			name:     "Example 3: Sortable with duplicates",
			nums:     []int{10, 5, 9, 3, 15},
			expected: true,
		},
		{
			name:     "Already sorted",
			nums:     []int{2, 3, 5, 7},
			expected: true,
		},
		{
			name:     "All primes",
			nums:     []int{3, 2, 5, 7},
			expected: false,
		},
		{
			name:     "All same number",
			nums:     []int{6, 6, 6},
			expected: true,
		},
		{
			name:     "Complex sortable case",
			nums:     []int{8, 9, 4, 2, 3, 6},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdSort(tt.nums); got != tt.expected {
				t.Errorf("gcdSort() = %v, want %v", got, tt.expected)
			}
		})
	}
}
