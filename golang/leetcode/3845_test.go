package main

import "testing"

func Test_maxXor(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "example 1",
			nums: []int{1, 5, 3},
			k:    2,
			want: 6,
		},
		{
			name: "example 2",
			nums: []int{2, 6, 7},
			k:    1,
			want: 7,
		},
		{
			name: "single element",
			nums: []int{5},
			k:    0,
			want: 5,
		},
		{
			name: "all same",
			nums: []int{3, 3, 3},
			k:    0,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxXor(tt.nums, tt.k); got != tt.want {
				t.Errorf("maxXor(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
