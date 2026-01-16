package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumEffortPath(t *testing.T) {
	tests := []struct {
		name        string
		heights     [][]int
		expected    int
		description string
	}{
		{
			name: "LeetCode Example 1",
			heights: [][]int{
				{1, 2, 2},
				{3, 8, 2},
				{5, 3, 5},
			},
			expected:    2,
			description: "Example with varying heights and minimum effort path",
		},
		{
			name: "LeetCode Example 2",
			heights: [][]int{
				{1, 2, 3},
				{3, 8, 4},
				{5, 3, 5},
			},
			expected:    1,
			description: "Example with diagonal path",
		},
		{
			name: "LeetCode Example 3",
			heights: [][]int{
				{1, 2, 1, 1, 1},
				{1, 2, 1, 2, 1},
				{1, 2, 1, 2, 1},
				{1, 2, 1, 2, 1},
				{1, 1, 1, 2, 1},
			},
			expected:    0,
			description: "Example with all 1's except some 2's, but path through 1's exists",
		},
		{
			name:        "Single cell matrix",
			heights:     [][]int{{10}},
			expected:    0,
			description: "Only one cell, effort is 0",
		},
		{
			name:        "1xN matrix",
			heights:     [][]int{{1, 3, 6, 10, 15}},
			expected:    5,
			description: "Horizontal path with increasing heights",
		},
		{
			name: "Nx1 matrix",
			heights: [][]int{
				{5},
				{10},
				{15},
				{20},
			},
			expected:    5,
			description: "Vertical path with increasing heights",
		},
		{
			name: "All same heights",
			heights: [][]int{
				{7, 7, 7},
				{7, 7, 7},
				{7, 7, 7},
			},
			expected:    0,
			description: "All cells have same height, effort is 0",
		},
		{
			name: "Path with steep climb",
			heights: [][]int{
				{1, 100, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected:    0,
			description: "Path can avoid steep climb by going around",
		},
		{
			name: "Only one possible path",
			heights: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			expected:    4,
			description: "Grid where each cell is 1 higher than the previous, minimum effort is 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumEffortPath(tt.heights)
			assert.Equal(t, tt.expected, result, "%s: Expected minimum effort %d, got %d", tt.description, tt.expected, result)
		})
	}
}
