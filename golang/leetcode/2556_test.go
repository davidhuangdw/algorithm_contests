package main

import (
	"testing"
)

func TestIsPossibleToCutPath(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want bool
	}{
		{
			name: "Example 1 - disconnected by 1 flip",
			grid: [][]int{
				{1, 1, 1},
				{1, 0, 0},
				{1, 1, 1},
			},
			want: true,
		},
		{
			name: "Example 2 - cannot be disconnected",
			grid: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: false,
		},
		{
			name: "Already disconnected",
			grid: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			want: true,
		},
		{
			name: "Multiple paths but can cut at a bottleneck",
			grid: [][]int{
				{1, 1, 1},
				{0, 1, 0},
				{1, 1, 1},
			},
			want: true,
		},
		{
			name: "Larger grid with 1 bottleneck",
			grid: [][]int{
				{1, 1, 0, 1},
				{1, 1, 1, 1},
				{0, 1, 1, 1},
				{0, 0, 1, 1},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// deep copy since the algorithm modifies the slice in place
			gridCopy := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}

			got := isPossibleToCutPath(gridCopy)
			if got != tt.want {
				t.Errorf("isPossibleToCutPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
