package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChange(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		coins    []int
		expected int
	}{
		{
			name:     "amount 0",
			amount:   0,
			coins:    []int{1, 2, 5},
			expected: 1,
		},
		{
			name:     "leetcode example 1",
			amount:   5,
			coins:    []int{1, 2, 5},
			expected: 4,
		},
		{
			name:     "leetcode example 2",
			amount:   3,
			coins:    []int{2},
			expected: 0,
		},
		{
			name:     "leetcode example 3",
			amount:   10,
			coins:    []int{10},
			expected: 1,
		},
		{
			name:     "single coin type",
			amount:   5,
			coins:    []int{1},
			expected: 1,
		},
		{
			name:     "multiple coin types",
			amount:   6,
			coins:    []int{1, 2, 3},
			expected: 7,
		},
		{
			name:     "no coins",
			amount:   5,
			coins:    []int{},
			expected: 0,
		},
		{
			name:     "large amount",
			amount:   100,
			coins:    []int{1, 5, 10, 25},
			expected: 242,
		},
		{
			name:     "coins with gaps",
			amount:   7,
			coins:    []int{2, 3, 5},
			expected: 2,
		},
		{
			name:     "impossible amount",
			amount:   1,
			coins:    []int{2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := change(tt.amount, tt.coins)
			assert.Equal(t, tt.expected, result, "change(%d, %v) = %d, want %d", tt.amount, tt.coins, result, tt.expected)
		})
	}
}
