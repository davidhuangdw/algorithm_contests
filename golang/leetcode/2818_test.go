package main

import (
	"testing"
)

func TestMaximumScore(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{8, 3, 9, 3, 8},
			k:        2,
			expected: 81,
		},
		{
			name:     "Example 2",
			nums:     []int{19, 12, 14, 6, 10, 18},
			k:        3,
			expected: 4788,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			k:        1,
			expected: 5,
		},
		{
			name:     "k larger than total subarrays",
			nums:     []int{2, 2},
			k:        10,
			expected: 8, // Subarrays: [2] (idx 0), [2] (idx 1), [2,2] (idx 0). Elements: 2, 2, 2. Product: 8.
		},
		{
			name:     "All same prime scores, different values",
			nums:     []int{10, 5, 10},
			k:        2,
			expected: 100, // Prime scores: [2, 1, 2]. 
			// Val 10 at idx 0 is max for subarrays starting at 0: [10], [10,5]. (2 subarrays)
			// Val 5 at idx 1 is max for subarrays starting at 1: [5]. (1 subarray)
			// Val 10 at idx 2 is max for subarrays ending at 2: [10,5,10], [5,10], [10]. (3 subarrays)
			// k=2, take 10 twice. 10*10=100.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumScore(tt.nums, tt.k)
			if got != tt.expected {
				t.Errorf("maximumScore(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}
