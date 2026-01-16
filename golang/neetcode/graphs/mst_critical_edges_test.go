package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCriticalAndPseudoCriticalEdges(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		edges       [][]int
		expected    [][]int
		description string
	}{
		{
			name: "LeetCode Example 1",
			n:    5,
			edges: [][]int{
				{0, 1, 1}, {1, 2, 1}, {2, 3, 2},
				{0, 3, 2}, {0, 4, 3}, {3, 4, 3},
				{1, 4, 6},
			},
			expected:    [][]int{{0, 1}, {2, 3, 4, 5}},
			description: "Example with both critical and pseudo-critical edges",
		},
		{
			name:        "Single edge",
			n:           2,
			edges:       [][]int{{0, 1, 1}},
			expected:    [][]int{{0}, {}},
			description: "Only one edge in the graph, which is critical",
		},
		{
			name: "Two edges with same weight",
			n:    3,
			edges: [][]int{
				{0, 1, 1}, {1, 2, 1}, {0, 2, 1},
			},
			expected:    [][]int{{}, {0, 1, 2}},
			description: "All edges have same weight, all are pseudo-critical",
		},
		{
			name: "No pseudo-critical edges",
			n:    4,
			edges: [][]int{
				{0, 1, 1}, {1, 2, 2}, {2, 3, 3}, {0, 3, 4},
			},
			expected:    [][]int{{0, 1, 2}, {}},
			description: "MST with unique edges of increasing weights",
		},
		{
			name:        "Disconnected graph",
			n:           3,
			edges:       [][]int{{0, 1, 1}},
			expected:    [][]int{{0}, {}},
			description: "Graph with two components, the single edge is critical for its component",
		},
		{
			name: "Multiple pseudo-critical edges in same weight class",
			n:    4,
			edges: [][]int{
				{0, 1, 1}, {0, 2, 1}, {1, 2, 1},
				{1, 3, 2}, {2, 3, 2},
			},
			expected:    [][]int{{}, {0, 1, 2, 3, 4}},
			description: "Two weight classes with multiple edges each",
		},
		{
			name: "Complete graph with unique MST",
			n:    4,
			edges: [][]int{
				{0, 1, 1}, {0, 2, 3}, {0, 3, 4},
				{1, 2, 2}, {1, 3, 5}, {2, 3, 6},
			},
			expected:    [][]int{{0, 2, 3}, {}},
			description: "Complete graph where MST is unique",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findCriticalAndPseudoCriticalEdges(tt.n, tt.edges)

			assert.ElementsMatch(t, result[0], tt.expected[0], "%s: Expected %v, got %v", tt.description, tt.expected, result)
			assert.ElementsMatch(t, result[1], tt.expected[1], "%s: Expected %v, got %v", tt.description, tt.expected, result)
		})
	}
}
