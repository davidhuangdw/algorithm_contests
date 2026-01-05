package backtrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "empty array",
			nums: []int{},
			want: [][]int{{}},
		},
		{
			name: "single element",
			nums: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "all unique elements",
			nums: []int{1, 2},
			want: [][]int{{1, 2}, {2, 1}},
		},
		{
			name: "with duplicates",
			nums: []int{1, 1, 2},
			want: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			name: "all duplicates",
			nums: []int{1, 1, 1},
			want: [][]int{{1, 1, 1}},
		},
		{
			name: "mixed duplicates",
			nums: []int{1, 2, 1},
			want: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			name: "duplicated",
			nums: []int{2, 3, 1, 3},
			want: [][]int{{1, 2, 3, 3}, {1, 3, 2, 3}, {1, 3, 3, 2}, {2, 1, 3, 3}, {2, 3, 1, 3}, {2, 3, 3, 1}, {3, 1, 2, 3}, {3, 1, 3, 2}, {3, 2, 1, 3}, {3, 2, 3, 1}, {3, 3, 1, 2}, {3, 3, 2, 1}},
		},
		{
			name: "negative numbers",
			nums: []int{-1, 0, 1},
			want: [][]int{{-1, 0, 1}, {-1, 1, 0}, {0, -1, 1}, {0, 1, -1}, {1, -1, 0}, {1, 0, -1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := permuteUnique(tt.nums)

			// Sort both result and expected for consistent comparison
			sortPermutations(result)
			sortPermutations(tt.want)

			assert.Equal(t, tt.want, result,
				"permuteUnique(%v) = %v, want %v", tt.nums, result, tt.want)
		})
	}
}
