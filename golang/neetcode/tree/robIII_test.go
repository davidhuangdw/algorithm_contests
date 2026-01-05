package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	tests := []struct {
		name     string
		tree     []interface{}
		expected int
	}{
		// Basic cases
		{"empty tree", []interface{}{}, 0},
		{"single node", []interface{}{3}, 3},
		{"single node negative", []interface{}{-5}, 0},

		// Classic cases - LeetCode examples
		{"LeetCode example 1", []interface{}{3, 2, 3, nil, 3, nil, 1}, 7},
		{"LeetCode example 2", []interface{}{3, 4, 5, 1, 3, nil, 1}, 9},

		// Simple trees
		{"two nodes - parent and left child", []interface{}{3, 4}, 4},
		{"two nodes - parent and right child", []interface{}{3, nil, 4}, 4},
		{"three nodes - parent and two children", []interface{}{3, 4, 5}, 9},

		// Complex cases
		{"balanced tree with optimal skipping", []interface{}{10, 1, 1, 100, 100}, 210},
		{"zigzag pattern", []interface{}{1, 2, nil, 3, nil, 4}, 6},
		{"alternating levels", []interface{}{10, 5, 15, 1, 1, 10, 10}, 32},

		// Edge cases
		{"all negative values", []interface{}{-1, -2, -3}, 0},
		{"mixed positive and negative", []interface{}{5, -3, 8, -1, -2, 6, 9}, 20},
		{"large tree", []interface{}{10, 5, 15, 3, 7, 12, 18, 1, 4, 6, 8, 11, 13, 16, 20}, 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the tree from slice representation
			root := createTree(tt.tree)

			// Calculate maximum robbery amount
			result := rob(root)

			assert.Equal(t, tt.expected, result,
				"rob(%v) should return %d, got %d",
				tt.tree, tt.expected, result)
		})
	}
}
