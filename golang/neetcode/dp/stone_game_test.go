package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoneGame(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		expected bool
	}{
		{
			name:     "empty piles",
			piles:    []int{},
			expected: true,
		},
		{
			name:     "single pile",
			piles:    []int{5},
			expected: true,
		},
		{
			name:     "two piles - Alice wins",
			piles:    []int{5, 3},
			expected: true,
		},
		{
			name:     "leetcode example 1",
			piles:    []int{5, 3, 4, 5},
			expected: true,
		},
		{
			name:     "three piles - Alice lose",
			piles:    []int{3, 7, 2},
			expected: false,
		},
		{
			name:     "four piles - Alice wins",
			piles:    []int{3, 2, 10, 4},
			expected: true,
		},
		{
			name:     "all same piles",
			piles:    []int{3, 3, 3, 3},
			expected: true,
		},
		{
			name:     "alternating high-low",
			piles:    []int{10, 1, 10, 1},
			expected: true,
		},
		{
			name:     "increasing order",
			piles:    []int{1, 2, 3, 4},
			expected: true,
		},
		{
			name:     "decreasing order",
			piles:    []int{4, 3, 2, 1},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stoneGame(tt.piles)
			assert.Equal(t, tt.expected, result)
		})
	}
}
