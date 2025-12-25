package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to find a node by value in the tree
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if val < root.Val {
		return findNode(root.Left, val)
	}
	return findNode(root.Right, val)
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name     string
		tree     []interface{}
		p        int
		q        int
		expected int
	}{
		// Basic cases
		{"root is LCA", []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 2, 8, 6},
		{"LCA in left subtree", []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 0, 5, 2},
		{"LCA in right subtree", []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 7, 9, 8},

		// Classic cases - LeetCode examples
		{"LeetCode example 1", []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 2, 8, 6},
		{"LeetCode example 2", []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 2, 4, 2},
		{"LeetCode example 3", []interface{}{2, 1}, 2, 1, 2},

		// Edge cases
		{"same nodes", []interface{}{6, 2, 8, 0, 4, 7, 9}, 2, 2, 2},
		{"nodes in straight line", []interface{}{5, 4, nil, 3, nil, 2, nil, 1}, 1, 4, 4},
		{"one node is ancestor", []interface{}{6, 2, 8, 0, 4, 7, 9}, 2, 3, 2},
		{"deep nodes", []interface{}{10, 5, 15, 3, 7, 12, 18, 1, 4, 6, 8, 11, 13, 16, 20}, 1, 8, 5},

		// Complex cases
		{"large tree LCA", []interface{}{20, 10, 30, 5, 15, 25, 35, 3, 7, 12, 17, 22, 27, 32, 37}, 3, 17, 10},
		{"negative numbers", []interface{}{0, -5, 5, -10, -3, 3, 10}, -10, 3, 0},
		{"mixed values", []interface{}{50, 25, 75, 10, 40, 60, 90, 5, 15, 35, 45, 55, 65, 85, 95}, 35, 45, 40},
		{"single node tree", []interface{}{1}, 1, 1, 1},
		{"two node tree", []interface{}{2, 1}, 1, 2, 2},

		// BST property tests
		{"p and q on different sides", []interface{}{10, 5, 15, 3, 7, 12, 18}, 3, 18, 10},
		{"p and q both on left", []interface{}{10, 5, 15, 3, 7, 12, 18}, 3, 7, 5},
		{"p and q both on right", []interface{}{10, 5, 15, 3, 7, 12, 18}, 12, 18, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the BST from slice representation
			root := createTree(tt.tree)

			// Find the p and q nodes
			pNode := findNode(root, tt.p)
			qNode := findNode(root, tt.q)

			// Find the lowest common ancestor
			result := lowestCommonAncestor(root, pNode, qNode)

			// Verify the result
			assert.NotNil(t, result, "LCA should not be nil")
			assert.Equal(t, tt.expected, result.Val,
				"lowestCommonAncestor(tree with root %v, %d, %d) should return %d, got %d",
				tt.tree[0], tt.p, tt.q, tt.expected, result.Val)
		})
	}
}
