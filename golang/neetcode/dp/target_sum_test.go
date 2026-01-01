package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{1, 1, 1, 1, 1},
			target:   3,
			expected: 5,
		},
		{
			name:     "leetcode example 2",
			nums:     []int{1},
			target:   1,
			expected: 1,
		},
		{
			name:     "single element zero target",
			nums:     []int{0},
			target:   0,
			expected: 2,
		},
		{
			name:     "impossible target",
			nums:     []int{1, 2, 3},
			target:   10,
			expected: 0,
		},
		{
			name:     "negative target",
			nums:     []int{1, 1, 1, 1, 1},
			target:   -3,
			expected: 5,
		},
		{
			name:     "multiple ways",
			nums:     []int{1, 2, 3},
			target:   0,
			expected: 2,
		},
		{
			name:     "large numbers",
			nums:     []int{100, 200, 300},
			target:   400,
			expected: 1,
		},
		{
			name:     "zero elements",
			nums:     []int{},
			target:   0,
			expected: 1,
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0, 0},
			target:   0,
			expected: 16,
		},
		{
			name:     "mixed numbers",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 3,
		},
		{
			name:     "edge case with zero",
			nums:     []int{0, 1},
			target:   1,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findTargetSumWays(tt.nums, tt.target)
			assert.Equal(t, tt.expected, result, "findTargetSumWays(%v, %d) = %d, want %d", tt.nums, tt.target, result, tt.expected)
		})
	}
}