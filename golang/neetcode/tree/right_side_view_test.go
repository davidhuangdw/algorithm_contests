package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name     string
		tree     []interface{}
		expected []int
	}{
		// Basic cases
		{"empty tree", []interface{}{}, nil},
		{"single node", []interface{}{1}, []int{1}},

		// Classic cases
		{"balanced tree", []interface{}{1, 2, 3, nil, 5, nil, 4}, []int{1, 3, 4}},
		{"LeetCode example", []interface{}{1, 2, 3, nil, 5, nil, 4}, []int{1, 3, 4}},
		{"three level tree", []interface{}{1, 2, 3, 4, nil, nil, 5}, []int{1, 3, 5}},

		// Edge cases
		{"left skewed", []interface{}{1, 2, nil, 3, nil, 4}, []int{1, 2, 3, 4}},
		{"right skewed", []interface{}{1, nil, 2, nil, 3}, []int{1, 2, 3}},
		{"incomplete levels", []interface{}{1, 2, 3, 4}, []int{1, 3, 4}},

		// Complex cases
		{"large tree", []interface{}{1, 2, 3, 4, 5, nil, 6, nil, nil, 7, nil, nil, nil, nil, 8}, []int{1, 3, 6, 7, 8}},
		{"negative values", []interface{}{-1, -2, -3, -4, -5}, []int{-1, -3, -5}},
		{"mixed values", []interface{}{5, -3, 8, 0, -1, 6, 9}, []int{5, 8, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.tree)
			result := rightSideView(root)
			assert.Equal(t, tt.expected, result)
		})
	}
}
