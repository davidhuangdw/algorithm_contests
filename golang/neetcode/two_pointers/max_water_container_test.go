package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		// Basic cases
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
		{[]int{1, 2, 4, 3}, 4},

		// Edge cases
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 1},
		{[]int{5, 5, 5, 5}, 15},
		{[]int{0, 0, 0, 0}, 0},
		{[]int{1, 0, 1}, 2},
		{[]int{0, 1, 0}, 0},

		// Complex cases
		{[]int{1, 3, 2, 5, 25, 24, 5}, 24},
		{[]int{2, 3, 10, 5, 7, 8, 9}, 36},
		{[]int{1, 8, 6, 2, 5, 4, 8, 25, 7}, 49},
		{[]int{10, 1, 1, 1, 1, 1, 1, 1, 10}, 80},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 25},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 25},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := maxArea(tt.height)
			assert.Equal(t, tt.want, result,
				"maxArea(%v) = %d, want %d", tt.height, result, tt.want)
		})
	}
}
