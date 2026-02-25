package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPathJohnson(t *testing.T) {
	// Test case 1: Simple graph with no negative cycles
	t.Run("SimpleGraph", func(t *testing.T) {
		// Graph: 0 -> 1 (1), 0 -> 2 (4), 1 -> 2 (2), 1 -> 3 (5), 2 -> 3 (1)
		adj := [][][2]int{
			{{1, 1}, {2, 4}},
			{{2, 2}, {3, 5}},
			{{3, 1}},
			{},
		}

		expected := [][]int{
			{0, 1, 3, 4},       // Distances from 0
			{1e9, 0, 2, 3},     // Distances from 1
			{1e9, 1e9, 0, 1},   // Distances from 2
			{1e9, 1e9, 1e9, 0}, // Distances from 3
		}

		result := shortestPathJohnson(adj)
		assert.NotNil(t, result, "Result should not be nil")
		assert.Equal(t, expected, result, "Shortest paths should match expected values")
	})

	// Test case 2: Graph with negative edges but no negative cycles
	t.Run("GraphWithNegativeEdges", func(t *testing.T) {
		// Graph: 0 -> 1 (2), 0 -> 2 (5), 1 -> 2 (-3), 1 -> 3 (1), 2 -> 3 (2)
		adj := [][][2]int{
			{{1, 2}, {2, 5}},
			{{2, -3}, {3, 1}},
			{{3, 2}},
			{},
		}

		expected := [][]int{
			{0, 2, -1, 1},      // Distances from 0
			{1e9, 0, -3, -1},   // Distances from 1
			{1e9, 1e9, 0, 2},   // Distances from 2
			{1e9, 1e9, 1e9, 0}, // Distances from 3
		}

		result := shortestPathJohnson(adj)
		assert.NotNil(t, result, "Result should not be nil")
		assert.Equal(t, expected, result, "Shortest paths should match expected values with negative edges")
	})

	// Test case 3: Graph with negative cycle
	t.Run("GraphWithNegativeCycle", func(t *testing.T) {
		// Graph: 0 -> 1 (1), 1 -> 2 (2), 2 -> 3 (3), 3 -> 1 (-7)
		adj := [][][2]int{
			{{1, 1}},
			{{2, 2}},
			{{3, 3}},
			{{1, -7}},
		}

		result := shortestPathJohnson(adj)
		assert.Nil(t, result, "Result should be nil for graph with negative cycle")
	})

	// Test case 4: Single node graph
	t.Run("SingleNodeGraph", func(t *testing.T) {
		// Adjacency list with one node (no outgoing edges)
		adj := [][][2]int{{}}

		expected := [][]int{{0}}

		result := shortestPathJohnson(adj)
		assert.NotNil(t, result, "Result should not be nil")
		assert.Equal(t, expected, result, "Single node should have distance 0 to itself")
	})

	// Test case 5: Disconnected graph
	t.Run("DisconnectedGraph", func(t *testing.T) {
		// Graph: 0 connected to 1, 2 connected to 3, no edges between components
		adj := [][][2]int{
			{{1, 5}},
			{},
			{{3, 10}},
			{},
		}

		expected := [][]int{
			{0, 5, 1e9, 1e9},   // Distances from 0
			{1e9, 0, 1e9, 1e9}, // Distances from 1
			{1e9, 1e9, 0, 10},  // Distances from 2
			{1e9, 1e9, 1e9, 0}, // Distances from 3
		}

		result := shortestPathJohnson(adj)
		assert.NotNil(t, result, "Result should not be nil")
		assert.Equal(t, expected, result, "Disconnected graph should have correct distances")
	})
}
