package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum4(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Basic case with [1,2,3] and target 4",
			nums:   []int{1, 2, 3},
			target: 4,
			want:   7, // 1+1+1+1, 1+1+2, 1+2+1, 2+1+1, 2+2, 1+3, 3+1
		},
		{
			name:   "Single number [9] with target 3",
			nums:   []int{9},
			target: 3,
			want:   0, // No combination possible
		},
		{
			name:   "Target 0",
			nums:   []int{1, 2, 3},
			target: 0,
			want:   1, // Empty combination
		},
		{
			name:   "Only 1s",
			nums:   []int{1},
			target: 5,
			want:   1, // Only one way: 1+1+1+1+1
		},
		{
			name:   "Numbers larger than target",
			nums:   []int{4, 5, 6},
			target: 3,
			want:   0, // No combination possible
		},
		{
			name:   "Mixed numbers [1,2] with target 3",
			nums:   []int{1, 2},
			target: 3,
			want:   3, // 1+1+1, 1+2, 2+1
		},
		{
			name:   "Larger target with [1,2,3]",
			nums:   []int{1, 2, 3},
			target: 5,
			want:   13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum4(tt.nums, tt.target)
			assert.Equal(t, tt.want, got, "combinationSum4(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
		})
	}
}
