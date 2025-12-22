package sliding_window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		// Basic cases
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, -1}, 1, []int{1, -1}},
		{[]int{9, 11}, 2, []int{11}},
		{[]int{4, -2}, 2, []int{4}},

		// Edge cases
		{[]int{}, 1, []int{}},
		{[]int{1, 2, 3}, 1, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 3, []int{3}},
		{[]int{3, 2, 1}, 3, []int{3}},
		{[]int{1, 2, 3, 4}, 2, []int{2, 3, 4}},
		{[]int{4, 3, 2, 1}, 2, []int{4, 3, 2}},

		// All same elements
		{[]int{5, 5, 5, 5, 5}, 1, []int{5, 5, 5, 5, 5}},
		{[]int{5, 5, 5, 5, 5}, 3, []int{5, 5, 5}},
		{[]int{5, 5, 5, 5, 5}, 5, []int{5}},

		// Increasing sequence
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 4, []int{4, 5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{5}},

		// Decreasing sequence
		{[]int{5, 4, 3, 2, 1}, 1, []int{5, 4, 3, 2, 1}},
		{[]int{5, 4, 3, 2, 1}, 2, []int{5, 4, 3, 2}},
		{[]int{5, 4, 3, 2, 1}, 3, []int{5, 4, 3}},
		{[]int{5, 4, 3, 2, 1}, 4, []int{5, 4}},
		{[]int{5, 4, 3, 2, 1}, 5, []int{5}},

		// Complex patterns
		{[]int{7, 2, 4}, 2, []int{7, 4}},
		{[]int{1, 3, 1, 2, 0, 5}, 3, []int{3, 3, 2, 5}},
		{[]int{8, 7, 6, 9}, 2, []int{8, 7, 9}},
		{[]int{6, 5, 4, 3, 2, 1, 7}, 3, []int{6, 5, 4, 3, 7}},
		{[]int{1, 2, 1, 0, 4, 2, 3}, 4, []int{2, 4, 4, 4}},

		// Large values and negative numbers
		{[]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4, []int{7, 7, 7, 7, 7}},
		{[]int{-1, -2, -3, -4, -5}, 2, []int{-1, -2, -3, -4}},
		{[]int{-1, -3, -2, -4, -5}, 3, []int{-1, -2, -2}},
		{[]int{1000000, 999999, -1000000}, 2, []int{1000000, 999999}},

		// Window size equals array length
		{[]int{1, 2, 3}, 3, []int{3}},
		{[]int{5, 4, 3, 2, 1}, 5, []int{5}},
		{[]int{-1, -2, -3}, 3, []int{-1}},

		// Single element arrays with different k
		{[]int{42}, 1, []int{42}},
		{[]int{0}, 1, []int{0}},
		{[]int{-100}, 1, []int{-100}},

		// Mixed positive and negative
		{[]int{1, -1, 2, -2, 3, -3}, 2, []int{1, 2, 2, 3, 3}},
		{[]int{1, -1, 2, -2, 3, -3}, 3, []int{2, 2, 3, 3}},
		{[]int{1, -1, 2, -2, 3, -3}, 4, []int{2, 3, 3}},

		// Large window sizes relative to array
		{[]int{1, 2, 3, 4}, 4, []int{4}},
		{[]int{1, 2, 3, 4, 5}, 4, []int{4, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, 4, []int{4, 5, 6}},

		// Zero and negative k (should handle gracefully)
		{[]int{1, 2, 3}, 0, []int{}},
		{[]int{}, 0, []int{}},
		{[]int{1, 2, 3}, -1, []int{}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := maxSlidingWindow(tt.nums, tt.k)
			assert.Equal(t, tt.want, result,
				"maxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, result, tt.want)
		})
	}
}
