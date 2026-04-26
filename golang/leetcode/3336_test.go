package main

import "testing"

func Test_subsequencePairCount(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4},
			want: 10,
		},
		{
			name: "Example 2",
			nums: []int{10, 20, 30},
			want: 2,
		},
		{
			name: "Example 3",
			nums: []int{1, 1, 1, 1},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsequencePairCount(tt.nums); got != tt.want {
				t.Errorf("subsequencePairCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
