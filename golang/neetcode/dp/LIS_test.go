package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty array",
			nums: []int{},
			want: 0,
		},
		{
			name: "single element",
			nums: []int{5},
			want: 1,
		},
		{
			name: "strictly increasing",
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "strictly decreasing",
			nums: []int{5, 4, 3, 2, 1},
			want: 1,
		},
		{
			name: "classic example",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			name: "all same elements",
			nums: []int{7, 7, 7, 7, 7},
			want: 1,
		},
		{
			name: "mixed with duplicates",
			nums: []int{4, 10, 4, 3, 8, 9},
			want: 3,
		},
		{
			name: "leetcode example 1",
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			name: "leetcode example 2",
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},
		{
			name: "negative numbers",
			nums: []int{-1, -2, -3, -4},
			want: 1,
		},
		{
			name: "mixed positive and negative",
			nums: []int{-2, -1, 0, 1, 2},
			want: 5,
		},
		{
			name: "large numbers",
			nums: []int{100, 200, 300, 400},
			want: 4,
		},
		{
			name: "complex pattern",
			nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLIS(tt.nums)
			assert.Equal(t, tt.want, got, "lengthOfLIS(%v) = %v, want %v", tt.nums, got, tt.want)
		})
	}
}
