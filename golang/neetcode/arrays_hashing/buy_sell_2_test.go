package arrays_hashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "LeetCode example 1 - increasing prices",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			name:     "LeetCode example 2 - always increasing",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "LeetCode example 3 - decreasing prices",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "Single day",
			prices:   []int{5},
			expected: 0,
		},
		{
			name:     "Empty prices",
			prices:   []int{},
			expected: 0,
		},
		{
			name:     "Two days - profit possible",
			prices:   []int{1, 2},
			expected: 1,
		},
		{
			name:     "Two days - no profit",
			prices:   []int{2, 1},
			expected: 0,
		},
		{
			name:     "Multiple peaks and valleys",
			prices:   []int{1, 5, 3, 8, 2, 10},
			expected: 17,
		},
		{
			name:     "All same prices",
			prices:   []int{5, 5, 5, 5, 5},
			expected: 0,
		},
		{
			name:     "Complex pattern with multiple transactions",
			prices:   []int{3, 3, 5, 0, 0, 3, 1, 4},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.prices)
			assert.Equal(t, tt.expected, result)
		})
	}
}