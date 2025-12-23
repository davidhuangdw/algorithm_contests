package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		h        int
		expected int
	}{
		// Basic cases
		{
			name:     "Single pile - exact time",
			piles:    []int{3},
			h:        3,
			expected: 1,
		},
		{
			name:     "Single pile - more time than needed",
			piles:    []int{3},
			h:        5,
			expected: 1,
		},
		{
			name:     "Single pile - less time than needed",
			piles:    []int{5},
			h:        3,
			expected: 2,
		},
		{
			name:     "Multiple piles - exact time",
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},

		// Classic examples
		{
			name:     "Classic example 1",
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},
		{
			name:     "Classic example 2",
			piles:    []int{30, 11, 23, 4, 20},
			h:        5,
			expected: 30,
		},
		{
			name:     "Classic example 3",
			piles:    []int{30, 11, 23, 4, 20},
			h:        6,
			expected: 23,
		},

		// Edge cases
		{
			name:     "Empty piles",
			piles:    []int{},
			h:        5,
			expected: 1,
		},
		{
			name:     "Single banana piles",
			piles:    []int{1, 1, 1, 1},
			h:        4,
			expected: 1,
		},
		{
			name:     "Large piles with ample time",
			piles:    []int{1000000000},
			h:        1000000000,
			expected: 1,
		},
		{
			name:     "Large piles with minimal time",
			piles:    []int{1000000000},
			h:        1,
			expected: 1000000000,
		},

		// Complex cases
		{
			name:     "Mixed pile sizes",
			piles:    []int{10, 20, 30, 40, 50},
			h:        15,
			expected: 10,
		},
		{
			name:     "All piles same size",
			piles:    []int{5, 5, 5, 5, 5},
			h:        5,
			expected: 5,
		},
		{
			name:     "All piles same size with more time",
			piles:    []int{5, 5, 5, 5, 5},
			h:        10,
			expected: 3,
		},
		{
			name:     "Very large piles",
			piles:    []int{1000, 2000, 3000},
			h:        10,
			expected: 667,
		},
		{
			name:     "Many small piles",
			piles:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			h:        10,
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minEatingSpeed(tt.piles, tt.h)
			assert.Equal(t, tt.expected, result,
				"minEatingSpeed(piles=%v, h=%d) = %d, expected %d",
				tt.piles, tt.h, result, tt.expected)
		})
	}
}
