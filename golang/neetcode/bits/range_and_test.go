package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeBitwiseAnd(t *testing.T) {
	tests := []struct {
		name     string
		left     int
		right    int
		expected int
	}{
		{
			name:     "LeetCode example 1",
			left:     5,
			right:    7,
			expected: 4,
		},
		{
			name:     "LeetCode example 2",
			left:     0,
			right:    0,
			expected: 0,
		},
		{
			name:     "LeetCode example 3",
			left:     1,
			right:    2147483647,
			expected: 0,
		},
		{
			name:     "same number",
			left:     10,
			right:    10,
			expected: 10,
		},
		{
			name:     "consecutive numbers",
			left:     8,
			right:    9,
			expected: 8,
		},
		{
			name:     "small range",
			left:     10,
			right:    11,
			expected: 10,
		},
		{
			name:     "medium range",
			left:     12,
			right:    15,
			expected: 12,
		},
		{
			name:     "large range",
			left:     100,
			right:    200,
			expected: 0,
		},
		{
			name:     "power of two boundary",
			left:     15,
			right:    16,
			expected: 0,
		},
		{
			name:     "another power of two boundary",
			left:     31,
			right:    32,
			expected: 0,
		},
		{
			name:     "range within same power of two",
			left:     16,
			right:    31,
			expected: 16,
		},
		{
			name:     "range crossing multiple powers",
			left:     10,
			right:    20,
			expected: 0,
		},
		{
			name:     "max int range",
			left:     2147483646,
			right:    2147483647,
			expected: 2147483646,
		},
		{
			name:     "zero to small number",
			left:     0,
			right:    5,
			expected: 0,
		},
		{
			name:     "one to small number",
			left:     1,
			right:    3,
			expected: 0,
		},
		{
			name:     "range with common prefix",
			left:     26,
			right:    30,
			expected: 24,
		},
		{
			name:     "range with no common prefix",
			left:     15,
			right:    16,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rangeBitwiseAnd(tt.left, tt.right)
			assert.Equal(t, tt.expected, result, "rangeBitwiseAnd(%d, %d)", tt.left, tt.right)
		})
	}
}
