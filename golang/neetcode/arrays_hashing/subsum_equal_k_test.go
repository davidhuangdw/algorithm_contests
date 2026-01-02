package arrays_hashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "LeetCode example 1 - basic case",
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			name:     "LeetCode example 2 - with negative numbers",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2,
		},
		{
			name:     "Single element equal to k",
			nums:     []int{5},
			k:        5,
			expected: 1,
		},
		{
			name:     "Single element not equal to k",
			nums:     []int{3},
			k:        5,
			expected: 0,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			k:        0,
			expected: 0,
		},
		{
			name:     "All zeros with k=0",
			nums:     []int{0, 0, 0, 0},
			k:        0,
			expected: 10, // n*(n+1)/2 = 4*5/2 = 10
		},
		{
			name:     "Negative numbers and zero sum",
			nums:     []int{-1, -1, 1},
			k:        0,
			expected: 1,
		},
		{
			name:     "Complex pattern with multiple subarrays",
			nums:     []int{3, 4, 7, 2, -3, 1, 4, 2},
			k:        7,
			expected: 4,
		},
		{
			name:     "Large k value",
			nums:     []int{1, 2, 3},
			k:        100,
			expected: 0,
		},
		{
			name:     "Negative k value",
			nums:     []int{1, -2, 3, -4},
			k:        -1,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subarraySum(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}