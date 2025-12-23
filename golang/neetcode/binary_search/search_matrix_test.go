package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		target   int
		expected bool
	}{
		// Basic cases
		{
			name: "Single element matrix - target found",
			matrix: [][]int{
				{5},
			},
			target:   5,
			expected: true,
		},
		{
			name: "Single element matrix - target not found",
			matrix: [][]int{
				{5},
			},
			target:   3,
			expected: false,
		},
		{
			name: "2x2 matrix - target found",
			matrix: [][]int{
				{1, 3},
				{5, 7},
			},
			target:   3,
			expected: true,
		},
		{
			name: "2x2 matrix - target not found",
			matrix: [][]int{
				{1, 3},
				{5, 7},
			},
			target:   4,
			expected: false,
		},

		// Classic examples
		{
			name: "Classic 3x4 matrix - target found",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   3,
			expected: true,
		},
		{
			name: "Classic 3x4 matrix - target found in last row",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   34,
			expected: true,
		},
		{
			name: "Classic 3x4 matrix - target not found",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   13,
			expected: false,
		},

		// Edge cases
		{
			name:     "Empty matrix",
			matrix:   [][]int{},
			target:   5,
			expected: false,
		},
		{
			name: "Single row matrix - target found",
			matrix: [][]int{
				{1, 3, 5, 7, 9},
			},
			target:   5,
			expected: true,
		},
		{
			name: "Single row matrix - target not found",
			matrix: [][]int{
				{1, 3, 5, 7, 9},
			},
			target:   6,
			expected: false,
		},
		{
			name: "Single column matrix - target found",
			matrix: [][]int{
				{1},
				{3},
				{5},
			},
			target:   3,
			expected: true,
		},
		{
			name: "Target smaller than first element",
			matrix: [][]int{
				{10, 20},
				{30, 40},
			},
			target:   5,
			expected: false,
		},
		{
			name: "Target larger than last element",
			matrix: [][]int{
				{10, 20},
				{30, 40},
			},
			target:   50,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := searchMatrix(tt.matrix, tt.target)
			assert.Equal(t, tt.expected, result,
				"searchMatrix(matrix=%v, target=%d) = %v, expected %v",
				tt.matrix, tt.target, result, tt.expected)
		})
	}
}
