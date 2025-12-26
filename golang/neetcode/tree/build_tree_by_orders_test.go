package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		expected []interface{}
	}{
		// Basic cases
		{"empty trees", []int{}, []int{}, []interface{}{}},
		{"single node", []int{1}, []int{1}, []interface{}{1}},

		// Classic cases - balanced trees
		{"simple balanced", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []interface{}{3, 9, 20, nil, nil, 15, 7}},
		{"LeetCode example", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []interface{}{3, 9, 20, nil, nil, 15, 7}},
		{"three nodes left", []int{1, 2, 3}, []int{2, 1, 3}, []interface{}{1, 2, 3}},
		{"three nodes right", []int{1, 2, 3}, []int{1, 2, 3}, []interface{}{1, nil, 2, nil, 3}},

		// Edge cases
		{"left skewed", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}, []interface{}{1, 2, nil, 3, nil, 4}},
		{"right skewed", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, []interface{}{1, nil, 2, nil, 3, nil, 4}},

		// Complex cases
		{"large balanced", []int{1, 2, 4, 5, 3, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7}, []interface{}{1, 2, 3, 4, 5, 6, 7}},
		{"negative values", []int{-1, -2, -3}, []int{-3, -2, -1}, []interface{}{-1, -2, nil, -3}},
		{"mixed values", []int{5, -3, 0, -1, 8, 6, 9}, []int{0, -3, -1, 5, 6, 8, 9}, []interface{}{5, -3, 8, 0, -1, 6, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build tree from preorder and inorder
			root := buildTree(tt.preorder, tt.inorder)

			// Convert the built tree back to slice representation for comparison
			result := treeToSlice(root)

			assert.Equal(t, tt.expected, result)
		})
	}
}
