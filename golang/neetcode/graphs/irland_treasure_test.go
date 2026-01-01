package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslandsAndTreasure(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "LeetCode example 1",
			input: [][]int{
				{2147483647, -1, 0, 2147483647},
				{2147483647, 2147483647, 2147483647, -1},
				{2147483647, -1, 2147483647, -1},
				{0, -1, 2147483647, 2147483647},
			},
			expected: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			name: "Single treasure",
			input: [][]int{
				{2147483647, 2147483647},
				{2147483647, 0},
			},
			expected: [][]int{
				{2, 1},
				{1, 0},
			},
		},
		{
			name: "Multiple treasures",
			input: [][]int{
				{0, 2147483647},
				{2147483647, 0},
			},
			expected: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			name: "All obstacles",
			input: [][]int{
				{-1, -1},
				{-1, -1},
			},
			expected: [][]int{
				{-1, -1},
				{-1, -1},
			},
		},
		{
			name: "Single cell treasure",
			input: [][]int{{0}},
			expected: [][]int{{0}},
		},
		{
			name: "Treasure surrounded by obstacles",
			input: [][]int{
				{-1, -1, -1},
				{-1, 0, -1},
				{-1, -1, -1},
			},
			expected: [][]int{
				{-1, -1, -1},
				{-1, 0, -1},
				{-1, -1, -1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a deep copy of the input grid
			gridCopy := make([][]int, len(tt.input))
			for i := range tt.input {
				gridCopy[i] = make([]int, len(tt.input[i]))
				copy(gridCopy[i], tt.input[i])
			}

			islandsAndTreasure(gridCopy)

			assert.Equal(t, tt.expected, gridCopy)
		})
	}
}