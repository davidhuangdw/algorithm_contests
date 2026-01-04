package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "leetcode example 2",
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			name:     "k larger than array length",
			nums:     []int{1, 2, 3},
			k:        5,
			expected: []int{2, 3, 1}, // 5 % 3 = 2 rotations
		},
		{
			name:     "k equal to array length",
			nums:     []int{1, 2, 3, 4},
			k:        4,
			expected: []int{1, 2, 3, 4}, // no change
		},
		{
			name:     "k is zero",
			nums:     []int{1, 2, 3},
			k:        0,
			expected: []int{1, 2, 3}, // no change
		},
		{
			name:     "single element array",
			nums:     []int{5},
			k:        10,
			expected: []int{5}, // no change
		},
		{
			name:     "empty array",
			nums:     []int{},
			k:        5,
			expected: []int{}, // no change
		},
		{
			name:     "negative k (should be handled as positive)",
			nums:     []int{1, 2, 3, 4},
			k:        -1, // equivalent to k=3
			expected: []int{2, 3, 4, 1},
		},
		{
			name:     "multiple of array length",
			nums:     []int{1, 2, 3},
			k:        6,
			expected: []int{1, 2, 3}, // no change
		},
		{
			name:     "complex rotation",
			nums:     []int{1, 2, 3, 4, 5, 6},
			k:        4,
			expected: []int{3, 4, 5, 6, 1, 2},
		},
		{
			name:     "gcd not 1",
			nums:     []int{1, 2, 3, 4, 5, 6},
			k:        2,
			expected: []int{5, 6, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy to avoid modifying the original test data
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			
			rotate(nums, tt.k)
			assert.Equal(t, tt.expected, nums, "input=%v, k=%d", tt.nums, tt.k)
		})
	}
}