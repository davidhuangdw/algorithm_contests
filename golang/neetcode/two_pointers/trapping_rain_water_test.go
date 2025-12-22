package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		// Basic cases
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{2, 0, 2}, 2},
		{[]int{3, 0, 2, 0, 4}, 7},
		{[]int{1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},

		// Edge cases
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 0},
		{[]int{2, 1}, 0},
		{[]int{1, 2, 1}, 0},
		{[]int{5, 4, 3, 2, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 0},
		{[]int{5, 5, 5, 5}, 0},
		{[]int{0, 0, 0, 0}, 0},
		{[]int{1, 0, 0, 0, 1}, 3},
		{[]int{1, 0, 0, 0, 2}, 3},
		{[]int{2, 0, 0, 0, 1}, 3},
		{[]int{3, 0, 0, 0, 2, 0, 0, 0, 4}, 19},
		{[]int{1, 0, 3, 0, 2, 0, 1}, 4},
		{[]int{3, 1, 2, 0, 4, 0, 2, 1, 3}, 12},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := trap(tt.height)
			assert.Equal(t, tt.want, result,
				"trap(%v) = %d, want %d", tt.height, result, tt.want)
		})
	}
}
