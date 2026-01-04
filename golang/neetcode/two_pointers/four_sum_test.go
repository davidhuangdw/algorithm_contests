package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected [][]int
	}{
		{
			name:   "leetcode example 1",
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			expected: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			name:     "leetcode example 2",
			nums:     []int{2, 2, 2, 2, 2},
			target:   8,
			expected: [][]int{{2, 2, 2, 2}},
		},
		{
			name:     "empty array",
			nums:     []int{},
			target:   0,
			expected: [][]int{},
		},
		{
			name:     "less than 4 elements",
			nums:     []int{1, 2, 3},
			target:   6,
			expected: [][]int{},
		},
		{
			name:   "negative target",
			nums:   []int{-3, -2, -1, 0, 0, 1, 2, 3},
			target: -5,
			expected: [][]int{
				{-3, -2, -1, 1},
				{-3, -2, 0, 0},
			},
		},
		{
			name:     "large numbers",
			nums:     []int{1000, 2000, 3000, 4000, 5000},
			target:   10000,
			expected: [][]int{{1000, 2000, 3000, 4000}},
		},
		{
			name:     "duplicates handling",
			nums:     []int{0, 0, 0, 0, 0, 0},
			target:   0,
			expected: [][]int{{0, 0, 0, 0}},
		},
		{
			name:     "no solution",
			nums:     []int{1, 2, 3, 4, 5},
			target:   100,
			expected: [][]int{},
		},
		{
			name:   "mixed positive negative",
			nums:   []int{-4, -3, -2, -1, 1, 2, 3, 4},
			target: 0,
			expected: [][]int{
				{-4, -3, 3, 4},
				{-4, -2, 2, 4},
				{-4, -1, 1, 4},
				{-4, -1, 2, 3},
				{-3, -2, 1, 4},
				{-3, -2, 2, 3},
				{-3, -1, 1, 3},
				{-2, -1, 1, 2},
			},
		},
		{
			name:     "exact four elements",
			nums:     []int{1, 2, 3, 4},
			target:   10,
			expected: [][]int{{1, 2, 3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fourSum(tt.nums, tt.target)
			assert.ElementsMatch(t, tt.expected, result, "nums=%v, target=%d", tt.nums, tt.target)
		})
	}
}
