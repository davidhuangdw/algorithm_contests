package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRotated(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		// Basic cases
		{
			name:     "Single element - target found",
			nums:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "Single element - target not found",
			nums:     []int{5},
			target:   3,
			expected: -1,
		},
		{
			name:     "Two elements sorted - target found",
			nums:     []int{1, 3},
			target:   3,
			expected: 1,
		},
		{
			name:     "Two elements rotated - target found",
			nums:     []int{3, 1},
			target:   1,
			expected: 1,
		},
		
		// Classic examples
		{
			name:     "Classic example 1 - target found",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "Classic example 1 - target not found",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "Classic example 2 - target found",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   5,
			expected: 1,
		},
		{
			name:     "Classic example 3",
			nums:     []int{1},
			target:   0,
			expected: -1,
		},
		
		// Edge cases
		{
			name:     "Empty array",
			nums:     []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "Already sorted - target found",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "Already sorted - target not found",
			nums:     []int{1, 2, 3, 4, 5},
			target:   6,
			expected: -1,
		},
		{
			name:     "Fully rotated - target found",
			nums:     []int{5, 1, 2, 3, 4},
			target:   1,
			expected: 1,
		},
		{
			name:     "All same elements - target found",
			nums:     []int{3, 3, 3, 3},
			target:   3,
			expected: 1, // Can be any index with value 3
		},
		{
			name:     "All same elements - target not found",
			nums:     []int{3, 3, 3, 3},
			target:   5,
			expected: -1,
		},
		
		// Complex cases
		{
			name:     "Large rotation - target found",
			nums:     []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5},
			target:   1,
			expected: 5,
		},
		{
			name:     "Large rotation - target in left part",
			nums:     []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5},
			target:   8,
			expected: 2,
		},
		{
			name:     "Small rotation - target found",
			nums:     []int{2, 3, 4, 5, 6, 7, 8, 9, 1},
			target:   1,
			expected: 8,
		},
		{
			name:     "Target at rotation point",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "Target smaller than all",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   -1,
			expected: -1,
		},
		{
			name:     "Target larger than all",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   10,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := search(tt.nums, tt.target)
			assert.Equal(t, tt.expected, result, 
				"search(%v, %d) = %d, expected %d",
				tt.nums, tt.target, result, tt.expected)
		})
	}
}