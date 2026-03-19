package main

import "testing"

func TestLongestAlternating(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "single element",
			nums: []int{5},
			want: 1,
		},
		{
			name: "two elements increasing",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "two elements decreasing",
			nums: []int{3, 1},
			want: 2,
		},
		{
			name: "three elements strictly alternating",
			nums: []int{1, 3, 2},
			want: 3,
		},
		{
			name: "three elements increasing",
			nums: []int{1, 2, 3},
			want: 2,
		},
		{
			name: "three elements decreasing",
			nums: []int{3, 2, 1},
			want: 2,
		},
		{
			name: "four elements alternating",
			nums: []int{1, 4, 2, 5},
			want: 4,
		},
		{
			name: "long alternating sequence",
			nums: []int{10, 22, 9, 33, 49, 50, 31, 60},
			want: 4,
		},
		{
			name: "alternating with same differences",
			nums: []int{1, 3, 1, 3, 1},
			want: 5,
		},
		{
			name: "all same elements",
			nums: []int{5, 5, 5, 5},
			want: 1,
		},
		{
			name: "alternating with larger jumps",
			nums: []int{1, 100, 2, 200, 3, 300},
			want: 6,
		},
		{
			name: "alternating then flat",
			nums: []int{1, 5, 2, 2, 2},
			want: 3,
		},
		{
			name: "complex pattern",
			nums: []int{3, 5, 2, 7, 9, 11, 8, 6, 4},
			want: 4,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := longestAlternating(tt.nums)
			if got != tt.want {
				t.Errorf("longestAlternating(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
