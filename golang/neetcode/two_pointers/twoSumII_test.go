package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		want    []int
	}{
		// Basic cases
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, 9, []int{4, 5}},
		{[]int{1, 2}, 3, []int{1, 2}},
		{[]int{1, 3, 4, 5, 7, 11}, 9, []int{3, 4}},
		{[]int{0, 0, 3, 4}, 0, []int{1, 2}},
		{[]int{-5, -3, -1, 0, 2}, -3, []int{1, 5}},

		// No solution cases
		{[]int{1, 2, 3, 4, 5}, 10, nil},
		{[]int{1}, 2, nil},
		{[]int{}, 0, nil},
		{[]int{1, 2, 3}, 10, nil},

		// Edge cases
		{[]int{}, 5, nil},
		{[]int{1}, 1, nil},
		{[]int{1, 1, 2, 3}, 2, []int{1, 2}},
		{[]int{-3, -1, 0, 2}, -3, []int{1, 3}},
		{[]int{1000, 2000, 3000}, 5000, []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := twoSum(tt.numbers, tt.target)
			assert.Equal(t, tt.want, result,
				"twoSum(%v, %d) = %v, want %v", tt.numbers, tt.target, result, tt.want)
		})
	}
}
