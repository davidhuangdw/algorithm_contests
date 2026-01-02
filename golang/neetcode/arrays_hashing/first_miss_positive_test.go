package arrays_hashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "LeetCode example 1 - basic case",
			nums:     []int{1, 2, 0},
			expected: 3,
		},
		{
			name:     "LeetCode example 2 - with duplicates and negatives",
			nums:     []int{3, 4, -1, 1},
			expected: 2,
		},
		{
			name:     "LeetCode example 3 - all large numbers",
			nums:     []int{7, 8, 9, 11, 12},
			expected: 1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 1,
		},
		{
			name:     "Single element - missing 1",
			nums:     []int{2},
			expected: 1,
		},
		{
			name:     "Single element - 1 present",
			nums:     []int{1},
			expected: 2,
		},
		{
			name:     "All numbers present in sequence",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 6,
		},
		{
			name:     "Mixed positive and negative numbers",
			nums:     []int{-1, -2, -3, 1, 3, 4},
			expected: 2,
		},
		{
			name:     "Large numbers with gaps",
			nums:     []int{100, 101, 102, 1, 3},
			expected: 2,
		},
		{
			name:     "Duplicates and missing numbers",
			nums:     []int{1, 1, 2, 2, 3, 3, 5},
			expected: 4,
		},
		{
			name:     "some test",
			nums:     []int{1, 2, 4, 5, 6, 3, 1},
			expected: 7,
		},
		{
			name:     "All negative numbers",
			nums:     []int{-1, -2, -3, -4},
			expected: 1,
		},
		{
			name:     "Zero and negative numbers only",
			nums:     []int{0, -1, -2},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy to avoid modifying the original test data
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)

			result := firstMissingPositive(nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}
