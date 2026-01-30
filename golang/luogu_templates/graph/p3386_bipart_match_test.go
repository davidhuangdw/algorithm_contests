package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBipartMaxMatch(t *testing.T) {
	tests := []struct {
		name string
		n    int
		m    int
		adj  [][]int
		want int
	}{
		{
			name: "empty graph - n=0",
			n:    0,
			m:    5,
			adj:  [][]int{},
			want: 0,
		},
		{
			name: "empty graph - m=0",
			n:    5,
			m:    0,
			adj:  make([][]int, 5),
			want: 0,
		},
		{
			name: "single edge",
			n:    1,
			m:    1,
			adj:  [][]int{{0}},
			want: 1,
		},
		{
			name: "perfect matching",
			n:    3,
			m:    3,
			adj:  [][]int{{0}, {1}, {2}},
			want: 3,
		},
		{
			name: "multiple edges, one matching per left node",
			n:    2,
			m:    3,
			adj:  [][]int{{0, 1, 2}, {1, 2}},
			want: 2,
		},
		{
			name: "matching with shared right nodes",
			n:    3,
			m:    2,
			adj:  [][]int{{0}, {0, 1}, {1}},
			want: 2,
		},
		{
			name: "complex bipartite graph",
			n:    4,
			m:    4,
			adj:  [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}},
			want: 4, // Perfect matching possible
		},
		{
			name: "disconnected components",
			n:    4,
			m:    5,
			adj:  [][]int{{0}, {}, {2, 3}, {4}},
			want: 3,
		},
		{
			name: "multiple edges to same right node",
			n:    3,
			m:    2,
			adj:  [][]int{{0, 1}, {0}, {0}},
			want: 2,
		},
		{
			name: "star-shaped graph",
			n:    5,
			m:    1,
			adj:  [][]int{{0}, {0}, {0}, {0}, {0}},
			want: 1,
		},
		{
			name: "bug case - requires backtracking with different edges",
			n:    3,
			m:    3,
			adj: [][]int{
				{0, 1},    // Left 0 can connect to right 0, 1
				{1, 2},    // Left 1 can connect to right 1, 2
				{0, 1, 2}, // Left 2 can connect to right 0, 1, 2
			},
			want: 3, // Perfect matching possible
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bipartMaxMatch(tt.n, tt.m, tt.adj)
			assert.Equal(t, tt.want, got, "bipartMaxMatch(%d, %d, adj) = %d, want %d", tt.n, tt.m, got, tt.want)
		})
	}
}
