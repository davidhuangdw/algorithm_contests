package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShipWithinDays(t *testing.T) {
	tests := []struct {
		name     string
		weights  []int
		days     int
		expected int
	}{
		{
			name:     "LeetCode Example 1",
			weights:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			days:     5,
			expected: 15,
		},
		{
			name:     "LeetCode Example 2",
			weights:  []int{3, 2, 2, 4, 1, 4},
			days:     3,
			expected: 6,
		},
		{
			name:     "LeetCode Example 3",
			weights:  []int{1, 2, 3, 1, 1},
			days:     4,
			expected: 3,
		},
		{
			name:     "Single day shipping",
			weights:  []int{1, 2, 3, 4, 5},
			days:     1,
			expected: 15,
		},
		{
			name:     "Each package per day",
			weights:  []int{1, 2, 3, 4, 5},
			days:     5,
			expected: 5,
		},
		{
			name:     "Large weights",
			weights:  []int{100, 200, 300, 400, 500},
			days:     2,
			expected: 900,
		},
		{
			name:     "Single package",
			weights:  []int{10},
			days:     1,
			expected: 10,
		},
		{
			name:     "Empty weights",
			weights:  []int{},
			days:     1,
			expected: 0,
		},
		{
			name:     "Equal weights",
			weights:  []int{5, 5, 5, 5, 5},
			days:     2,
			expected: 15,
		},
		{
			name:     "Complex distribution",
			weights:  []int{10, 50, 100, 10, 20, 30},
			days:     3,
			expected: 100,
		},
		{
			name:     "Minimum days required",
			weights:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			days:     9,
			expected: 9,
		},
		{
			name:     "Large number of packages",
			weights:  []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			days:     2,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := shipWithinDays(tt.weights, tt.days)
			assert.Equal(t, tt.expected, result)
		})
	}
}