package main

import "testing"

func Test_dinicMaxFlow(t *testing.T) {
	tests := []struct {
		name string
		s, t int
		adj  []map[int]int
		want int
	}{
		{
			name: "simple linear flow",
			s:    0,
			t:    2,
			adj: []map[int]int{
				{1: 10},      // 0->1: 10
				{2: 5},       // 1->2: 5
				{},           // 2
			},
			want: 5, // bottleneck at 1->2
		},
		{
			name: "multiple paths",
			s:    0,
			t:    3,
			adj: []map[int]int{
				{1: 10, 2: 10}, // 0->1: 10, 0->2: 10
				{3: 4},         // 1->3: 4
				{3: 9},         // 2->3: 9
				{},             // 3
			},
			want: 13, // path 0->1->3: 4, path 0->2->3: 9
		},
		{
			name: "direct edge",
			s:    0,
			t:    1,
			adj: []map[int]int{
				{1: 100}, // 0->1: 100
				{},       // 1
			},
			want: 100,
		},
		{
			name: "no path",
			s:    0,
			t:    2,
			adj: []map[int]int{
				{1: 10}, // 0->1: 10
				{},      // 1 (no edge to 2)
				{},      // 2
			},
			want: 0,
		},
		{
			name: "complex network",
			s:    0,
			t:    5,
			adj: []map[int]int{
				{1: 16, 2: 13},    // 0
				{2: 10, 3: 12},    // 1
				{1: 4, 4: 14},     // 2
				{2: 9, 5: 20},     // 3
				{3: 7, 5: 4},      // 4
				{},                // 5
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Deep copy adj to avoid modification affecting test data
			adjCopy := make([]map[int]int, len(tt.adj))
			for i := range tt.adj {
				adjCopy[i] = make(map[int]int)
				for k, v := range tt.adj[i] {
					adjCopy[i][k] = v
				}
			}
			got := dinicMaxFlow(tt.s, tt.t, adjCopy)
			if got != tt.want {
				t.Errorf("dinicMaxFlow() = %v, want %v", got, tt.want)
			}
		})
	}
}
