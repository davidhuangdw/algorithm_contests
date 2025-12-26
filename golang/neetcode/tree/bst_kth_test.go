package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name     string
		tree     []interface{}
		k        int
		expected int
	}{
		// Basic cases
		{"single node k=1", []interface{}{5}, 1, 5},

		// Classic cases - balanced BSTs
		{"balanced BST k=1", []interface{}{3, 1, 4, nil, 2}, 1, 1},
		{"balanced BST k=3", []interface{}{3, 1, 4, nil, 2}, 3, 3},
		{"LeetCode example", []interface{}{3, 1, 4, nil, 2}, 1, 1},
		{"LeetCode example k=2", []interface{}{3, 1, 4, nil, 2}, 2, 2},
		{"LeetCode example k=3", []interface{}{3, 1, 4, nil, 2}, 3, 3},

		// Edge cases
		{"left skewed k=1", []interface{}{5, 3, nil, 1}, 1, 1},
		{"left skewed k=2", []interface{}{5, 3, nil, 1}, 2, 3},
		{"left skewed k=3", []interface{}{5, 3, nil, 1}, 3, 5},
		{"right skewed k=1", []interface{}{1, nil, 3, nil, nil, nil, 5}, 1, 1},
		{"right skewed k=2", []interface{}{1, nil, 3, nil, nil, nil, 5}, 2, 3},
		{"right skewed k=3", []interface{}{1, nil, 3, nil, 5}, 3, 5},

		// Complex cases
		{"large BST k=5", []interface{}{5, 3, 7, 1, 4, 6, 8}, 5, 6},
		{"large BST k=7", []interface{}{5, 3, 7, 1, 4, 6, 8}, 7, 8},
		{"negative values k=2", []interface{}{0, -3, 2, -5, -1}, 2, -3},
		{"mixed values k=3", []interface{}{10, 5, 15, 3, 7, 12, 18}, 3, 7},
		{"duplicate values k=4", []interface{}{2, 1, 3, 1, 2, 3}, 4, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.tree)
			result := kthSmallest(root, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}
