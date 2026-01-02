package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		name     string
		edges    [][]int
		expected []int
	}{
		{
			name:     "LeetCode example 1 - simple cycle",
			edges:    [][]int{{1, 2}, {1, 3}, {2, 3}},
			expected: []int{2, 3},
		},
		{
			name:     "LeetCode example 2 - larger cycle",
			edges:    [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			expected: []int{1, 4},
		},
		{
			name:     "Cycle with multiple connections",
			edges:    [][]int{{1, 2}, {2, 3}, {3, 1}, {3, 4}},
			expected: []int{3, 1},
		},
		{
			name:     "Star pattern with cycle",
			edges:    [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 4}},
			expected: []int{2, 4},
		},
		{
			name:     "Complex graph with multiple paths",
			edges:    [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 1}, {2, 6}, {6, 7}},
			expected: []int{5, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findRedundantConnection(tt.edges)
			assert.Equal(t, tt.expected, result)
		})
	}
}
