package arrays_hashing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
		desc string
	}{
		{
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
			desc: "consecutive sequence 1,2,3,4",
		},
		{
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
			desc: "long consecutive sequence 0-8",
		},
		{
			nums: []int{},
			want: 0,
			desc: "empty array",
		},
		{
			nums: []int{1},
			want: 1,
			desc: "single element",
		},
		{
			nums: []int{1, 3, 5, 7, 9},
			want: 1,
			desc: "all elements isolated",
		},
		{
			nums: []int{1, 2, 0, 1},
			want: 3,
			desc: "with duplicates, sequence 0,1,2",
		},
		{
			nums: []int{-7, -1, 3, -9, -4, -7, 2, 3, 3, 7, -2, 3, 4, 9, 4, 9, 8, 1, 0, -4, -1, 9},
			want: 7,
			desc: "complex case with negatives and duplicates",
		},
		{
			nums: []int{-1, 0, 1},
			want: 3,
			desc: "sequence with negative numbers",
		},
		{
			nums: []int{-3, -2, -1, 0, 1},
			want: 5,
			desc: "long sequence with negative numbers",
		},
		{
			nums: []int{1, 1, 1, 1},
			want: 1,
			desc: "all duplicates",
		},
		{
			nums: []int{10, 20, 30, 40},
			want: 1,
			desc: "large gaps between numbers",
		},
		{
			nums: []int{5, 4, 3, 2, 1},
			want: 5,
			desc: "reverse sorted consecutive",
		},
		{
			nums: []int{7, 6, 5, 4, 3, 2, 1},
			want: 7,
			desc: "long reverse sorted consecutive",
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			want: 7,
			desc: "long sorted consecutive",
		},
		{
			nums: []int{1, 3, 5, 2, 4},
			want: 5,
			desc: "unsorted consecutive sequence",
		},
		{
			nums: []int{-10, -9, -8, -7, -6},
			want: 5,
			desc: "negative consecutive sequence",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.desc), func(t *testing.T) {
			got := longestConsecutive(tt.nums)
			assert.Equal(t, tt.want, got, "longestConsecutive(%v) should return %d", tt.nums, tt.want)
		})
	}
}
