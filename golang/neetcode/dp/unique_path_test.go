package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "1x1 grid",
			m:    1,
			n:    1,
			want: 1,
		},
		{
			name: "1x2 grid",
			m:    1,
			n:    2,
			want: 1,
		},
		{
			name: "2x1 grid",
			m:    2,
			n:    1,
			want: 1,
		},
		{
			name: "2x2 grid",
			m:    2,
			n:    2,
			want: 2,
		},
		{
			name: "3x2 grid",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "3x3 grid",
			m:    3,
			n:    3,
			want: 6,
		},
		{
			name: "3x7 grid",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "7x3 grid",
			m:    7,
			n:    3,
			want: 28,
		},
		{
			name: "leetcode example 1",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "leetcode example 2",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "large grid 10x10",
			m:    10,
			n:    10,
			want: 48620,
		},
		{
			name: "tall grid 1x10",
			m:    1,
			n:    10,
			want: 1,
		},
		{
			name: "wide grid 10x1",
			m:    10,
			n:    1,
			want: 1,
		},
		{
			name: "medium grid 5x5",
			m:    5,
			n:    5,
			want: 70,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePaths(tt.m, tt.n)
			assert.Equal(t, tt.want, got, "uniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, got, tt.want)
		})
	}
}
