package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_edgeBiconnected(t *testing.T) {
	tests := []struct {
		name string
		adj  [][]int
		want [][]int
	}{
		{
			name: "single vertex",
			adj:  [][]int{nil, []int{}},
			want: [][]int{{1}},
		},
		{
			name: "two connected vertices",
			adj:  [][]int{nil, []int{2}, []int{1}},
			want: [][]int{{2}, {1}},
		},
		{
			name: "triangle - single component",
			adj:  [][]int{nil, []int{2, 3}, []int{1, 3}, []int{1, 2}},
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "chain with bridge - separate components",
			adj:  [][]int{nil, []int{2}, []int{1, 3}, []int{2, 4}, []int{3}},
			want: [][]int{{4}, {3}, {2}, {1}},
		},
		{
			name: "two triangles connected by bridge",
			adj:  [][]int{nil, []int{2, 3}, []int{1, 3}, []int{1, 2, 4}, []int{3, 5, 6}, []int{4, 6}, []int{4, 5}},
			want: [][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			name: "square - single component",
			adj:  [][]int{nil, []int{2, 4}, []int{1, 3}, []int{2, 4}, []int{1, 3}},
			want: [][]int{{1, 2, 3, 4}},
		},
		{
			name: "disconnected components",
			adj:  [][]int{nil, []int{2}, []int{1}, []int{4}, []int{3}},
			want: [][]int{{2}, {1}, {4}, {3}},
		},
		{
			name: "star graph - all bridges",
			adj:  [][]int{nil, []int{2, 3, 4}, []int{1}, []int{1}, []int{1}},
			want: [][]int{{4}, {3}, {2}, {1}},
		},
		{
			name: "complex graph with cycle and bridges",
			adj:  [][]int{nil, []int{2}, []int{1, 3, 4}, []int{2, 4}, []int{2, 3, 5}, []int{4}},
			want: [][]int{{5}, {2, 3, 4}, {1}},
		},
		{
			name: "two nodes with double edge - single component",
			adj:  [][]int{nil, []int{2, 2}, []int{1, 1}},
			want: [][]int{{1, 2}},
		},
		{
			name: "multi-edges between nodes",
			adj:  [][]int{nil, []int{2, 2, 3}, []int{1, 1, 3}, []int{1, 2}},
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "chain with double edge in middle",
			adj:  [][]int{nil, []int{2}, []int{1, 3, 3}, []int{2, 2, 4}, []int{3}},
			want: [][]int{{4}, {2, 3}, {1}},
		},
		{
			name: "mixed single and double edges",
			adj:  [][]int{nil, []int{2, 2}, []int{1, 1, 3}, []int{2, 4}, []int{3}},
			want: [][]int{{4}, {3}, {1, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := edgeBiconnected(tt.adj)

			// Sort for comparison since order may vary
			for i := range got {
				sort.Ints(got[i])
			}
			for i := range tt.want {
				sort.Ints(tt.want[i])
			}
			sort.Slice(got, func(i, j int) bool {
				if len(got[i]) != len(got[j]) {
					return len(got[i]) < len(got[j])
				}
				for k := 0; k < len(got[i]); k++ {
					if got[i][k] != got[j][k] {
						return got[i][k] < got[j][k]
					}
				}
				return false
			})
			sort.Slice(tt.want, func(i, j int) bool {
				if len(tt.want[i]) != len(tt.want[j]) {
					return len(tt.want[i]) < len(tt.want[j])
				}
				for k := 0; k < len(tt.want[i]); k++ {
					if tt.want[i][k] != tt.want[j][k] {
						return tt.want[i][k] < tt.want[j][k]
					}
				}
				return false
			})

			assert.Equal(t, tt.want, got)
		})
	}
}
