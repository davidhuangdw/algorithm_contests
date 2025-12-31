package math_geo

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyPow(t *testing.T) {
	tests := []struct {
		name     string
		x        float64
		n        int
		expected float64
	}{
		{
			name:     "leetcode example 1",
			x:        2.0,
			n:        10,
			expected: 1024.0,
		},
		{
			name:     "leetcode example 2",
			x:        2.1,
			n:        3,
			expected: 9.261,
		},
		{
			name:     "leetcode example 3",
			x:        2.0,
			n:        -2,
			expected: 0.25,
		},
		{
			name:     "power of zero",
			x:        5.0,
			n:        0,
			expected: 1.0,
		},
		{
			name:     "negative base with positive exponent",
			x:        -2.0,
			n:        3,
			expected: -8.0,
		},
		{
			name:     "negative base with negative exponent",
			x:        -2.0,
			n:        -3,
			expected: -0.125,
		},
		{
			name:     "fractional base with positive exponent",
			x:        0.5,
			n:        3,
			expected: 0.125,
		},
		{
			name:     "fractional base with negative exponent",
			x:        0.5,
			n:        -3,
			expected: 8.0,
		},
		{
			name:     "large positive exponent",
			x:        1.0001,
			n:        10000,
			expected: math.Pow(1.0001, 10000),
		},
		{
			name:     "large negative exponent",
			x:        1.0001,
			n:        -10000,
			expected: math.Pow(1.0001, -10000),
		},
		{
			name:     "base 1 with any exponent",
			x:        1.0,
			n:        1000,
			expected: 1.0,
		},
		{
			name:     "base -1 with even exponent",
			x:        -1.0,
			n:        100,
			expected: 1.0,
		},
		{
			name:     "base -1 with odd exponent",
			x:        -1.0,
			n:        101,
			expected: -1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := myPow(tt.x, tt.n)
			assert.InEpsilon(t, tt.expected, result, 1e-9, "myPow(%f, %d) = %f, want %f", tt.x, tt.n, result, tt.expected)
		})
	}
}