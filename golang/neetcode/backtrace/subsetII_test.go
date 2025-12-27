package backtrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "empty array",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: [][]int{{}, {1}},
		},
		{
			name:     "two elements no duplicates",
			nums:     []int{1, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name:     "with duplicates - classic case",
			nums:     []int{1, 2, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
		{
			name: "multiple duplicates",
			nums: []int{1, 1, 2, 2},
			expected: [][]int{
				{}, {1}, {2}, {1, 1}, {1, 2}, {2, 2},
				{1, 1, 2}, {1, 2, 2}, {1, 1, 2, 2},
			},
		},
		{
			name:     "all duplicates",
			nums:     []int{2, 2, 2},
			expected: [][]int{{}, {2}, {2, 2}, {2, 2, 2}},
		},
		{
			name: "negative numbers with duplicates",
			nums: []int{-1, 0, 0, 1},
			expected: [][]int{
				{}, {-1}, {0}, {1}, {-1, 0}, {-1, 1}, {0, 0}, {0, 1},
				{-1, 0, 0}, {-1, 0, 1}, {0, 0, 1}, {-1, 0, 0, 1},
			},
		},
		{
			name: "three elements no duplicates",
			nums: []int{1, 2, 3},
			expected: [][]int{
				{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subsetsWithDup(tt.nums)

			// Sort both result and expected for consistent comparison
			sortSubsets(result)
			sortSubsets(tt.expected)

			assert.Equal(t, tt.expected, result,
				"subsetsWithDup(%v) should equal %v", tt.nums, tt.expected)
		})
	}
}
