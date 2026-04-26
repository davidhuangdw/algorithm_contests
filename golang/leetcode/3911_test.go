package main

import (
	"reflect"
	"testing"
)

func Test_kthRemainingInteger(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		queries [][]int
		want    []int
	}{
		{
			name:    "Example 1",
			nums:    []int{1, 4, 7},
			queries: [][]int{{0, 2, 1}, {1, 1, 2}, {0, 0, 3}},
			want:    []int{2, 6, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthRemainingInteger(tt.nums, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthRemainingInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
