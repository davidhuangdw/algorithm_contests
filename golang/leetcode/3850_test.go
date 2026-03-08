package main

import "testing"

func Test_countSequences(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int64
		want int
	}{
		{
			name: "simple k=2",
			nums: []int{2, 1},
			k:    2,
			want: 3, // skip 2, use 2, or 2/1=2
		},
		{
			name: "k=1 equal division",
			nums: []int{2, 2},
			k:    1,
			want: 3, // 2/2=1, and other combinations
		},
		{
			name: "single element equals k",
			nums: []int{3},
			k:    3,
			want: 1, // just use 3 in numerator
		},
		{
			name: "single element not k",
			nums: []int{5},
			k:    3,
			want: 0, // cannot achieve k=3 with just 5
		},
		{
			name: "k=4 with factors",
			nums: []int{2, 2},
			k:    4,
			want: 1, // 2*2/1 = 4
		},
		{
			name: "multiple ways",
			nums: []int{2, 3, 6},
			k:    6,
			want: 2, // 6/1=6, or 2*3/1=6
		},
		{
			name: "k=1 with three elements",
			nums: []int{2, 3, 6},
			k:    1,
			want: 3, // 2/2=1 variations
		},
		{
			name: "impossible ratio",
			nums: []int{2, 3},
			k:    5,
			want: 0, // cannot make 5 from 2 and 3
		},
		{
			name: "k=12 achievable",
			nums: []int{2, 2, 3},
			k:    12,
			want: 1, // 2*2*3/1 = 12
		},
		{
			name: "all ones k=1",
			nums: []int{1, 1, 1},
			k:    1,
			want: 27, // 3^3 combinations all work since 1s are neutral
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSequences(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("countSequences(%v, %v) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
