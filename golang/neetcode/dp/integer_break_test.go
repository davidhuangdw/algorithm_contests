package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerBreak(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n=2",
			n:    2,
			want: 1, // 1+1
		},
		{
			name: "n=3",
			n:    3,
			want: 2, // 1+2
		},
		{
			name: "n=4",
			n:    4,
			want: 4, // 2+2
		},
		{
			name: "n=5",
			n:    5,
			want: 6, // 2+3
		},
		{
			name: "n=6",
			n:    6,
			want: 9, // 3+3
		},
		{
			name: "n=7",
			n:    7,
			want: 12, // 3+4
		},
		{
			name: "n=8",
			n:    8,
			want: 18, // 3+3+2
		},
		{
			name: "n=9",
			n:    9,
			want: 27, // 3+3+3
		},
		{
			name: "n=10",
			n:    10,
			want: 36, // 3+3+4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := integerBreak(tt.n)
			assert.Equal(t, tt.want, got, "integerBreak(%d) = %d, want %d", tt.n, got, tt.want)
		})
	}
}
