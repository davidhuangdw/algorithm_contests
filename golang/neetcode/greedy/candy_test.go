package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCandy(t *testing.T) {
	tests := []struct {
		name     string
		ratings  []int
		expected int
	}{
		{
			name:     "LeetCode example 1",
			ratings:  []int{1, 0, 2},
			expected: 5,
		},
		{
			name:     "LeetCode example 2",
			ratings:  []int{1, 2, 2},
			expected: 4,
		},
		{
			name:     "single child",
			ratings:  []int{5},
			expected: 1,
		},
		{
			name:     "all same ratings",
			ratings:  []int{3, 3, 3, 3},
			expected: 4,
		},
		{
			name:     "strictly increasing",
			ratings:  []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "strictly decreasing",
			ratings:  []int{5, 4, 3, 2, 1},
			expected: 15,
		},
		{
			name:     "mountain shape",
			ratings:  []int{1, 3, 5, 3, 1},
			expected: 9,
		},
		{
			name:     "valley shape",
			ratings:  []int{5, 3, 1, 3, 5},
			expected: 11,
		},
		{
			name:     "alternating pattern",
			ratings:  []int{1, 2, 1, 2, 1},
			expected: 7,
		},
		{
			name:     "plateau in middle",
			ratings:  []int{1, 2, 2, 1},
			expected: 6,
		},
		{
			name:     "complex pattern 1",
			ratings:  []int{1, 3, 2, 1},
			expected: 7,
		},
		{
			name:     "complex pattern 2",
			ratings:  []int{1, 2, 3, 1, 0},
			expected: 9,
		},
		{
			name:     "large ratings difference",
			ratings:  []int{10, 1, 10},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy(tt.ratings)
			assert.Equal(t, tt.expected, result, "candy(%v)", tt.ratings)
		})
	}
}
