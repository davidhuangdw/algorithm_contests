package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinEnd(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		x        int
		expected int
	}{
		{
			name:     "LeetCode example 1",
			n:        3,
			x:        4,
			expected: 6,
		},
		{
			name:     "LeetCode example 2",
			n:        2,
			x:        7,
			expected: 15,
		},
		{
			name:     "n=1, x=5",
			n:        1,
			x:        5,
			expected: 5,
		},
		{
			name:     "n=2, x=1",
			n:        2,
			x:        1,
			expected: 3,
		},
		{
			name:     "n=3, x=1",
			n:        3,
			x:        1,
			expected: 5,
		},
		{
			name:     "n=4, x=1",
			n:        4,
			x:        1,
			expected: 7,
		},
		{
			name:     "n=5, x=1",
			n:        5,
			x:        1,
			expected: 9,
		},
		{
			name:     "n=2, x=2",
			n:        2,
			x:        2,
			expected: 3,
		},
		{
			name:     "n=3, x=2",
			n:        3,
			x:        2,
			expected: 6,
		},
		{
			name:     "n=4, x=2",
			n:        4,
			x:        2,
			expected: 7,
		},
		{
			name:     "n=2, x=8",
			n:        2,
			x:        8,
			expected: 9,
		},
		{
			name:     "n=3, x=8",
			n:        3,
			x:        8,
			expected: 10,
		},
		{
			name:     "n=4, x=8",
			n:        4,
			x:        8,
			expected: 11,
		},
		{
			name:     "n=5, x=8",
			n:        5,
			x:        8,
			expected: 12,
		},
		{
			name:     "n=10, x=3",
			n:        10,
			x:        3,
			expected: 39,
		},
		{
			name:     "n=100, x=1",
			n:        100,
			x:        1,
			expected: 199,
		},
		{
			name:     "n=1000, x=10",
			n:        1000,
			x:        10,
			expected: 3999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minEnd(tt.n, tt.x)
			assert.Equal(t, tt.expected, result, "minEnd(%d, %d)", tt.n, tt.x)
		})
	}
}
