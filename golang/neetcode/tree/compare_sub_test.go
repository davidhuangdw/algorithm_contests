package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		name     string
		root     []interface{}
		subRoot  []interface{}
		expected bool
	}{
		// Basic cases
		{"both nil", []interface{}{}, []interface{}{}, true},
		{"root nil, subRoot not nil", []interface{}{}, []interface{}{1}, false},
		{"subRoot nil", []interface{}{1}, []interface{}{}, true},
		{"identical single nodes", []interface{}{1}, []interface{}{1}, true},
		{"different single nodes", []interface{}{1}, []interface{}{2}, false},

		// Classic cases - subtree at root
		{"subtree at root", []interface{}{3, 4, 5, 1, 2}, []interface{}{4, 1, 2}, true},
		{"LeetCode example 1", []interface{}{3, 4, 5, 1, 2}, []interface{}{4, 1, 2}, true},
		{"LeetCode example 2", []interface{}{3, 4, 5, 1, 2, nil, nil, nil, nil, 0}, []interface{}{4, 1, 2}, false},

		// Classic cases - subtree in left/right
		{"subtree in left", []interface{}{3, 4, 5, 1, 2, 6, 7}, []interface{}{4, 1, 2}, true},
		{"subtree in right", []interface{}{3, 4, 5, 1, 2, 6, 7}, []interface{}{5, 6, 7}, true},
		{"subtree deep in left", []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}, []interface{}{4, 8, 9}, true},

		// Edge cases
		{"subtree with different structure", []interface{}{1, 2, 3}, []interface{}{1, 2}, false},
		{"subtree with extra nodes", []interface{}{1, 2, 3}, []interface{}{1, 2, 3, 4}, false},
		{"same values different structure", []interface{}{1, 2, 1}, []interface{}{1, 1, 2}, false},

		// Complex cases
		{"large tree with subtree", []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, []interface{}{6, 12, 13}, true},
		{"subtree with same values but different structure", []interface{}{1, 2, 3, 4, 5, 6, 7}, []interface{}{2, 4, 5}, true},
		{"negative numbers", []interface{}{-1, -2, -3, -4, -5}, []interface{}{-2, -4, -5}, true},
		{"mixed positive and negative", []interface{}{1, -2, 3, -4, 5}, []interface{}{-2, -4, nil}, false},
		{"duplicate values valid subtree", []interface{}{2, 2, 2, 3, 4, 5, 6}, []interface{}{2, 3, 4}, true},
		{"duplicate values invalid subtree", []interface{}{2, 2, 2, 3, 4, 5, 6}, []interface{}{2, 4, 5}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create trees from slice representations
			root := createTree(tt.root)
			subRoot := createTree(tt.subRoot)

			// Check if subRoot is a subtree of root
			result := isSubtree(root, subRoot)

			assert.Equal(t, tt.expected, result,
				"isSubtree(%v, %v) should return %v, got %v",
				tt.root, tt.subRoot, tt.expected, result)
		})
	}
}
