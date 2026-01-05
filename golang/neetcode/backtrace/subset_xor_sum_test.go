package backtrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetXORSum(t *testing.T) {
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
			nums: []int{1},
			want: 1,
		},
		{
			name: "two elements",
			nums: []int{1, 3},
			want: 6,
		},
		{
			name: "three elements",
			nums: []int{5, 1, 6},
			want: 28,
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0},
			want: 0,
		},
		{
			name: "mixed positive numbers",
			nums: []int{2, 4, 8},
			want: 56,
		},
		{
			name: "with negative numbers",
			nums: []int{-1, 2},
			want: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subsetXORSum(tt.nums)
			assert.Equal(t, tt.want, result,
				"subsetXORSum(%v) = %d, want %d", tt.nums, result, tt.want)
		})
	}
}