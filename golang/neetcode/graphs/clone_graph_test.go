package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloneGraph(t *testing.T) {
	t.Run("nil node", func(t *testing.T) {
		result := cloneGraph(nil)
		assert.Nil(t, result)
	})

	t.Run("single node", func(t *testing.T) {
		node := &Node{Val: 1}
		clone := cloneGraph(node)
		
		assert.NotNil(t, clone)
		assert.Equal(t, 1, clone.Val)
		assert.Empty(t, clone.Neighbors)
		assert.NotSame(t, node, clone) // Ensure it's a different object
	})

	t.Run("two connected nodes", func(t *testing.T) {
		node1 := &Node{Val: 1}
		node2 := &Node{Val: 2}
		node1.Neighbors = []*Node{node2}
		node2.Neighbors = []*Node{node1}

		clone := cloneGraph(node1)
		
		assert.NotNil(t, clone)
		assert.Equal(t, 1, clone.Val)
		assert.Len(t, clone.Neighbors, 1)
		assert.Equal(t, 2, clone.Neighbors[0].Val)
		assert.Len(t, clone.Neighbors[0].Neighbors, 1)
		assert.Equal(t, 1, clone.Neighbors[0].Neighbors[0].Val)
		
		// Ensure original graph is not modified
		assert.Same(t, node1, node1.Neighbors[0].Neighbors[0])
		assert.NotSame(t, node1, clone.Neighbors[0].Neighbors[0])
	})

	t.Run("complex graph with cycle", func(t *testing.T) {
		node1 := &Node{Val: 1}
		node2 := &Node{Val: 2}
		node3 := &Node{Val: 3}
		node4 := &Node{Val: 4}
		
		node1.Neighbors = []*Node{node2, node4}
		node2.Neighbors = []*Node{node1, node3}
		node3.Neighbors = []*Node{node2, node4}
		node4.Neighbors = []*Node{node1, node3}

		clone := cloneGraph(node1)
		
		// Verify structure
		assert.Equal(t, 1, clone.Val)
		assert.Len(t, clone.Neighbors, 2)
		
		// Find neighbors with values 2 and 4
		var clone2, clone4 *Node
		for _, neighbor := range clone.Neighbors {
			if neighbor.Val == 2 {
				clone2 = neighbor
			} else if neighbor.Val == 4 {
				clone4 = neighbor
			}
		}
		
		assert.NotNil(t, clone2)
		assert.NotNil(t, clone4)
		assert.Len(t, clone2.Neighbors, 2)
		assert.Len(t, clone4.Neighbors, 2)
		
		// Verify cycle is maintained
		var clone3 *Node
		for _, neighbor := range clone2.Neighbors {
			if neighbor.Val == 3 {
				clone3 = neighbor
			}
		}
		assert.NotNil(t, clone3)
		assert.Equal(t, 3, clone3.Val)
		
		// Check that clone3 connects back to clone2 and clone4
		foundClone2, foundClone4 := false, false
		for _, neighbor := range clone3.Neighbors {
			if neighbor.Val == 2 {
				foundClone2 = true
			} else if neighbor.Val == 4 {
				foundClone4 = true
			}
		}
		assert.True(t, foundClone2)
		assert.True(t, foundClone4)
	})

	t.Run("node with self-loop", func(t *testing.T) {
		node := &Node{Val: 1}
		node.Neighbors = []*Node{node} // Self-loop

		clone := cloneGraph(node)
		
		assert.NotNil(t, clone)
		assert.Equal(t, 1, clone.Val)
		assert.Len(t, clone.Neighbors, 1)
		assert.Same(t, clone, clone.Neighbors[0]) // Should point to itself
	})

	t.Run("disconnected nodes", func(t *testing.T) {
		node1 := &Node{Val: 1}
		node2 := &Node{Val: 2}
		// node1 and node2 are not connected

		clone1 := cloneGraph(node1)
		clone2 := cloneGraph(node2)
		
		assert.Equal(t, 1, clone1.Val)
		assert.Empty(t, clone1.Neighbors)
		assert.Equal(t, 2, clone2.Val)
		assert.Empty(t, clone2.Neighbors)
	})
}