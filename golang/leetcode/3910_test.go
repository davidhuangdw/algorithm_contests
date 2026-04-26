package main

import "testing"

func Test_evenSumSubgraphs(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		edges [][]int
		want  int
	}{
		{
			name:  "Example 1",
			nums:  []int{1, 0, 1},
			edges: [][]int{{0, 1}, {1, 2}},
			want:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evenSumSubgraphs(tt.nums, tt.edges); got != tt.want {
				t.Errorf("evenSumSubgraphs() = %v, want %v", got, tt.want)
			}
		})
	}
}
