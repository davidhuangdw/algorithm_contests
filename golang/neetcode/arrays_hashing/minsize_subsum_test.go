package arrays_hashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		nums     []int
		expected int
	}{
		{
			name:     "leetcode example 1",
			target:   7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			expected: 2, // [4,3]
		},
		{
			name:     "leetcode example 2",
			target:   4,
			nums:     []int{1, 4, 4},
			expected: 1, // [4]
		},
		{
			name:     "leetcode example 3",
			target:   11,
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			expected: 0, // no subarray >= 11
		},
		{
			name:     "single element equal to target",
			target:   5,
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "single element less than target",
			target:   6,
			nums:     []int{5},
			expected: 0,
		},
		{
			name:     "empty array",
			target:   5,
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "all elements needed",
			target:   15,
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "negative numbers",
			target:   5,
			nums:     []int{-1, 2, 3, 4, -2, 6},
			expected: 1, // [6]
		},
		{
			name:     "multiple valid subarrays",
			target:   7,
			nums:     []int{2, 3, 1, 1, 5, 1, 2, 1},
			expected: 3,
		},
		{
			name:     "target is zero",
			target:   0,
			nums:     []int{1, 2, 3},
			expected: 1, // any non-empty subarray >= 0
		},
		{
			name:     "large numbers",
			target:   1000,
			nums:     []int{100, 200, 300, 400, 500},
			expected: 3, // [300,400,500] = 1200
		},
		{
			name:     "exact match needed",
			target:   10,
			nums:     []int{1, 2, 3, 4},
			expected: 4, // [1,2,3,4] = 10
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minSubArrayLen(tt.target, tt.nums)
			assert.Equal(t, tt.expected, result, "target=%d, nums=%v", tt.target, tt.nums)
		})
	}
}
