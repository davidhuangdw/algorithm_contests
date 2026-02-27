package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_vertexBiconnected(t *testing.T) {
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
			want: [][]int{{1, 2}},
		},
		{
			name: "triangle - single biconnected component",
			adj:  [][]int{nil, []int{2, 3}, []int{1, 3}, []int{1, 2}},
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "chain with cut vertices - three components",
			adj:  [][]int{nil, []int{2}, []int{1, 3}, []int{2, 4}, []int{3}},
			want: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			name: "star graph - multiple components sharing center",
			adj:  [][]int{nil, []int{2, 3, 4}, []int{1}, []int{1}, []int{1}},
			want: [][]int{{1, 4}, {1, 3}, {1, 2}},
		},
		{
			name: "two triangles connected by bridge",
			adj:  [][]int{nil, []int{2, 3}, []int{1, 3}, []int{1, 2, 4}, []int{3, 5, 6}, []int{4, 6}, []int{4, 5}},
			want: [][]int{{4, 5, 6}, {3, 4}, {1, 2, 3}},
		},
		{
			name: "disconnected components",
			adj:  [][]int{nil, []int{2}, []int{1}, []int{4}, []int{3}},
			want: [][]int{{1, 2}, {3, 4}},
		},
		{
			name: "single edge",
			adj:  [][]int{nil, []int{2}, []int{1}},
			want: [][]int{{1, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := vertexBiconnected(tt.adj)

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
