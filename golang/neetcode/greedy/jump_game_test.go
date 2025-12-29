package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "single element",
			nums: []int{0},
			want: true,
		},
		{
			name: "can reach end",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "cannot reach end",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "leetcode example 1",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "leetcode example 2",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "all zeros except first",
			nums: []int{1, 0, 0, 0, 0},
			want: false,
		},
		{
			name: "large jump at start",
			nums: []int{5, 0, 0, 0, 0},
			want: true,
		},
		{
			name: "exactly reach end",
			nums: []int{1, 1, 1, 1, 1},
			want: true,
		},
		{
			name: "zero at start",
			nums: []int{0, 1, 2, 3},
			want: false,
		},
		{
			name: "large numbers",
			nums: []int{10, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: true,
		},
		{
			name: "alternating pattern",
			nums: []int{2, 0, 2, 0, 2, 0},
			want: true,
		},
		{
			name: "complex pattern reachable",
			nums: []int{2, 0, 1, 0, 1, 0},
			want: false,
		},
		{
			name: "edge case with one element",
			nums: []int{1},
			want: true,
		},
		{
			name: "empty array",
			nums: []int{},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canJump(tt.nums)
			assert.Equal(t, tt.want, got, "canJump(%v) = %v, want %v", tt.nums, got, tt.want)
		})
	}
}
