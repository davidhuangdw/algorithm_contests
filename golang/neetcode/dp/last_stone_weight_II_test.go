package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastStoneWeightII(t *testing.T) {
	tests := []struct {
		name     string
		stones   []int
		expected int
	}{
		{
			name:     "empty array",
			stones:   []int{},
			expected: 0,
		},
		{
			name:     "single stone",
			stones:   []int{5},
			expected: 5,
		},
		{
			name:     "two stones same weight",
			stones:   []int{2, 2},
			expected: 0,
		},
		{
			name:     "complex case",
			stones:   []int{9, 1, 2, 3, 5},
			expected: 0,
		},
		{
			name:     "two stones different weight",
			stones:   []int{2, 7},
			expected: 5,
		},
		{
			name:     "three stones - leetcode example 1",
			stones:   []int{2, 7, 4},
			expected: 1,
		},
		{
			name:     "three stones - all same",
			stones:   []int{3, 3, 3},
			expected: 3,
		},
		{
			name:     "leetcode example 2",
			stones:   []int{31, 26, 33, 21, 40},
			expected: 5,
		},
		{
			name:     "four stones",
			stones:   []int{1, 1, 4, 2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lastStoneWeightII(tt.stones)
			assert.Equal(t, tt.expected, result)
		})
	}
}
