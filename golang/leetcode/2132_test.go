package main

import "testing"

func TestPossibleToStamp(t *testing.T) {
	tests := []struct {
		grid        [][]int
		stampHeight int
		stampWidth  int
		want        bool
	}{
		{
			[][]int{
				{1, 0, 0, 0},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
			},
			4, 3,
			true,
		},
		{
			[][]int{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
			2, 2,
			false,
		},
		{
			[][]int{{0, 0, 0}, {0, 0, 0}},
			2, 2,
			true,
		},
	}
	for i, tt := range tests {
		ans := possibleToStamp(tt.grid, tt.stampHeight, tt.stampWidth)
		if ans != tt.want {
			t.Errorf("case %d: possibleToStamp() = %v; want %v", i, ans, tt.want)
		}
	}
}
