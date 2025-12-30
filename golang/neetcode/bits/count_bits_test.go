package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "n = 0",
			n:    0,
			want: []int{0},
		},
		{
			name: "n = 1",
			n:    1,
			want: []int{0, 1},
		},
		{
			name: "n = 2",
			n:    2,
			want: []int{0, 1, 1},
		},
		{
			name: "n = 3",
			n:    3,
			want: []int{0, 1, 1, 2},
		},
		{
			name: "n = 4",
			n:    4,
			want: []int{0, 1, 1, 2, 1},
		},
		{
			name: "n = 5",
			n:    5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
		{
			name: "n = 6",
			n:    6,
			want: []int{0, 1, 1, 2, 1, 2, 2},
		},
		{
			name: "n = 7",
			n:    7,
			want: []int{0, 1, 1, 2, 1, 2, 2, 3},
		},
		{
			name: "n = 8",
			n:    8,
			want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1},
		},
		{
			name: "leetcode example 1",
			n:    2,
			want: []int{0, 1, 1},
		},
		{
			name: "leetcode example 2",
			n:    5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countBits(tt.n)
			assert.Equal(t, tt.want, got, "countBits(%d) = %v, want %v", tt.n, got, tt.want)
		})
	}
}