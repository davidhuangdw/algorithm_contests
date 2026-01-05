package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name     string
		tree     []interface{}
		key      int
		expected []interface{}
	}{
		// Basic cases
		{"empty tree", []interface{}{}, 5, []interface{}{}},
		{"single node delete", []interface{}{5}, 5, []interface{}{}},
		{"single node not found", []interface{}{5}, 3, []interface{}{5}},

		// Classic cases - LeetCode examples
		{"LeetCode example 1", []interface{}{5, 3, 6, 2, 4, nil, 7}, 3, []interface{}{5, 4, 6, 2, nil, nil, 7}},
		{"LeetCode example 2", []interface{}{5, 3, 6, 2, 4, nil, 7}, 0, []interface{}{5, 3, 6, 2, 4, nil, 7}},

		// Edge cases
		{"delete root with two children", []interface{}{5, 3, 7}, 5, []interface{}{7, 3}},
		{"delete leaf node", []interface{}{5, 3, 7}, 3, []interface{}{5, nil, 7}},
		{"delete node with one child", []interface{}{5, 3, 7, nil, nil, 6}, 7, []interface{}{5, 3, 6}},

		// Complex cases
		{"delete from large tree", []interface{}{8, 3, 10, 1, 6, nil, 14, nil, nil, 4, 7, nil, nil, 13}, 6, []interface{}{8, 3, 10, 1, 7, nil, 14, nil, nil, 4, nil, nil, nil, 13}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the original tree
			root := createTree(tt.tree)

			// Delete the node
			result := deleteNode(root, tt.key)

			// Convert the result back to slice representation
			resultSlice := treeToSlice(result)

			assert.Equal(t, tt.expected, resultSlice,
				"deleteNode(%v, %d) should return %v, got %v",
				tt.tree, tt.key, tt.expected, resultSlice)
		})
	}
}
