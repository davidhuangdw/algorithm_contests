package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "LeetCode example 1",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "LeetCode example 2",
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			expected: -1,
		},
		{
			name: "LeetCode example 3",
			grid: [][]int{
				{0, 2},
			},
			expected: 0,
		},
		{
			name: "All rotten initially",
			grid: [][]int{
				{2, 2},
				{2, 2},
			},
			expected: 0,
		},
		{
			name: "All fresh initially",
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
			expected: -1,
		},
		{
			name: "Single rotten orange",
			grid: [][]int{
				{2},
			},
			expected: 0,
		},
		{
			name: "Single fresh orange",
			grid: [][]int{
				{1},
			},
			expected: -1,
		},
		{
			name: "Empty cell",
			grid: [][]int{
				{0},
			},
			expected: 0,
		},
		{
			name: "Immediate neighbors",
			grid: [][]int{
				{2, 1, 2},
				{1, 1, 1},
				{2, 1, 2},
			},
			expected: 2,
		},
		{
			name: "Complex pattern",
			grid: [][]int{
				{2, 0, 1, 1, 1},
				{1, 0, 1, 0, 1},
				{1, 1, 1, 0, 2},
			},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a deep copy of the input grid
			gridCopy := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}

			result := orangesRotting(gridCopy)
			assert.Equal(t, tt.expected, result)
		})
	}
}
