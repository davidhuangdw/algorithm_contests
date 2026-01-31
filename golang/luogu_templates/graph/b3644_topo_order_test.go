package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopoOrder(t *testing.T) {
	// Test case 1: Simple DAG with clear topological order
	// Nodes: 1->2, 1->3, 2->4, 3->4
	t.Run("Simple DAG", func(t *testing.T) {
		n := 4
		adj := make([][]int, n+1)
		adj[1] = []int{2, 3}
		adj[2] = []int{4}
		adj[3] = []int{4}
		
		result := topoOrder(n, adj)
		// The topological order must have 1 before 2 and 3, and 2 and 3 before 4
		assert.Len(t, result, 4)
		assert.Equal(t, 1, result[0]) // 1 must be first
		
		// Check positions: 2 < 4 and 3 < 4
		pos2 := -1
		pos3 := -1
		pos4 := -1
		for i, v := range result {
			switch v {
			case 2:
				pos2 = i
			case 3:
				pos3 = i
			case 4:
				pos4 = i
			}
		}
		assert.True(t, pos2 < pos4)
		assert.True(t, pos3 < pos4)
	})

	// Test case 2: Single node
	t.Run("Single Node", func(t *testing.T) {
		n := 1
		adj := make([][]int, n+1)
		result := topoOrder(n, adj)
		assert.Equal(t, []int{1}, result)
	})

	// Test case 3: Chain-like DAG
	// 1->2->3->4->5
	t.Run("Chain DAG", func(t *testing.T) {
		n := 5
		adj := make([][]int, n+1)
		for i := 1; i < n; i++ {
			adj[i] = []int{i + 1}
		}
		
		result := topoOrder(n, adj)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
	})

	// Test case 4: DAG with multiple roots
	// Roots: 1, 2
	// 1->3, 2->3, 2->4
	t.Run("Multiple Roots", func(t *testing.T) {
		n := 4
		adj := make([][]int, n+1)
		adj[1] = []int{3}
		adj[2] = []int{3, 4}
		
		result := topoOrder(n, adj)
		assert.Len(t, result, 4)
		
		// 3 must come after 1 and 2, 4 must come after 2
		pos1 := -1
		pos2 := -1
		pos3 := -1
		pos4 := -1
		for i, v := range result {
			switch v {
			case 1:
				pos1 = i
			case 2:
				pos2 = i
			case 3:
				pos3 = i
			case 4:
				pos4 = i
			}
		}
		assert.True(t, pos1 < pos3)
		assert.True(t, pos2 < pos3)
		assert.True(t, pos2 < pos4)
	})

	// Test case 5: Complex DAG
	// 1->2, 1->3, 2->4, 3->4, 3->5, 4->6, 5->6
	t.Run("Complex DAG", func(t *testing.T) {
		n := 6
		adj := make([][]int, n+1)
		adj[1] = []int{2, 3}
		adj[2] = []int{4}
		adj[3] = []int{4, 5}
		adj[4] = []int{6}
		adj[5] = []int{6}
		
		result := topoOrder(n, adj)
		assert.Len(t, result, 6)
		
		// Check dependencies
		assert.Equal(t, 1, result[0]) // 1 must be first
		
		pos2 := -1
		pos3 := -1
		pos4 := -1
		pos5 := -1
		pos6 := -1
		for i, v := range result {
			switch v {
			case 2:
				pos2 = i
			case 3:
				pos3 = i
			case 4:
				pos4 = i
			case 5:
				pos5 = i
			case 6:
				pos6 = i
			}
		}
		
		// All dependencies must be satisfied
		assert.True(t, pos2 < pos4)
		assert.True(t, pos3 < pos4)
		assert.True(t, pos3 < pos5)
		assert.True(t, pos4 < pos6)
		assert.True(t, pos5 < pos6)
	})
}