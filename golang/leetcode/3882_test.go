package main

import "testing"

func Test_minCost(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"Test Case 1", [][]int{{0, 1}, {1, 0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.grid); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
