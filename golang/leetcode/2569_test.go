package main

import (
	"reflect"
	"testing"
)

func TestHandleQuery(t *testing.T) {
	tests := []struct {
		name    string
		nums1   []int
		nums2   []int
		queries [][]int
		want    []int64
	}{
		{
			name:    "Example 1",
			nums1:   []int{1, 0, 1},
			nums2:   []int{0, 0, 0},
			queries: [][]int{{1, 1, 1}, {2, 1, 0}, {3, 0, 0}},
			want:    []int64{3},
		},
		{
			name:    "Single Element",
			nums1:   []int{1},
			nums2:   []int{5},
			queries: [][]int{{2, 10, 0}, {3, 0, 0}, {1, 0, 0}, {2, 10, 0}, {3, 0, 0}},
			want:    []int64{15, 15},
		},
		{
			name:    "Range Flip and Multiple Type 2",
			nums1:   []int{0, 0, 0},
			nums2:   []int{1, 1, 1},
			queries: [][]int{{1, 0, 2}, {2, 5, 0}, {3, 0, 0}, {1, 1, 1}, {2, 10, 0}, {3, 0, 0}},
			want:    []int64{18, 38},
			// Initial: res = 3, nums1 = [0,0,0]
			// Q1 [1,0,2]: nums1 = [1,1,1], sm[1] = 3
			// Q2 [2,5,0]: res = 3 + 3*5 = 18
			// Q3 [3,0,0]: ans = [18]
			// Q4 [1,1,1]: flip index 1. nums1 = [1,0,1], sm[1] = 2
			// Q5 [2,10,0]: res = 18 + 2*10 = 38
			// Q6 [3,0,0]: ans = [18, 38]
		},
		{
			name:    "Empty result",
			nums1:   []int{1, 1},
			nums2:   []int{1, 1},
			queries: [][]int{{1, 0, 1}, {2, 10, 0}},
			want:    []int64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleQuery(tt.nums1, tt.nums2, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
