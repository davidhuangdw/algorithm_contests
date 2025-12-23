package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		// Basic cases
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Two elements sorted",
			nums:     []int{1, 2},
			expected: 1,
		},
		{
			name:     "Two elements rotated",
			nums:     []int{2, 1},
			expected: 1,
		},
		{
			name:     "Three elements sorted",
			nums:     []int{1, 2, 3},
			expected: 1,
		},
		{
			name:     "Three elements rotated once",
			nums:     []int{2, 3, 1},
			expected: 1,
		},
		{
			name:     "Three elements rotated twice",
			nums:     []int{3, 1, 2},
			expected: 1,
		},
		
		// Classic examples
		{
			name:     "Classic example 1",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "Classic example 2",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "Classic example 3",
			nums:     []int{11, 13, 15, 17},
			expected: 11,
		},
		
		// Edge cases
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Already sorted",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Fully rotated",
			nums:     []int{5, 1, 2, 3, 4},
			expected: 1,
		},
		{
			name:     "All same elements",
			nums:     []int{3, 3, 3, 3},
			expected: 3,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-5, -3, -1, 0, 2},
			expected: -5,
		},
		{
			name:     "Negative numbers rotated",
			nums:     []int{2, -5, -3, -1, 0},
			expected: -5,
		},
		
		// Complex cases
		{
			name:     "Large rotation",
			nums:     []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Small rotation",
			nums:     []int{2, 3, 4, 5, 6, 7, 8, 9, 1},
			expected: 1,
		},
		{
			name:     "Minimum at beginning",
			nums:     []int{0, 1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "Minimum at end",
			nums:     []int{1, 2, 3, 4, 5, 0},
			expected: 0,
		},
		{
			name:     "Minimum in middle",
			nums:     []int{4, 5, 0, 1, 2, 3},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMin(tt.nums)
			assert.Equal(t, tt.expected, result, 
				"findMin(%v) = %d, expected %d",
				tt.nums, result, tt.expected)
		})
	}
}