package main

import (
	"reflect"
	"testing"
)

func Test_moQueries(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		q    [][2]int
		want []int
	}{
		{
			name: "single query all different",
			nums: []int{0, 1, 2, 3, 4}, // 1-indexed: [1,2,3,4]
			q:    [][2]int{{1, 4}},
			want: []int{4}, // each appears once: 1^2*4 = 4
		},
		{
			name: "single query with duplicates",
			nums: []int{0, 1, 1, 2, 2}, // 1-indexed: [1,1,2,2]
			q:    [][2]int{{1, 4}},
			want: []int{8}, // 1 appears 2x, 2 appears 2x: 2^2+2^2 = 8
		},
		{
			name: "multiple queries",
			nums: []int{0, 1, 2, 1, 2, 1}, // 1-indexed: [1,2,1,2,1]
			q:    [][2]int{{1, 3}, {2, 5}},
			want: []int{5, 8}, // [1,2,1]: 2^2+1^2=5, [2,1,2,1]: 2^2+2^2=8
		},
		{
			name: "all same value",
			nums: []int{0, 5, 5, 5},
			q:    [][2]int{{1, 3}},
			want: []int{9}, // 5 appears 3x: 3^2 = 9
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := moQueries(tt.nums, tt.q)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
