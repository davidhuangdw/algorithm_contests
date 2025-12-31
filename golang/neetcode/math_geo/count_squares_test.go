package math_geo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSquares(t *testing.T) {
	tests := []struct {
		name     string
		adds     [][]int
		counts   [][]int
		expected []int
	}{
		{
			name:     "single point",
			adds:     [][]int{{1, 1}},
			counts:   [][]int{{1, 1}},
			expected: []int{0},
		},
		{
			name:     "two points no square",
			adds:     [][]int{{1, 1}, {2, 2}},
			counts:   [][]int{{1, 1}, {2, 2}},
			expected: []int{0, 0},
		},
		{
			name:     "three points forming square",
			adds:     [][]int{{1, 1}, {1, 2}, {2, 1}},
			counts:   [][]int{{2, 2}},
			expected: []int{1},
		},
		{
			name:     "four points forming two squares",
			adds:     [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
			counts:   [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "points on same x-axis",
			adds:     [][]int{{1, 1}, {1, 2}, {1, 3}},
			counts:   [][]int{{1, 1}, {1, 2}, {1, 3}},
			expected: []int{0, 0, 0},
		},
		{
			name:     "points on same y-axis",
			adds:     [][]int{{1, 1}, {2, 1}, {3, 1}},
			counts:   [][]int{{1, 1}, {2, 1}, {3, 1}},
			expected: []int{0, 0, 0},
		},
		{
			name:     "large square",
			adds:     [][]int{{0, 0}, {0, 5}, {5, 0}, {5, 5}},
			counts:   [][]int{{0, 0}, {0, 5}, {5, 0}, {5, 5}},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "duplicate points",
			adds:     [][]int{{1, 1}, {1, 1}, {1, 2}, {2, 1}},
			counts:   [][]int{{2, 2}},
			expected: []int{2},
		},
		{
			name:     "negative coordinates",
			adds:     [][]int{{-1, -1}, {-1, 0}, {0, -1}, {0, 0}},
			counts:   [][]int{{-1, -1}, {-1, 0}, {0, -1}, {0, 0}},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "multiple squares from same point",
			adds:     [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
			counts:   [][]int{{1, 1}},
			expected: []int{4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := Constructor()
			
			// Add all points
			for _, point := range tt.adds {
				co.Add(point)
			}
			
			// Count squares for each test point
			for i, point := range tt.counts {
				result := co.Count(point)
				assert.Equal(t, tt.expected[i], result, "Count(%v) = %d, want %d", point, result, tt.expected[i])
			}
		})
	}
}