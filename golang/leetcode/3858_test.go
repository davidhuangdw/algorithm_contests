package main

import "testing"

func Test_minimumOR(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "can achieve zero",
			grid: [][]int{{0, 1}, {0, 2}},
			want: 0, // select 0 from both rows
		},
		{
			name: "eliminate high bit",
			grid: [][]int{{1, 2}, {4, 3}},
			want: 3, // bit 4 can be eliminated, select 2 from row 0, 3 from row 1 -> 2|3=3
		},
		{
			name: "all same values",
			grid: [][]int{{7, 7}, {7, 7}},
			want: 7, // no choice, must be 7
		},
		{
			name: "single row - pick minimum",
			grid: [][]int{{1, 2, 4}},
			want: 1, // single row, just pick minimum value
		},
		{
			name: "single column - no choice",
			grid: [][]int{{1}, {2}, {4}},
			want: 7, // must select all -> 1|2|4=7
		},
		{
			name: "cannot eliminate any bit",
			grid: [][]int{{5, 5}, {5, 5}},
			want: 5, // all elements are 5
		},
		{
			name: "3x3 grid with zeros",
			grid: [][]int{{0, 1, 2}, {0, 3, 4}, {0, 5, 6}},
			want: 0, // all rows have 0
		},
		{
			name: "bit pattern test",
			grid: [][]int{{8, 7}, {4, 11}, {2, 13}},
			want: 7, // 8=1000, 7=0111, 4=0100, 11=1011, 2=0010, 13=1101
		},
		{
			name: "partial elimination",
			grid: [][]int{{3, 6}, {5, 6}},
			want: 6, // 3=011, 6=110, 5=101 -> bit 4 cannot be eliminated
		},
		{
			name: "multiple rows same pattern",
			grid: [][]int{{1, 8}, {2, 8}, {4, 8}},
			want: 7, // select 1, 2, 4 -> 1|2|4=7
		},
		{
			name: "all bits eliminable except one",
			grid: [][]int{{1, 2, 3}, {1, 2, 3}},
			want: 1, // can select 1 from both rows
		},
		{
			name: "large values",
			grid: [][]int{{255, 128}, {255, 64}},
			want: 192, // 255=11111111, 128=10000000, 64=01000000 -> 128|64=192
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gridCopy := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}
			got := minimumOR(gridCopy)
			if got != tt.want {
				t.Errorf("minimumOR() = %v (binary: %b), want %v (binary: %b)", got, got, tt.want, tt.want)
			}
		})
	}
}
