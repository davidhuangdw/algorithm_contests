package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCoins(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		// LeetCode examples
		{[]int{3, 1, 5, 8}, 167},
		{[]int{1, 5}, 10},

		// Edge cases
		{[]int{}, 0},
		{[]int{5}, 5},
		{[]int{1, 2, 3}, 12},

		// Simple cases
		{[]int{2, 3}, 9},
		{[]int{4, 2, 3}, 40},

		// Complex cases
		{[]int{7, 9, 8, 0, 7, 1, 3, 5, 5, 2, 3}, 1654},
		{[]int{8, 2, 6, 8, 9, 8, 1, 4, 1, 5, 3, 0, 7, 7, 0, 4, 2, 2}, 3446},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := maxCoins(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}
