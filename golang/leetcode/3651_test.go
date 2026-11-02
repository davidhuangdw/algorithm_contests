package main

import "testing"

func Test_minCost1(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		k    int
		want int
	}{
		{"1", [][]int{{1, 3, 3}, {2, 5, 4}, {4, 3, 5}}, 2, 7},
		{"Example 2", [][]int{{1, 2}, {2, 3}, {3, 4}}, 1, 9},
		{"No Teleportation", [][]int{{1, 1}, {1, 1}}, 0, 2},
		{"Teleport All Zero Cost", [][]int{{5, 5, 5}, {5, 5, 5}, {5, 5, 5}}, 2, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCost1(tt.grid, tt.k)
			if got != tt.want {
				t.Errorf("minCost1() = %v, want %v", got, tt.want)
			}
		})
	}
}
