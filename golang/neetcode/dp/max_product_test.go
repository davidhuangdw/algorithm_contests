package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProduct(t *testing.T) {
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
			name: "two positive numbers",
			nums: []int{2, 3},
			want: 6,
		},
		{
			name: "two negative numbers",
			nums: []int{-2, -3},
			want: 6,
		},
		{
			name: "mixed positive and negative",
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			name: "classic example with zero",
			nums: []int{2, 3, -2, 4, -1},
			want: 48,
		},
		{
			name: "all negative numbers",
			nums: []int{-2, -3, -4},
			want: 12,
		},
		{
			name: "with zeros",
			nums: []int{-2, 0, -1},
			want: 0,
		},
		{
			name: "multiple zeros",
			nums: []int{0, 0, 0, 2, 3},
			want: 6,
		},
		{
			name: "large positive numbers",
			nums: []int{10, 20, 30},
			want: 6000,
		},
		{
			name: "edge case with one",
			nums: []int{-2, 1, -3},
			want: 6,
		},
		{
			name: "leetcode example 1",
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			name: "leetcode example 2",
			nums: []int{-2, 0, -1},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProduct(tt.nums)
			assert.Equal(t, tt.want, got, "maxProduct(%v) = %v, want %v", tt.nums, got, tt.want)
		})
	}
}
