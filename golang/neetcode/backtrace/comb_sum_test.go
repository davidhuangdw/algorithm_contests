package backtrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected [][]int
	}{
		{
			name:     "basic case",
			nums:     []int{2, 3, 6, 7},
			target:   7,
			expected: [][]int{{7}, {2, 2, 3}},
		},
		{
			name:     "multiple combinations",
			nums:     []int{2, 3, 5},
			target:   8,
			expected: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:     "no combinations possible",
			nums:     []int{2},
			target:   1,
			expected: [][]int{},
		},
		{
			name:     "single number multiple times",
			nums:     []int{1},
			target:   3,
			expected: [][]int{{1, 1, 1}},
		},
		{
			name:     "empty result for zero target",
			nums:     []int{1, 2, 3},
			target:   0,
			expected: [][]int{[]int{}},
		},
		{
			name:     "with duplicates in input",
			nums:     []int{2, 2, 3},
			target:   5,
			expected: [][]int{{2, 3}, {2, 3}},
		},
		{
			name:     "large numbers",
			nums:     []int{10, 20, 30},
			target:   50,
			expected: [][]int{{10, 10, 10, 10, 10}, {10, 10, 10, 20}, {10, 10, 30}, {20, 30}, {10, 20, 20}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := combinationSum(tt.nums, tt.target)

			// Sort both result and expected for consistent comparison
			sortCombinations(result)
			sortCombinations(tt.expected)

			assert.Equal(t, tt.expected, result,
				"combinationSum(%v, %d) should equal %v", tt.nums, tt.target, tt.expected)
		})
	}
}
