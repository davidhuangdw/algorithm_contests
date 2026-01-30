package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEulerPath(t *testing.T) {
	// Test case 1: Simple directed Euler path (1→2→3→4)
	// Node 1: out=1, in=0 → odd (out-in)
	// Node 4: out=0, in=1 → odd (out-in)
	adj1 := [][]int{
		nil,
		{2}, // 1
		{3}, // 2
		{4}, // 3
		{},  // 4
	}
	path1 := getDirectedEulerPath(4, adj1)
	assert.NotNil(t, path1)
	assert.Equal(t, 4, len(path1)) // 3 edges + 1 nodes
	assert.Equal(t, []int{1, 2, 3, 4}, path1)

	// Test case 2: Directed Euler circuit (1→2→3→1) - should return nil
	// All nodes have equal in-degree and out-degree → 0 odd nodes
	adj2 := [][]int{
		nil,
		{2}, // 1
		{3}, // 2
		{1}, // 3
	}
	path2 := getDirectedEulerPath(3, adj2)
	assert.Equal(t, []int{1, 2, 3, 1}, path2)

	// Test case 3: Directed graph with more than 2 nodes having odd (out-in) difference - should return nil
	adj3 := [][]int{
		nil,
		{2},    // 1: out=1, in=0 → odd
		{3},    // 2: out=1, in=1 → even
		{4, 5}, // 3: out=2, in=1 → odd
		{},     // 4: out=0, in=1 → odd
		{},     // 5: out=0, in=1 → odd
	}
	path3 := getDirectedEulerPath(5, adj3)
	assert.Nil(t, path3)

	// Test case 4: Disconnected directed graph - should return nil
	adj4 := [][]int{
		nil,
		{2}, // 1 → 2 chain (component 1)
		{},
		{4}, // 3 → 4 chain (component 2)
		{},
	}
	path4 := getDirectedEulerPath(4, adj4)
	assert.Nil(t, path4)

	// Test case 5: Complex directed Euler path (1→2→3→2→4→3→5)
	// Node 1: out=1, in=0 → odd
	// Node 5: out=0, in=1 → odd
	adj5 := [][]int{
		nil,
		{2},    // 1
		{3, 4}, // 2
		{2, 5}, // 3
		{3},    // 4
		{},     // 5
	}
	path5 := getDirectedEulerPath(5, adj5) // n=6 but only nodes 1-5 used
	assert.NotNil(t, path5)
	assert.Equal(t, 7, len(path5)) // 5 edges + 1 nodes

	// Test case 6: Single directed edge (1→2)
	// Node 1: out=1, in=0 → odd
	// Node 2: out=0, in=1 → odd
	adj6 := [][]int{
		nil,
		{2}, // 1
		{},  // 2
	}
	path6 := getDirectedEulerPath(2, adj6)
	assert.NotNil(t, path6)
	assert.Equal(t, []int{1, 2}, path6)

	// Test case 7: Directed chain (1→2→3)
	// Node 1: out=1, in=0 → odd
	// Node 3: out=0, in=1 → odd
	adj7 := [][]int{
		nil,
		{2}, // 1
		{3}, // 2
		{},  // 3
	}
	path7 := getDirectedEulerPath(3, adj7)
	assert.NotNil(t, path7)
	assert.Equal(t, []int{1, 2, 3}, path7)

	// Test case 8: Directed graph with self-loop (u==v)
	// Node 1: out=1, in=1 → even (1→2)
	// Node 2: out=2, in=2 → even (2→3, 2→2, 1→2, 3→2)
	// Node 3: out=1, in=1 → even (3→2)
	adj8 := [][]int{
		nil,
		{2},    // 1
		{3, 2}, // 2 (with self-loop)
		{2},    // 3
	}
	path8 := getDirectedEulerPath(3, adj8)
	assert.NotNil(t, path8)
	assert.Equal(t, 5, len(path8)) // 4 edges (1→2, 2→3, 3→2, 2→2) + 1 nodes

	// Test case 9: Complex directed Euler circuit
	// All nodes have equal in-degree and out-degree
	// 1→2→4→1, 2→3→4→2
	adj9 := [][]int{
		nil,
		{2}, // 1: out=1, in=1
		{3, 4}, // 2: out=2, in=2
		{4}, // 3: out=1, in=1
		{1, 2}, // 4: out=2, in=2
	}
	path9 := getDirectedEulerPath(4, adj9)
	assert.NotNil(t, path9)
	assert.Equal(t, 7, len(path9)) // 6 edges + 1 nodes
	// Check it's a proper circuit (starts and ends at same node)
	assert.Equal(t, path9[0], path9[len(path9)-1])
}