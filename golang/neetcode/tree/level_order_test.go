package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		tree     []interface{}
		expected [][]int
	}{
		// Basic cases
		{"empty tree", []interface{}{}, nil},
		{"single node", []interface{}{1}, [][]int{{1}}},

		// Classic cases - balanced trees
		{"two level balanced", []interface{}{3, 9, 20, nil, nil, 15, 7}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"LeetCode example", []interface{}{3, 9, 20, nil, nil, 15, 7}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"three level balanced", []interface{}{1, 2, 3, 4, 5, 6, 7}, [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},

		// Edge cases - unbalanced trees
		{"left skewed", []interface{}{1, 2, nil, 3, nil, 4}, [][]int{{1}, {2}, {3}, {4}}},
		{"right skewed", []interface{}{1, nil, 2, nil, 3, nil, 4}, [][]int{{1}, {2}, {3}, {4}}},
		{"incomplete last level", []interface{}{1, 2, 3, 4, nil, nil, nil}, [][]int{{1}, {2, 3}, {4}}},

		// Complex cases
		{"large balanced tree", []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			[][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15}}},
		{"negative numbers", []interface{}{-1, -2, -3, -4, -5}, [][]int{{-1}, {-2, -3}, {-4, -5}}},
		{"mixed values", []interface{}{5, -3, 8, 0, -1, 6, 9}, [][]int{{5}, {-3, 8}, {0, -1, 6, 9}}},
		{"duplicate values", []interface{}{2, 2, 2, 3, 3, 3, 3}, [][]int{{2}, {2, 2}, {3, 3, 3, 3}}},

		// Special cases
		{"single branch", []interface{}{1, 2, nil, 3, nil, nil, 4}, [][]int{{1}, {2}, {3}, {4}}},
		{"zigzag pattern", []interface{}{1, 2, nil, nil, 3, 4, nil}, [][]int{{1}, {2}, {3}, {4}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the tree from slice representation
			root := createTree(tt.tree)

			// Perform level order traversal
			result := levelOrder(root)

			assert.Equal(t, tt.expected, result,
				"levelOrder(%v) should return %v, got %v",
				tt.tree, tt.expected, result)
		})
	}
}
