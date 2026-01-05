package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		tree     []interface{}
		expected []int
	}{
		// Basic cases
		{"empty tree", []interface{}{}, []int{}},
		{"single node", []interface{}{1}, []int{1}},
		{"single node negative", []interface{}{-5}, []int{-5}},

		// Simple trees
		{"two nodes - left child", []interface{}{2, 1}, []int{1, 2}},
		{"two nodes - right child", []interface{}{1, nil, 2}, []int{2, 1}},
		{"three nodes - complete", []interface{}{2, 1, 3}, []int{1, 3, 2}},

		// Classic cases - LeetCode examples
		{"LeetCode example 1", []interface{}{1, nil, 2, nil, 3}, []int{3, 2, 1}},
		{"LeetCode example 2", []interface{}{}, []int{}},
		{"LeetCode example 3", []interface{}{1}, []int{1}},

		// Complex cases
		{"balanced tree", []interface{}{4, 2, 6, 1, 3, 5, 7}, []int{1, 3, 2, 5, 7, 6, 4}},
		{"left skewed", []interface{}{4, 3, nil, 2, nil, 1}, []int{1, 2, 3, 4}},
		{"right skewed", []interface{}{1, nil, 2, nil, 3}, []int{3, 2, 1}},

		// Edge cases
		{"incomplete tree", []interface{}{1, 2, nil, 3}, []int{3, 2, 1}},
		{"large tree", []interface{}{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15},
			[]int{1, 3, 2, 5, 7, 6, 4, 9, 11, 10, 13, 15, 14, 12, 8}},
		{"negative values", []interface{}{0, -2, 2, -3, -1, 1, 3}, []int{-3, -1, -2, 1, 3, 2, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the tree from slice representation
			root := createTree(tt.tree)

			// Perform postorder traversal
			result := postorderTraversal(root)

			assert.Equal(t, tt.expected, result,
				"postorderTraversal(%v) should return %v, got %v",
				tt.tree, tt.expected, result)
		})
	}
}
