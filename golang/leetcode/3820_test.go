package main

import (
	"testing"
)

func Test_specialNodes(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		edges  [][]int
		x      int
		y      int
		z      int
		want   int
	}{
		{
			name:  "single node",
			n:     1,
			edges: [][]int{},
			x:     0,
			y:     0,
			z:     0,
			want:  1,
		},
		{
			name:  "two nodes line",
			n:     2,
			edges: [][]int{{0, 1}},
			x:     0,
			y:     1,
			z:     0,
			want:  1,
		},
		{
			name:  "three nodes star",
			n:     3,
			edges: [][]int{{0, 1}, {0, 2}},
			x:     0,
			y:     1,
			z:     2,
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := specialNodes(tt.n, tt.edges, tt.x, tt.y, tt.z)
			if got != tt.want {
				t.Errorf("specialNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}