package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxDirectedPath(t *testing.T) {
	// Test case 1: Single node
	adj1 := [][]int{nil, []int{}}
	wei1 := []int{0, 5}
	assert.Equal(t, 5, getMaxDirectedPath(1, wei1, adj1))

	// Test case 2: Linear chain (no cycles)
	// 1 -> 2 -> 3
	adj2 := [][]int{nil, []int{2}, []int{3}, []int{}}
	wei2 := []int{0, 1, 2, 3}
	assert.Equal(t, 6, getMaxDirectedPath(3, wei2, adj2))

	// Test case 3: Graph with a cycle
	// 1 -> 2 -> 3 -> 1
	adj3 := [][]int{nil, []int{2}, []int{3}, []int{1}}
	wei3 := []int{0, 1, 2, 3}
	assert.Equal(t, 6, getMaxDirectedPath(3, wei3, adj3))

	// Test case 4: Two separate SCCs with edge between them
	// SCC1: 1 <-> 2, SCC2: 3 <-> 4, Edge: 2 -> 3
	adj4 := [][]int{nil, []int{2}, []int{1, 3}, []int{4}, []int{3}}
	wei4 := []int{0, 1, 2, 3, 4}
	assert.Equal(t, 10, getMaxDirectedPath(4, wei4, adj4))

	// Test case 5: No edges (each node is own SCC)
	adj5 := [][]int{nil, []int{}, []int{}, []int{}}
	wei5 := []int{0, 3, 1, 4}
	assert.Equal(t, 4, getMaxDirectedPath(3, wei5, adj5))

	// Test case 6: Diamond-shaped graph
	// 1 -> 2, 1 -> 3, 2 -> 4, 3 -> 4
	adj6 := [][]int{nil, []int{2, 3}, []int{4}, []int{4}, []int{}}
	wei6 := []int{0, 1, 2, 3, 4}
	assert.Equal(t, 8, getMaxDirectedPath(4, wei6, adj6))

	// Test case 7: Complex graph with multiple paths
	// 1 -> 2 -> 5, 1 -> 3 -> 5, 2 -> 4 -> 5, 3 -> 4
	adj7 := [][]int{nil, []int{2, 3}, []int{5, 4}, []int{5, 4}, []int{5}, []int{}}
	wei7 := []int{0, 1, 2, 3, 4, 5}
	assert.Equal(t, 13, getMaxDirectedPath(5, wei7, adj7))

	// Test case 8: Graph with self-loop (should be ignored)
	adj8 := [][]int{nil, []int{1, 2}, []int{}}
	wei8 := []int{0, 3, 5}
	assert.Equal(t, 8, getMaxDirectedPath(2, wei8, adj8))
}
