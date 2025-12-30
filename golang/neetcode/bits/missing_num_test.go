package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "missing 0",
			nums: []int{1, 2, 3},
			want: 0,
		},
		{
			name: "missing 1",
			nums: []int{0, 2, 3},
			want: 1,
		},
		{
			name: "missing 2",
			nums: []int{0, 1, 3},
			want: 2,
		},
		{
			name: "missing 3",
			nums: []int{0, 1, 2},
			want: 3,
		},
		{
			name: "leetcode example 1",
			nums: []int{3, 0, 1},
			want: 2,
		},
		{
			name: "leetcode example 2",
			nums: []int{0, 1},
			want: 2,
		},
		{
			name: "leetcode example 3",
			nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},
		{
			name: "single element missing 0",
			nums: []int{1},
			want: 0,
		},
		{
			name: "single element missing 1",
			nums: []int{0},
			want: 1,
		},
		{
			name: "empty array",
			nums: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := missingNumber(tt.nums)
			assert.Equal(t, tt.want, got, "missingNumber(%v) = %v, want %v", tt.nums, got, tt.want)
		})
	}
}