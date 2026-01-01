package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestIncreasingPath(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected int
	}{
		{
			name:     "leetcode example 1",
			matrix:   [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}},
			expected: 4,
		},
		{
			name:     "leetcode example 2",
			matrix:   [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}},
			expected: 4,
		},
		{
			name:     "single element",
			matrix:   [][]int{{1}},
			expected: 1,
		},
		{
			name:     "all same values",
			matrix:   [][]int{{1, 1}, {1, 1}},
			expected: 1,
		},
		{
			name:     "strictly increasing",
			matrix:   [][]int{{1, 2}, {4, 3}},
			expected: 4,
		},
		{
			name:     "strictly decreasing",
			matrix:   [][]int{{4, 3}, {1, 2}},
			expected: 4,
		},
		{
			name:     "spiral path",
			matrix:   [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
			expected: 9,
		},
		{
			name:     "diagonal path",
			matrix:   [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
			expected: 5,
		},
		{
			name:     "large matrix",
			matrix:   [][]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}, {4, 5, 6, 7}},
			expected: 7,
		},
		{
			name:     "complex path",
			matrix:   [][]int{{7, 8, 9}, {9, 7, 6}, {7, 2, 3}},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestIncreasingPath(tt.matrix)
			assert.Equal(t, tt.expected, result, "longestIncreasingPath(%v) = %d, want %d", tt.matrix, result, tt.expected)
		})
	}
}
