package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodec(t *testing.T) {
	tests := []struct {
		name string
		tree []interface{}
	}{
		// Basic cases
		{"empty tree", []interface{}{}},
		{"single node", []interface{}{1}},
		{"single node negative", []interface{}{-5}},

		// Classic cases
		{"balanced tree", []interface{}{1, 2, 3, 4, 5, 6, 7}},
		{"LeetCode example", []interface{}{1, 2, 3, nil, nil, 4, 5}},
		{"three nodes", []interface{}{1, 2, 3}},

		// Edge cases
		{"left skewed", []interface{}{1, 2, nil, 3, nil, nil, nil, 4}},
		{"right skewed", []interface{}{1, nil, 2, nil, nil, nil, 3}},
		{"incomplete levels", []interface{}{1, 2, 3, 4, nil, nil, nil}},

		// Complex cases
		{"large tree", []interface{}{5, 3, 7, 1, 4, 6, 8, nil, 2, nil, nil, nil, nil, nil, 9}},
		{"negative values", []interface{}{-1, -2, -3, -4, -5}},
		{"mixed values", []interface{}{10, -5, 15, -3, 7, 12, 20}},
		{"duplicate values", []interface{}{2, 2, 2, 3, 3, 3, 3}},
		{"zero values", []interface{}{0, -1, 1, -2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create original tree
			originalTree := createTree(tt.tree)

			// Initialize codec
			codec := Constructor()

			// Serialize the tree
			serialized := codec.serialize(originalTree)

			// Deserialize back to tree
			deserializedTree := codec.deserialize(serialized)

			// Convert both trees to slice representation for comparison
			originalSlice := treeToSlice(originalTree)
			deserializedSlice := treeToSlice(deserializedTree)

			// Assert that the deserialized tree matches the original
			assert.Equal(t, originalSlice, deserializedSlice,
				"Tree should be identical after serialization and deserialization")
		})
	}
}
