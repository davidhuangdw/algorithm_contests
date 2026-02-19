package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasNegCycle(t *testing.T) {
	// Test with a simple graph with no negative cycle
	adj := make([][][]int, 4)
	adj[1] = append(adj[1], []int{2, 5})
	adj[2] = append(adj[2], []int{3, 3})
	adj[3] = append(adj[3], []int{1, 2})
	assert.False(t, hasNegCycle(adj), "Graph with positive weights should have no negative cycle")

	// Test with a simple graph with a negative cycle
	adj = make([][][]int, 4)
	adj[1] = append(adj[1], []int{2, 1})
	adj[2] = append(adj[2], []int{3, -2})
	adj[3] = append(adj[3], []int{1, -2})
	assert.True(t, hasNegCycle(adj), "Graph with negative cycle should return true")

	// Test with a graph with a single negative edge but no cycle
	adj = make([][][]int, 4)
	adj[1] = append(adj[1], []int{2, -5})
	adj[2] = append(adj[2], []int{3, 3})
	assert.False(t, hasNegCycle(adj), "Graph with negative edge but no cycle should return false")

	// Test with a graph with a self-loop negative edge
	adj = make([][][]int, 3)
	adj[1] = append(adj[1], []int{1, -1})
	assert.True(t, hasNegCycle(adj), "Graph with self-loop negative edge should return true")

	// Test with an empty graph
	adj = make([][][]int, 2)
	assert.False(t, hasNegCycle(adj), "Empty graph should have no negative cycle")

	// Test with a graph with multiple edges but no negative cycle
	adj = make([][][]int, 5)
	adj[1] = append(adj[1], []int{2, 1})
	adj[1] = append(adj[1], []int{3, 4})
	adj[2] = append(adj[2], []int{3, 2})
	adj[3] = append(adj[3], []int{4, 1})
	assert.False(t, hasNegCycle(adj), "Graph with multiple edges but no negative cycle should return false")

	// Test with a simple negative cycle (1->2->3->1 with total weight -1)
	adj = make([][][]int, 4)
	adj[1] = append(adj[1], []int{2, 1})
	adj[2] = append(adj[2], []int{3, 1})
	adj[3] = append(adj[3], []int{1, -3})
	assert.True(t, hasNegCycle(adj), "Simple negative cycle should return true")
}
