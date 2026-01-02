package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwimInWater(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "LeetCode example 1 - simple 2x2 grid",
			grid: [][]int{
				{0, 2},
				{1, 3},
			},
			expected: 3,
		},
		{
			name: "LeetCode example 2 - 3x3 grid",
			grid: [][]int{
				{0, 1, 2, 3, 4},
				{24, 23, 22, 21, 5},
				{12, 13, 14, 15, 16},
				{11, 17, 18, 19, 20},
				{10, 9, 8, 7, 6},
			},
			expected: 16,
		},
		{
			name: "Single cell grid",
			grid: [][]int{
				{5},
			},
			expected: 5,
		},
		{
			name: "2x1 grid",
			grid: [][]int{
				{3},
				{1},
			},
			expected: 3,
		},
		{
			name: "1x2 grid",
			grid: [][]int{
				{2, 5},
			},
			expected: 5,
		},
		{
			name: "Increasing values along path",
			grid: [][]int{
				{0, 1},
				{2, 3},
			},
			expected: 3,
		},
		{
			name: "Alternative path with lower maximum",
			grid: [][]int{
				{0, 5},
				{3, 4},
			},
			expected: 4,
		},
		{
			name: "Complex 3x3 grid with multiple paths",
			grid: [][]int{
				{0, 8, 7},
				{2, 1, 6},
				{3, 4, 5},
			},
			expected: 5,
		},
		{
			name: "Grid with high values requiring detour",
			grid: [][]int{
				{0, 10, 2},
				{1, 9, 3},
				{4, 8, 5},
			},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := swimInWater(tt.grid)
			assert.Equal(t, tt.expected, result)
		})
	}
}
