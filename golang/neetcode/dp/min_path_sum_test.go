package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "empty grid",
			grid: [][]int{},
			want: 0,
		},
		{
			name: "1x1 grid",
			grid: [][]int{{5}},
			want: 5,
		},
		{
			name: "2x2 grid",
			grid: [][]int{{1, 3}, {1, 5}},
			want: 7, // 1→1→5 = 7
		},
		{
			name: "LeetCode example 1",
			grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			want: 7, // 1→3→1→1→1 = 7
		},
		{
			name: "single row",
			grid: [][]int{{1, 2, 3}},
			want: 6, // 1+2+3 = 6
		},
		{
			name: "single column",
			grid: [][]int{{1}, {2}, {3}},
			want: 6, // 1+2+3 = 6
		},
		{
			name: "all same values",
			grid: [][]int{{2, 2}, {2, 2}},
			want: 6, // 2+2+2 = 6
		},
		{
			name: "negative values",
			grid: [][]int{{-1, -2}, {-3, -4}},
			want: -8, // -1→-2→-4 = -7 or -1→-3→-4 = -8, min is -8
		},
		{
			name: "mixed positive and negative",
			grid: [][]int{{1, -2}, {3, -4}},
			want: -5, // 1→-2→-4 = -5 or 1→3→-4 = 0, min is -5
		},
		{
			name: "large grid",
			grid: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			want: 30, // 1→2→3→4→8→12 = 30
		},
		{
			name: "grid with zero values",
			grid: [][]int{{0, 0}, {0, 0}},
			want: 0,
		},
		{
			name: "minimum path through bottom",
			grid: [][]int{{1, 2}, {1, 1}},
			want: 3, // 1→1→1 = 3
		},
		{
			name: "minimum path through right",
			grid: [][]int{{1, 1}, {2, 1}},
			want: 3, // 1→1→1 = 3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minPathSum(tt.grid)
			assert.Equal(t, tt.want, result,
				"minPathSum(%v) = %d, want %d", tt.grid, result, tt.want)
		})
	}
}

func TestMinPathSumEdgeCases(t *testing.T) {
	t.Run("nil grid", func(t *testing.T) {
		result := minPathSum(nil)
		assert.Equal(t, 0, result)
	})

	t.Run("grid with one element zero", func(t *testing.T) {
		grid := [][]int{{0}}
		result := minPathSum(grid)
		assert.Equal(t, 0, result)
	})

	t.Run("grid with large values", func(t *testing.T) {
		grid := [][]int{{100, 200}, {300, 400}}
		result := minPathSum(grid)
		assert.Equal(t, 700, result) // 100→200→400 = 700
	})
}
