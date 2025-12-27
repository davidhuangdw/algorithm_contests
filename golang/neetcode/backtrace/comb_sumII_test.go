package backtrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "basic case with duplicates",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expected:   [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			name:       "no duplicates in result",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expected:   [][]int{{1, 2, 2}, {5}},
		},
		{
			name:       "no combinations possible",
			candidates: []int{2, 3},
			target:     1,
			expected:   [][]int{},
		},
		{
			name:       "single number exact match",
			candidates: []int{1, 1},
			target:     2,
			expected:   [][]int{{1, 1}},
		},
		{
			name:       "empty result for zero target",
			candidates: []int{1, 2, 3},
			target:     0,
			expected:   [][]int{[]int{}},
		},
		{
			name:       "all numbers used",
			candidates: []int{1, 2, 3},
			target:     6,
			expected:   [][]int{{1, 2, 3}},
		},
		{
			name:       "multiple duplicates",
			candidates: []int{1, 1, 1, 2, 2},
			target:     4,
			expected:   [][]int{{1, 1, 2}, {2, 2}},
		},
		{
			name:       "large numbers",
			candidates: []int{10, 10, 20, 30},
			target:     50,
			expected:   [][]int{{10, 10, 30}, {20, 30}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := combinationSum2(tt.candidates, tt.target)

			// Sort both result and expected for consistent comparison
			sortCombinations(result)
			sortCombinations(tt.expected)

			assert.Equal(t, tt.expected, result,
				"combinationSum2(%v, %d) should equal %v", tt.candidates, tt.target, tt.expected)
		})
	}
}
