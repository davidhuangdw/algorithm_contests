package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostConnectPoints(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		expected int
	}{
		{
			name:     "LeetCode example 1 - simple 3 points",
			points:   [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}},
			expected: 20,
		},
		{
			name:     "LeetCode example 2 - 3 points in line",
			points:   [][]int{{3, 12}, {-2, 5}, {-4, 1}},
			expected: 18,
		},
		{
			name:     "Two points - direct connection",
			points:   [][]int{{0, 0}, {1, 1}},
			expected: 2,
		},
		{
			name:     "Single point - zero cost",
			points:   [][]int{{0, 0}},
			expected: 0,
		},
		{
			name:     "Four points in square",
			points:   [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
			expected: 3,
		},
		{
			name:     "Points in straight line",
			points:   [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
			expected: 3,
		},
		{
			name:     "Large grid with multiple paths",
			points:   [][]int{{0, 0}, {2, 0}, {0, 2}, {2, 2}, {1, 1}},
			expected: 8,
		},
		{
			name:     "Negative coordinates",
			points:   [][]int{{-1, -1}, {1, 1}, {-1, 1}, {1, -1}},
			expected: 6,
		},
		{
			name:     "Complex pattern with optimal path",
			points:   [][]int{{0, 0}, {4, 0}, {0, 4}, {4, 4}, {2, 2}},
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minCostConnectPoints(tt.points)
			assert.Equal(t, tt.expected, result)
		})
	}
}
