package main

import (
	"reflect"
	"testing"
)

func TestMaximizeXor(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		queries [][]int
		want    []int
	}{
		{
			name: "Example 1",
			nums: []int{0, 1, 2, 3, 4},
			queries: [][]int{
				{3, 1},
				{1, 3},
				{5, 6},
			},
			want: []int{3, 3, 7},
		},
		{
			name: "Example 2",
			nums: []int{5, 2, 4, 6, 6, 3},
			queries: [][]int{
				{12, 4},
				{8, 1},
				{6, 3},
			},
			want: []int{15, -1, 5},
		},
		{
			name: "No elements satisfy condition",
			nums: []int{10, 20, 30},
			queries: [][]int{
				{5, 5},
			},
			want: []int{-1},
		},
		{
			name: "All elements satisfy condition",
			nums: []int{1, 2, 3},
			queries: [][]int{
				{4, 10},
			},
			want: []int{7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeXor(tt.nums, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximizeXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
