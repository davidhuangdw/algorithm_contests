package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchR2(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected bool
	}{
		{
			name:     "LeetCode Example 1",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   0,
			expected: true,
		},
		{
			name:     "LeetCode Example 2",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   3,
			expected: false,
		},
		{
			name:     "Target at beginning",
			nums:     []int{1, 0, 1, 1, 1},
			target:   1,
			expected: true,
		},
		{
			name:     "Target at end",
			nums:     []int{1, 1, 1, 0, 1},
			target:   1,
			expected: true,
		},
		{
			name:     "All duplicates",
			nums:     []int{1, 1, 1, 1, 1},
			target:   1,
			expected: true,
		},
		{
			name:     "All duplicates, target not found",
			nums:     []int{1, 1, 1, 1, 1},
			target:   2,
			expected: false,
		},
		{
			name:     "Single element found",
			nums:     []int{5},
			target:   5,
			expected: true,
		},
		{
			name:     "Single element not found",
			nums:     []int{5},
			target:   3,
			expected: false,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			target:   5,
			expected: false,
		},
		{
			name:     "Complex rotation with duplicates",
			nums:     []int{1, 1, 2, 0, 0, 0, 1},
			target:   2,
			expected: true,
		},
		{
			name:     "Target in left sorted part",
			nums:     []int{4, 5, 6, 0, 1, 2, 3},
			target:   5,
			expected: true,
		},
		{
			name:     "Target in right sorted part",
			nums:     []int{4, 5, 6, 0, 1, 2, 3},
			target:   2,
			expected: true,
		},
		{
			name:     "Target at pivot",
			nums:     []int{4, 5, 6, 0, 1, 2, 3},
			target:   0,
			expected: true,
		},
		{
			name:     "Multiple duplicates at pivot",
			nums:     []int{1, 0, 1, 1, 1},
			target:   0,
			expected: true,
		},
		{
			name:     "Large array with duplicates",
			nums:     []int{1, 1, 1, 2, 2, 3, 0, 0, 0, 1},
			target:   3,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := searchR2(tt.nums, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}
