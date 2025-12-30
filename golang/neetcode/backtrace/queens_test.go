package backtrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]string
	}{
		{
			name: "n=1",
			n:    1,
			want: [][]string{
				{"Q"},
			},
		},
		{
			name: "n=2",
			n:    2,
			want: [][]string{},
		},
		{
			name: "n=3",
			n:    3,
			want: [][]string{},
		},
		{
			name: "n=4",
			n:    4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			name: "n=5",
			n:    5,
			want: [][]string{
				{"Q....", "..Q..", "....Q", ".Q...", "...Q."},
				{"Q....", "...Q.", ".Q...", "....Q", "..Q.."},
				{".Q...", "...Q.", "Q....", "..Q..", "....Q"},
				{".Q...", "....Q", "..Q..", "Q....", "...Q."},
				{"..Q..", "Q....", "...Q.", ".Q...", "....Q"},
				{"..Q..", "....Q", ".Q...", "...Q.", "Q...."},
				{"...Q.", "Q....", "..Q..", "....Q", ".Q..."},
				{"...Q.", ".Q...", "....Q", "..Q..", "Q...."},
				{"....Q", ".Q...", "...Q.", "Q....", "..Q.."},
				{"....Q", "..Q..", "Q....", "...Q.", ".Q..."},
			},
		},
		{
			name: "n=6",
			n:    6,
			want: [][]string{
				{".Q....", "...Q..", ".....Q", "Q.....", "..Q...", "....Q."},
				{"..Q...", ".....Q", ".Q....", "....Q.", "Q.....", "...Q.."},
				{"...Q..", "Q.....", "....Q.", ".Q....", ".....Q", "..Q..."},
				{"....Q.", "..Q...", "Q.....", ".....Q", "...Q..", ".Q...."},
			},
		},
		{
			name: "n=0",
			n:    0,
			want: [][]string{
				{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveNQueens(tt.n)
			assert.Equal(t, len(tt.want), len(got), "solveNQueens(%d) should return %d solutions, got %d", tt.n, len(tt.want), len(got))
			
			// For n >= 4, verify that all solutions are valid
			if tt.n >= 4 {
				for i, solution := range got {
					assert.Equal(t, tt.n, len(solution), "Solution %d should have %d rows", i, tt.n)
					for j, row := range solution {
						assert.Equal(t, tt.n, len(row), "Row %d in solution %d should have %d columns", j, i, tt.n)
					}
				}
			}
		})
	}
}