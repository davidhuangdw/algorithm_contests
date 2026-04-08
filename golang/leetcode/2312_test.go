package main

import "testing"

func TestSellingWood(t *testing.T) {
	tests := []struct {
		m      int
		n      int
		prices [][]int
		want   int64
	}{
		{3, 5, [][]int{{1, 4, 2}, {2, 2, 7}, {2, 1, 3}}, 19},
		{4, 6, [][]int{{3, 2, 10}, {1, 4, 2}, {4, 1, 3}}, 32},
	}
	for i, tt := range tests {
		if got := sellingWood(tt.m, tt.n, tt.prices); got != tt.want {
			t.Errorf("case %d: sellingWood(%v, %v, %v) = %v; want %v", i, tt.m, tt.n, tt.prices, got, tt.want)
		}
	}
}
