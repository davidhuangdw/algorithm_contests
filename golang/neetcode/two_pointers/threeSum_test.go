package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		// Basic cases
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			[]int{0, 1, 1},
			[][]int{},
		},
		{
			[]int{0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{
			[]int{-2, 0, 1, 1, 2},
			[][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			[]int{-1, 0, 1, 0},
			[][]int{{-1, 0, 1}},
		},
		{
			[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			[][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		},

		// Edge cases
		{
			[]int{},
			[][]int{},
		},
		{
			[]int{0},
			[][]int{},
		},
		{
			[]int{1, 2, 3},
			[][]int{},
		},
		{
			[]int{0, 0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{
			[]int{-10, -5, -5, 0, 5, 10, 15},
			[][]int{{-10, -5, 15}, {-10, 0, 10}, {-5, -5, 10}, {-5, 0, 5}},
		},
		{
			[]int{-1, -1, -1, 2},
			[][]int{{-1, -1, 2}},
		},
		{
			[]int{1, 2, 3, 4},
			[][]int{},
		},
		{
			[]int{-2, 1, 1},
			[][]int{{-2, 1, 1}},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := threeSum(tt.nums)
			assert.ElementsMatch(t, tt.want, result,
				"threeSum(%v) = %v, want %v", tt.nums, result, tt.want)
		})
	}
}
