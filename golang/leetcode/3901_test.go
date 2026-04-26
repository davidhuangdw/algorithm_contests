package main

import "testing"

func Test_countGoodSubseq(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		p       int
		queries [][]int
		want    int
	}{
		{
			name:    "Example 1",
			nums:    []int{2, 4, 6},
			p:       2,
			queries: [][]int{{0, 1}},
			want:    1,
		},
		{
			name:    "Example 2",
			nums:    []int{1, 2, 3},
			p:       1,
			queries: [][]int{{0, 2}},
			want:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodSubseq(tt.nums, tt.p, tt.queries); got != tt.want {
				t.Errorf("countGoodSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
