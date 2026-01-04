package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStockSpanner(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected []int
	}{
		{
			name:     "LeetCode Example 1",
			prices:   []int{100, 80, 60, 70, 60, 75, 85},
			expected: []int{1, 1, 1, 2, 1, 4, 6},
		},
		{
			name:     "Increasing prices",
			prices:   []int{10, 20, 30, 40, 50},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Decreasing prices",
			prices:   []int{50, 40, 30, 20, 10},
			expected: []int{1, 1, 1, 1, 1},
		},
		{
			name:     "Constant prices",
			prices:   []int{10, 10, 10, 10, 10},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Mixed prices with peaks",
			prices:   []int{10, 5, 15, 7, 20, 3, 25},
			expected: []int{1, 1, 3, 1, 5, 1, 7},
		},
		{
			name:     "Single price",
			prices:   []int{100},
			expected: []int{1},
		},
		{
			name:     "Empty prices",
			prices:   []int{},
			expected: []int{},
		},
		{
			name:     "Large span calculation",
			prices:   []int{1, 2, 3, 2, 1, 4, 3, 2, 1, 5},
			expected: []int{1, 2, 3, 1, 1, 6, 1, 1, 1, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spanner := ConstructorSS()
			results := make([]int, 0, len(tt.prices))
			
			for _, price := range tt.prices {
				span := spanner.Next(price)
				results = append(results, span)
			}
			
			assert.Equal(t, tt.expected, results)
		})
	}
}