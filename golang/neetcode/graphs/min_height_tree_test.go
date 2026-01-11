package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinHeightTrees(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected []int
	}{
		{
			name:     "single node",
			n:        1,
			edges:    [][]int{},
			expected: []int{0},
		},
		{
			name:     "two nodes",
			n:        2,
			edges:    [][]int{{0, 1}},
			expected: []int{0, 1},
		},
		{
			name:     "three nodes in line",
			n:        3,
			edges:    [][]int{{0, 1}, {1, 2}},
			expected: []int{1},
		},
		{
			name:     "four nodes in line",
			n:        4,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}},
			expected: []int{1, 2},
		},
		{
			name:     "LeetCode example 1",
			n:        4,
			edges:    [][]int{{1, 0}, {1, 2}, {1, 3}},
			expected: []int{1},
		},
		{
			name:     "LeetCode example 2",
			n:        6,
			edges:    [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			expected: []int{3, 4},
		},
		{
			name:     "star graph with center",
			n:        5,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
			expected: []int{0},
		},
		{
			name:     "complex tree with multiple centers",
			n:        7,
			edges:    [][]int{{0, 1}, {1, 2}, {1, 3}, {0, 4}, {4, 5}, {4, 6}},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinHeightTrees(tt.n, tt.edges)
			assert.ElementsMatch(t, tt.expected, result, "findMinHeightTrees(%d, %v)", tt.n, tt.edges)
		})
	}
}
