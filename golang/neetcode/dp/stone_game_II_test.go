package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoneGameII(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		expected int
	}{
		{
			name:     "empty piles",
			piles:    []int{},
			expected: 0,
		},
		{
			name:     "single pile",
			piles:    []int{5},
			expected: 5,
		},
		{
			name:     "two piles - leetcode example",
			piles:    []int{2, 7},
			expected: 9,
		},
		{
			name:     "leetcode example 1",
			piles:    []int{2, 7, 9, 4, 4},
			expected: 10,
		},
		{
			name:     "three piles",
			piles:    []int{1, 2, 3},
			expected: 3,
		},
		{
			name:     "four piles",
			piles:    []int{1, 2, 3, 4},
			expected: 5,
		},
		{
			name:     "all same piles",
			piles:    []int{3, 3, 3, 3},
			expected: 6,
		},
		{
			name:     "alternating high-low",
			piles:    []int{10, 1, 10, 1},
			expected: 11,
		},
		{
			name:     "increasing order",
			piles:    []int{1, 2, 3, 4, 5},
			expected: 8,
		},
		{
			name:     "decreasing order",
			piles:    []int{5, 4, 3, 2, 1},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stoneGameII(tt.piles)
			assert.Equal(t, tt.expected, result)
		})
	}
}
