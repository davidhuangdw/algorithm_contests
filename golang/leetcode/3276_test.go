package main

import (
	"testing"
)

func Test_maxScore(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{{1, 2, 3}, {4, 3, 2}, {1, 1, 1}},
			want: 8,
		},
		{
			name: "Example 2",
			grid: [][]int{{8, 7, 6}, {8, 3, 2}},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.grid); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
