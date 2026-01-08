package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSquares(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n=1 (perfect square)",
			n:    1,
			want: 1, // 1
		},
		{
			name: "n=2 (1+1)",
			n:    2,
			want: 2, // 1+1
		},
		{
			name: "n=3 (1+1+1)",
			n:    3,
			want: 3, // 1+1+1
		},
		{
			name: "n=4 (perfect square)",
			n:    4,
			want: 1, // 4
		},
		{
			name: "n=5 (4+1)",
			n:    5,
			want: 2, // 4+1
		},
		{
			name: "n=6 (4+1+1)",
			n:    6,
			want: 3, // 4+1+1
		},
		{
			name: "n=7 (4+1+1+1)",
			n:    7,
			want: 4, // 4+1+1+1
		},
		{
			name: "n=8 (4+4)",
			n:    8,
			want: 2, // 4+4
		},
		{
			name: "n=9 (perfect square)",
			n:    9,
			want: 1, // 9
		},
		{
			name: "n=10 (9+1)",
			n:    10,
			want: 2, // 9+1
		},
		{
			name: "n=12 (4+4+4)",
			n:    12,
			want: 3, // 4+4+4
		},
		{
			name: "n=13 (9+4)",
			n:    13,
			want: 2, // 9+4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numSquares(tt.n)
			assert.Equal(t, tt.want, got, "numSquares(%d) = %d, want %d", tt.n, got, tt.want)
		})
	}
}
