package main

import "testing"

func TestInteractionCosts(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		group    []int
		expected int64
	}{
		{
			name:     "Simple path, all same group",
			n:        3,
			edges:    [][]int{{0, 1}, {1, 2}},
			group:    []int{0, 0, 0},
			expected: 4,
		},
		{
			name:     "Simple path, mixed groups",
			n:        3,
			edges:    [][]int{{0, 1}, {1, 2}},
			group:    []int{0, 1, 0},
			expected: 2,
		},
		{
			name:     "Star graph, all same group",
			n:        4,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}},
			group:    []int{0, 0, 0, 0},
			expected: 9,
		},
		{
			name:     "Star graph, mixed groups",
			n:        4,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}},
			group:    []int{0, 1, 0, 1},
			expected: 3, // Corrected: group 0 dist is 1, group 1 dist is 2. Total = 3.
		},
		{
			name:     "Disconnected nodes (not a tree, but handled by code)",
			n:        2,
			edges:    [][]int{},
			group:    []int{0, 0},
			expected: 0,
		},
		{
			name:     "Two nodes, one edge, different groups",
			n:        2,
			edges:    [][]int{{0, 1}},
			group:    []int{0, 1},
			expected: 0,
		},
		{
			name:     "Two nodes, one edge, same group",
			n:        2,
			edges:    [][]int{{0, 1}},
			group:    []int{0, 0},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := interactionCosts(tt.n, tt.edges, tt.group); got != tt.expected {
				t.Errorf("interactionCosts() = %v, want %v", got, tt.expected)
			}
		})
	}
}
