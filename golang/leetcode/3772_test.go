package main

import (
	"reflect"
	"testing"
)

func TestMaxSubgraphScore(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		good     []int
		expected []int
	}{
		{
			name:     "Line graph with alternating weights",
			n:        3,
			edges:    [][]int{{0, 1}, {1, 2}},
			good:     []int{1, 0, 1},
			expected: []int{1, 1, 1},
		},
		{
			name:     "Star graph with negative center",
			n:        4,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}},
			good:     []int{0, 1, 1, 1},
			expected: []int{2, 2, 2, 2},
		},
		{
			name:     "More complex tree",
			n:        5,
			edges:    [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 4}},
			good:     []int{1, 0, 1, 1, 0},
			expected: []int{2, 2, 2, 2, 1},
		},
		{
			name:     "Single node good",
			n:        1,
			edges:    [][]int{},
			good:     []int{1},
			expected: []int{1},
		},
		{
			name:     "Single node bad",
			n:        1,
			edges:    [][]int{},
			good:     []int{0},
			expected: []int{-1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubgraphScore(tt.n, tt.edges, tt.good); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("maxSubgraphScore() = %v, want %v", got, tt.expected)
			}
		})
	}
}
