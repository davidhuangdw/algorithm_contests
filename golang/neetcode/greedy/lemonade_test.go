package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLemonadeChange(t *testing.T) {
	tests := []struct {
		name     string
		bills    []int
		expected bool
	}{
		{
			name:     "all $5 bills",
			bills:    []int{5, 5, 5, 5, 5},
			expected: true,
		},
		{
			name:     "LeetCode example 1",
			bills:    []int{5, 5, 5, 10, 20},
			expected: true,
		},
		{
			name:     "LeetCode example 2",
			bills:    []int{5, 5, 10, 10, 20},
			expected: false,
		},
		{
			name:     "exact change with $10 bills",
			bills:    []int{5, 10, 5, 10, 20},
			expected: false,
		},
		{
			name:     "insufficient change for $20",
			bills:    []int{5, 5, 10, 20},
			expected: true,
		},
		{
			name:     "multiple $20 bills with insufficient $5",
			bills:    []int{5, 5, 10, 20, 20},
			expected: false,
		},
		{
			name:     "complex scenario with mixed bills",
			bills:    []int{5, 5, 5, 10, 20, 10, 5, 20},
			expected: true,
		},
		{
			name:     "edge case: first customer pays with $20",
			bills:    []int{20},
			expected: false,
		},
		{
			name:     "edge case: single $10 bill",
			bills:    []int{10},
			expected: false,
		},
		{
			name:     "optimal change strategy",
			bills:    []int{5, 5, 10, 10, 5, 20},
			expected: true,
		},
		{
			name:     "large sequence of customers",
			bills:    []int{5, 5, 5, 5, 10, 10, 10, 10, 20, 20, 20, 20},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lemonadeChange(tt.bills)
			assert.Equal(t, tt.expected, result, "lemonadeChange(%v)", tt.bills)
		})
	}
}
