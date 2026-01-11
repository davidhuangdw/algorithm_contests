package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxTurbulenceSize(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "single element",
			arr:      []int{5},
			expected: 1,
		},
		{
			name:     "LeetCode example 1",
			arr:      []int{9, 4, 2, 10, 7, 8, 8, 1, 9},
			expected: 5,
		},
		{
			name:     "LeetCode example 2",
			arr:      []int{4, 8, 12, 16},
			expected: 2,
		},
		{
			name:     "LeetCode example 3",
			arr:      []int{100},
			expected: 1,
		},
		{
			name:     "alternating pattern",
			arr:      []int{1, 2, 1, 2, 1, 2, 1},
			expected: 7,
		},
		{
			name:     "strictly increasing",
			arr:      []int{1, 2, 3, 4, 5},
			expected: 2,
		},
		{
			name:     "strictly decreasing",
			arr:      []int{5, 4, 3, 2, 1},
			expected: 2,
		},
		{
			name:     "all equal elements",
			arr:      []int{5, 5, 5, 5, 5},
			expected: 1,
		},
		{
			name:     "complex turbulent array",
			arr:      []int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24},
			expected: 8,
		},
		{
			name:     "two elements increasing",
			arr:      []int{1, 2},
			expected: 2,
		},
		{
			name:     "two elements decreasing",
			arr:      []int{2, 1},
			expected: 2,
		},
		{
			name:     "two elements equal",
			arr:      []int{2, 2},
			expected: 1,
		},
		{
			name:     "long turbulent sequence",
			arr:      []int{1, 10, 1, 10, 1, 10, 1, 10, 1, 10},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxTurbulenceSize(tt.arr)
			assert.Equal(t, tt.expected, result, "maxTurbulenceSize(%v)", tt.arr)
		})
	}
}
