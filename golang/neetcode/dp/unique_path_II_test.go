package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name         string
		obstacleGrid [][]int
		want         int
	}{
		{
			name:         "1x1 grid without obstacle",
			obstacleGrid: [][]int{{0}},
			want:         1,
		},
		{
			name:         "1x1 grid with obstacle",
			obstacleGrid: [][]int{{1}},
			want:         0,
		},
		{
			name:         "2x2 grid without obstacles",
			obstacleGrid: [][]int{{0, 0}, {0, 0}},
			want:         2,
		},
		{
			name:         "3x3 grid without obstacles",
			obstacleGrid: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			want:         6,
		},
		{
			name:         "grid with obstacle in start",
			obstacleGrid: [][]int{{1, 0}, {0, 0}},
			want:         0,
		},
		{
			name:         "grid with obstacle in end",
			obstacleGrid: [][]int{{0, 0}, {0, 1}},
			want:         0,
		},
		{
			name:         "grid with obstacle blocking path",
			obstacleGrid: [][]int{{0, 1}, {0, 0}},
			want:         1,
		},
		{
			name:         "LeetCode example 1",
			obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			want:         2,
		},
		{
			name:         "LeetCode example 2",
			obstacleGrid: [][]int{{0, 1}, {0, 0}},
			want:         1,
		},
		{
			name:         "single row with obstacle",
			obstacleGrid: [][]int{{0, 0, 1, 0}},
			want:         0,
		},
		{
			name:         "single column with obstacle",
			obstacleGrid: [][]int{{0}, {0}, {1}, {0}},
			want:         0,
		},
		{
			name: "large grid with obstacles",
			obstacleGrid: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 0},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uniquePathsWithObstacles(tt.obstacleGrid)
			assert.Equal(t, tt.want, result,
				"uniquePathsWithObstacles(%v) = %d, want %d", tt.obstacleGrid, result, tt.want)
		})
	}
}

func TestUniquePathsWithObstaclesEdgeCases(t *testing.T) {
	t.Run("nil grid", func(t *testing.T) {
		result := uniquePathsWithObstacles(nil)
		assert.Equal(t, 0, result)
	})

	t.Run("grid with all obstacles", func(t *testing.T) {
		grid := [][]int{{1, 1}, {1, 1}}
		result := uniquePathsWithObstacles(grid)
		assert.Equal(t, 0, result)
	})

	t.Run("single row without obstacles", func(t *testing.T) {
		grid := [][]int{{0, 0, 0, 0}}
		result := uniquePathsWithObstacles(grid)
		assert.Equal(t, 1, result)
	})

	t.Run("single column without obstacles", func(t *testing.T) {
		grid := [][]int{{0}, {0}, {0}, {0}}
		result := uniquePathsWithObstacles(grid)
		assert.Equal(t, 1, result)
	})
}
