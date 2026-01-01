package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		grid     [][]int
		expected int
	}{
		// LeetCode examples
		{
			[][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			}, 6,
		},

		// Edge cases
		{[][]int{{0}}, 0},
		{[][]int{{1}}, 1},
		{[][]int{{0, 0}, {0, 0}}, 0},
		{[][]int{{1, 1}, {1, 1}}, 4},

		// Simple cases
		{[][]int{{1, 0}, {0, 1}}, 1},
		{[][]int{{1, 1}, {0, 1}}, 3},
		{[][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}, 1},

		// Complex cases
		{
			[][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 1},
			}, 4,
		},
		{
			[][]int{
				{1, 0, 1, 1, 0},
				{1, 1, 0, 0, 1},
				{0, 1, 1, 1, 0},
				{1, 0, 0, 0, 1},
			}, 6,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			// Create a copy of the grid since the function modifies it
			gridCopy := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}
			result := maxAreaOfIsland(gridCopy)
			assert.Equal(t, tt.expected, result)
		})
	}
}
