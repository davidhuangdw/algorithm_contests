package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		// Basic cases
		{
			name:     "Single bar",
			heights:  []int{5},
			expected: 5,
		},
		{
			name:     "All same height",
			heights:  []int{2, 2, 2},
			expected: 6,
		},
		{
			name:     "Increasing heights",
			heights:  []int{1, 2, 3},
			expected: 4,
		},
		{
			name:     "Decreasing heights",
			heights:  []int{3, 2, 1},
			expected: 4,
		},

		// Classic examples
		{
			name:     "Classic histogram example",
			heights:  []int{2, 1, 5, 6, 2, 3},
			expected: 10,
		},
		{
			name:     "Another classic example",
			heights:  []int{6, 2, 5, 4, 5, 1, 6},
			expected: 12,
		},

		// Edge cases
		{
			name:     "Empty histogram",
			heights:  []int{},
			expected: 0,
		},
		{
			name:     "Single zero height",
			heights:  []int{0},
			expected: 0,
		},
		{
			name:     "All zeros",
			heights:  []int{0, 0, 0},
			expected: 0,
		},
		{
			name:     "Large single bar",
			heights:  []int{100},
			expected: 100,
		},

		// Complex patterns
		{
			name:     "Peak in middle",
			heights:  []int{1, 3, 5, 3, 1},
			expected: 9,
		},
		{
			name:     "Valley in middle",
			heights:  []int{5, 3, 2, 3, 5},
			expected: 10,
		},
		{
			name:     "Alternating heights",
			heights:  []int{4, 2, 4, 2, 4},
			expected: 10,
		},
		{
			name:     "Random pattern 1",
			heights:  []int{3, 1, 4, 2, 1, 3},
			expected: 6,
		},
		{
			name:     "Random pattern 2",
			heights:  []int{1, 2, 1, 2, 1, 2},
			expected: 6,
		},

		// Large numbers
		{
			name:     "Large numbers",
			heights:  []int{1000, 500, 1000},
			expected: 1500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestRectangleArea(tt.heights)
			assert.Equal(t, tt.expected, result,
				"largestRectangleArea(%v) = %d, expected %d",
				tt.heights, result, tt.expected)
		})
	}
}
