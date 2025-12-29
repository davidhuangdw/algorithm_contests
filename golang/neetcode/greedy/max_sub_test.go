package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "single positive number",
			nums: []int{5},
			want: 5,
		},
		{
			name: "single negative number",
			nums: []int{-5},
			want: -5,
		},
		{
			name: "all positive numbers",
			nums: []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "all negative numbers",
			nums: []int{-1, -2, -3, -4, -5},
			want: -1,
		},
		{
			name: "classic example",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "leetcode example 1",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "leetcode example 2",
			nums: []int{1},
			want: 1,
		},
		{
			name: "leetcode example 3",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			name: "mixed with zeros",
			nums: []int{0, -1, 0, 2, 0, -3, 0},
			want: 2,
		},
		{
			name: "alternating pattern",
			nums: []int{1, -2, 3, -4, 5},
			want: 5,
		},
		{
			name: "large positive sum",
			nums: []int{10, 20, 30, 40},
			want: 100,
		},
		{
			name: "edge case with one negative",
			nums: []int{-1, -2, -3, 0},
			want: 0,
		},
		{
			name: "complex pattern",
			nums: []int{8, -19, 5, -4, 20},
			want: 21,
		},
		{
			name: "empty array",
			nums: []int{},
			want: -1000000000, // -INF value
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubArray(tt.nums)
			assert.Equal(t, tt.want, got, "maxSubArray(%v) = %v, want %v", tt.nums, got, tt.want)
		})
	}
}
