package dp

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
			name:     "empty prices",
			prices:   []int{},
			expected: 0,
		},
		{
			name:     "single price",
			prices:   []int{1},
			expected: 0,
		},
			{
			name:     "decreasing prices",
			prices:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "leetcode example 1",
			prices:   []int{1, 2, 3, 0, 2},
			expected: 3,
		},
		{
			name:     "leetcode example 2",
			prices:   []int{1},
			expected: 0,
		},
		{
			name:     "simple profit",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "cooldown required",
			prices:   []int{1, 3, 2, 4},
			expected: 3,
		},
		{
			name:     "multiple peaks and valleys",
			prices:   []int{3, 1, 4, 2, 5, 1, 6},
			expected: 8,
		},
		{
			name:     "all same prices",
			prices:   []int{5, 5, 5, 5, 5},
			expected: 0,
		},
		{
			name:     "large price range",
			prices:   []int{100, 1, 200, 3, 300},
			expected: 299,
		},
		{
			name:     "cooldown prevents consecutive buys",
			prices:   []int{1, 2, 1, 2, 1, 2},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.prices)
			assert.Equal(t, tt.expected, result, "maxProfit(%v) = %d, want %d", tt.prices, result, tt.expected)
		})
	}
}