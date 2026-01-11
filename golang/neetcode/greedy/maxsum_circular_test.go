package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubarraySumCircular(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "single positive number",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "single negative number",
			nums:     []int{-5},
			expected: -5,
		},
		{
			name:     "LeetCode example 1",
			nums:     []int{1, -2, 3, -2},
			expected: 3,
		},
		{
			name:     "LeetCode example 2",
			nums:     []int{5, -3, 5},
			expected: 10,
		},
		{
			name:     "LeetCode example 3",
			nums:     []int{-3, -2, -3},
			expected: -2,
		},
		{
			name:     "all positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "all negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			expected: -1,
		},
		{
			name:     "circular wrap gives maximum",
			nums:     []int{3, -1, 2, -1},
			expected: 4,
		},
		{
			name:     "no circular wrap needed",
			nums:     []int{2, -1, 2},
			expected: 4,
		},
		{
			name:     "mixed with zero",
			nums:     []int{0, 5, -3, 5},
			expected: 10,
		},
		{
			name:     "large circular sum",
			nums:     []int{-2, 4, -5, 4, -5, 9},
			expected: 11,
		},
		{
			name:     "edge case: all zeros",
			nums:     []int{0, 0, 0, 0},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSubarraySumCircular(tt.nums)
			assert.Equal(t, tt.expected, result, "maxSubarraySumCircular(%v)", tt.nums)
		})
	}
}
