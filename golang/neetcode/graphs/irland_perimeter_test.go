package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslandPerimeter(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name:     "empty grid",
			grid:     [][]int{},
			expected: 0,
		},
		{
			name:     "single land cell",
			grid:     [][]int{{1}},
			expected: 4,
		},
		{
			name:     "single water cell",
			grid:     [][]int{{0}},
			expected: 0,
		},
		{
			name:     "two adjacent land cells horizontal",
			grid:     [][]int{{1, 1}},
			expected: 6,
		},
		{
			name:     "two adjacent land cells vertical",
			grid:     [][]int{{1}, {1}},
			expected: 6,
		},
		{
			name:     "leetcode example 1",
			grid:     [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}},
			expected: 16,
		},
		{
			name:     "2x2 land island",
			grid:     [][]int{{1, 1}, {1, 1}},
			expected: 8,
		},
		{
			name:     "single row with isolated land",
			grid:     [][]int{{1, 0, 1}},
			expected: 8,
		},
		{
			name:     "single column with isolated land",
			grid:     [][]int{{1}, {0}, {1}},
			expected: 8,
		},
		{
			name:     "L-shaped island",
			grid:     [][]int{{1, 0}, {1, 1}},
			expected: 8,
		},
		{
			name:     "all water grid",
			grid:     [][]int{{0, 0}, {0, 0}},
			expected: 0,
		},
		{
			name:     "large grid with complex shape",
			grid:     [][]int{{1, 1, 0, 0}, {1, 0, 0, 0}, {0, 0, 0, 1}, {0, 0, 1, 1}},
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := islandPerimeter(tt.grid)
			assert.Equal(t, tt.expected, result)
		})
	}
}