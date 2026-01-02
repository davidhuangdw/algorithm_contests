package arrays_hashing

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "LeetCode example 1 - two majority elements",
			nums:     []int{3, 2, 3},
			expected: []int{3},
		},
		{
			name:     "LeetCode example 2 - single majority element",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "LeetCode example 3 - two majority elements",
			nums:     []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "No majority elements",
			nums:     []int{1, 1, 2, 2, 3, 3},
			expected: []int{},
		},
		{
			name:     "Single element repeated",
			nums:     []int{5, 5, 5, 5},
			expected: []int{5},
		},
		{
			name:     "Two elements with equal count",
			nums:     []int{1, 1, 2, 2, 1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Three elements with one majority",
			nums:     []int{3, 1, 3, 1, 3, 1, 3},
			expected: []int{1, 3},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "Complex pattern with multiple candidates",
			nums:     []int{1, 1, 1, 2, 2, 2, 3, 3, 1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Large array with clear majority",
			nums:     []int{1, 1, 1, 1, 1, 2, 2, 2, 3, 3},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := majorityElement(tt.nums)

			// Sort both slices for comparison since order doesn't matter
			sort.Ints(result)
			sort.Ints(tt.expected)

			assert.Equal(t, tt.expected, result)
		})
	}
}
