package backtrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalNQueens(t *testing.T) {
	// Test cases with known solutions for N-Queens problem
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n=1",
			n:    1,
			want: 1,
		},
		{
			name: "n=2",
			n:    2,
			want: 0,
		},
		{
			name: "n=3",
			n:    3,
			want: 0,
		},
		{
			name: "n=4",
			n:    4,
			want: 2,
		},
		{
			name: "n=5",
			n:    5,
			want: 10,
		},
		{
			name: "n=6",
			n:    6,
			want: 4,
		},
		{
			name: "n=7",
			n:    7,
			want: 40,
		},
		{
			name: "n=8",
			n:    8,
			want: 92,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := totalNQueens(tt.n)
			assert.Equal(t, tt.want, got, "totalNQueens(%d) = %d, want %d", tt.n, got, tt.want)
		})
	}
}
