package arrays_hashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "LeetCode example 1 - mixed colors",
			nums:     []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "LeetCode example 2 - already sorted",
			nums:     []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "All twos",
			nums:     []int{2, 2, 2, 2},
			expected: []int{2, 2, 2, 2},
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "Reverse order",
			nums:     []int{2, 2, 1, 1, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Complex pattern with many swaps",
			nums:     []int{2, 0, 1, 0, 2, 1, 0, 2, 1},
			expected: []int{0, 0, 0, 1, 1, 1, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy to avoid modifying the original test data
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			
			sortColors(nums)
			assert.Equal(t, tt.expected, nums)
		})
	}
}