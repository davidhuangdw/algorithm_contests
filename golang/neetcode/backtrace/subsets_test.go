package backtrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "empty set",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: [][]int{{}, {1}},
		},
		{
			name:     "two elements",
			nums:     []int{1, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name:     "three elements",
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:     "with duplicates",
			nums:     []int{1, 2, 2},
			expected: [][]int{{}, {1}, {2}, {2}, {1, 2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, 0, 1},
			expected: [][]int{{}, {-1}, {0}, {1}, {-1, 0}, {-1, 1}, {0, 1}, {-1, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subsets(tt.nums)

			// Sort both result and expected for comparison
			sortSubsets(result)
			sortSubsets(tt.expected)

			assert.Equal(t, tt.expected, result,
				"subsets(%v) should equal %v", tt.nums, tt.expected)
		})
	}
}
