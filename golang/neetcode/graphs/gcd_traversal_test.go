package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanTraverseAllPairs(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		expected    bool
		description string
	}{
		{
			name:        "LeetCode Example 1",
			nums:        []int{2, 3, 6},
			expected:    true,
			description: "2 connected to 6, 3 connected to 6",
		},
		{
			name:        "LeetCode Example 2",
			nums:        []int{3, 9, 5},
			expected:    false,
			description: "3 and 9 are connected, but 5 is isolated",
		},
		{
			name:        "LeetCode Example 3",
			nums:        []int{4, 3, 12, 8},
			expected:    true,
			description: "All connected through 2 and 3",
		},
		{
			name:        "Array with 1",
			nums:        []int{2, 3, 6, 1},
			expected:    false,
			description: "Contains 1, which can't connect to any other number",
		},
		{
			name:        "Single element",
			nums:        []int{5},
			expected:    true,
			description: "Single element, trivially connected",
		},
		{
			name:        "Two elements with GCD > 1",
			nums:        []int{4, 6},
			expected:    true,
			description: "4 and 6 share GCD 2",
		},
		{
			name:        "Two elements with GCD = 1",
			nums:        []int{5, 7},
			expected:    false,
			description: "5 and 7 are coprime (GCD 1)",
		},
		{
			name:        "Multiple prime factors",
			nums:        []int{2, 4, 8, 16, 3},
			expected:    false,
			description: "All even numbers connected, but 3 is isolated",
		},
		{
			name:        "Chain connection",
			nums:        []int{2, 6, 12, 15, 25},
			expected:    true,
			description: "2-6-12-15-25 (connected through 2, 3, and 5)",
		},
		{
			name:        "Two disconnected groups",
			nums:        []int{2, 4, 8, 9, 27},
			expected:    false,
			description: "Even numbers (2,4,8) and multiples of 3 (9,27) are disconnected",
		},
		{
			name:        "All elements connected through multiple paths",
			nums:        []int{6, 10, 15, 30, 42, 56},
			expected:    true,
			description: "Connected through 2, 3, 5, 7",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canTraverseAllPairs(tt.nums)
			assert.Equal(t, tt.expected, result, "%s: Expected %t, got %t", tt.description, tt.expected, result)
		})
	}
}
