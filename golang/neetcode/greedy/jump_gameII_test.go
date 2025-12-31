package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "leetcode example 1",
			nums: []int{2, 3, 1, 1, 4},
			want: 2,
		},
		{
			name: "leetcode example 2",
			nums: []int{2, 3, 0, 1, 4},
			want: 2,
		},
		{
			name: "single element",
			nums: []int{0},
			want: 0,
		},
		{
			name: "already at end",
			nums: []int{1},
			want: 0,
		},
		{
			name: "minimum jumps needed",
			nums: []int{1, 1, 1, 1, 1},
			want: 4,
		},
		{
			name: "large jump at start",
			nums: []int{5, 0, 0, 0, 0},
			want: 1,
		},
		{
			name: "optimal path not greedy",
			nums: []int{2, 3, 1, 1, 4, 1},
			want: 3,
		},
		{
			name: "complex pattern",
			nums: []int{1, 2, 3, 4, 5},
			want: 3,
		},
		{
			name: "edge case with zeros",
			nums: []int{3, 2, 1, 0, 4},
			want: 2,
		},
		{
			name: "all ones",
			nums: []int{1, 1, 1, 1},
			want: 3,
		},
		{
			name: "large numbers",
			nums: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jump(tt.nums)
			assert.Equal(t, tt.want, got, "jump(%v) = %d, want %d", tt.nums, got, tt.want)
		})
	}
}
