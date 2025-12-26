package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name     string
		nodes    []interface{}
		expected int
	}{
		// Basic cases
		{"single positive node", []interface{}{5}, 5},
		{"single negative node", []interface{}{-5}, -5},
		{"two nodes positive", []interface{}{1, 2}, 3},
		{"two nodes negative", []interface{}{-1, -2}, -1},

		// Classic cases - LeetCode examples
		{"LeetCode example 1", []interface{}{1, 2, 3}, 6},
		{"LeetCode example 2", []interface{}{-10, 9, 20, nil, nil, 15, 7}, 42},
		{"LeetCode example 3", []interface{}{2, -1}, 2},

		// Edge cases
		{"all negative", []interface{}{-1, -2, -3}, -1},
		{"mixed with zero", []interface{}{-1, 0, -2}, 0},
		{"single branch", []interface{}{1, nil, 2, nil, 3}, 6},

		// Complex cases
		{"large balanced tree", []interface{}{10, 5, 15, 3, 7, 12, 18}, 55},
		{"negative root with positive children", []interface{}{-10, 20, 30}, 40},
		{"path through root", []interface{}{10, 5, 15, nil, nil, nil, 20}, 50},
		{"zigzag path", []interface{}{1, 2, nil, 3, nil, 4}, 10},
		{"all positive", []interface{}{1, 2, 3, 4, 5, 6, 7}, 18},
		{"large negative path", []interface{}{-1, -2, -3, -4, -5, -6, -7}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.nodes)
			result := maxPathSum(root)
			assert.Equal(t, tt.expected, result,
				"maxPathSum(%v) = %d, expected %d",
				tt.nodes, result, tt.expected)
		})
	}
}
