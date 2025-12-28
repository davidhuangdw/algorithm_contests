package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidTree(t *testing.T) {
	t.Run("single node no edges", func(t *testing.T) {
		result := validTree(1, [][]int{})
		assert.True(t, result)
	})

	t.Run("two nodes connected", func(t *testing.T) {
		edges := [][]int{{0, 1}}
		result := validTree(2, edges)
		assert.True(t, result)
	})

	t.Run("three nodes in line", func(t *testing.T) {
		edges := [][]int{{0, 1}, {1, 2}}
		result := validTree(3, edges)
		assert.True(t, result)
	})

	t.Run("three nodes star pattern", func(t *testing.T) {
		edges := [][]int{{0, 1}, {0, 2}}
		result := validTree(3, edges)
		assert.True(t, result)
	})

	t.Run("four nodes tree", func(t *testing.T) {
		edges := [][]int{{0, 1}, {1, 2}, {2, 3}}
		result := validTree(4, edges)
		assert.True(t, result)
	})

	t.Run("cycle in three nodes", func(t *testing.T) {
		edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
		result := validTree(3, edges)
		assert.False(t, result)
	})

	t.Run("cycle in four nodes", func(t *testing.T) {
		edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}
		result := validTree(4, edges)
		assert.False(t, result)
	})

	t.Run("disconnected two components", func(t *testing.T) {
		edges := [][]int{{0, 1}, {2, 3}} // Two separate edges
		result := validTree(4, edges)
		assert.False(t, result)
	})

	t.Run("self-loop", func(t *testing.T) {
		edges := [][]int{{0, 0}} // Self-loop
		result := validTree(1, edges)
		assert.False(t, result)
	})

	t.Run("too many edges for nodes", func(t *testing.T) {
		edges := [][]int{{0, 1}, {0, 2}, {1, 2}} // 3 edges for 3 nodes (should be 2)
		result := validTree(3, edges)
		assert.False(t, result)
	})

	t.Run("too few edges for nodes", func(t *testing.T) {
		edges := [][]int{{0, 1}} // Only 1 edge for 3 nodes
		result := validTree(3, edges)
		assert.False(t, result)
	})

	t.Run("complex tree without cycle", func(t *testing.T) {
		edges := [][]int{
		{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6},
	}
		result := validTree(7, edges)
		assert.True(t, result)
	})

	t.Run("complex graph with cycle", func(t *testing.T) {
		edges := [][]int{
		{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}, {4, 6}, // Creates cycle 0-2-6-4-1-0
	}
		result := validTree(7, edges)
		assert.False(t, result)
	})

	t.Run("single node with self-loop", func(t *testing.T) {
		edges := [][]int{{0, 0}}
		result := validTree(1, edges)
		assert.False(t, result)
	})

	t.Run("empty graph with no nodes", func(t *testing.T) {
		result := validTree(0, [][]int{})
		assert.True(t, result)
	})

	t.Run("large tree", func(t *testing.T) {
		edges := make([][]int, 99)
		for i := 0; i < 99; i++ {
			edges[i] = []int{i, i + 1}
		}
		result := validTree(100, edges)
		assert.True(t, result)
	})
}