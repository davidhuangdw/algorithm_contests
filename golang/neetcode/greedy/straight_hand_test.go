package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNStraightHand(t *testing.T) {
	tests := []struct {
		name      string
		hand      []int
		groupSize int
		want      bool
	}{
		{
			name:      "leetcode example 1",
			hand:      []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
			groupSize: 3,
			want:      true,
		},
		{
			name:      "leetcode example 2",
			hand:      []int{1, 2, 3, 4, 5},
			groupSize: 4,
			want:      false,
		},
		{
			name:      "single group size 1",
			hand:      []int{1},
			groupSize: 1,
			want:      true,
		},
		{
			name:      "multiple groups size 2",
			hand:      []int{1, 2, 3, 4},
			groupSize: 2,
			want:      true,
		},
		{
			name:      "consecutive groups",
			hand:      []int{1, 2, 3, 4, 5, 6},
			groupSize: 3,
			want:      true,
		},
		{
			name:      "non-consecutive cards",
			hand:      []int{1, 2, 4, 5},
			groupSize: 4,
			want:      false,
		},
		{
			name:      "duplicate cards insufficient",
			hand:      []int{1, 1, 2, 2, 3, 3},
			groupSize: 3,
			want:      true,
		},
		{
			name:      "empty hand",
			hand:      []int{},
			groupSize: 1,
			want:      true,
		},
		{
			name:      "hand size not divisible by group size",
			hand:      []int{1, 2, 3},
			groupSize: 2,
			want:      false,
		},
		{
			name:      "large group size",
			hand:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			groupSize: 5,
			want:      true,
		},
		{
			name:      "gaps in sequence",
			hand:      []int{1, 3, 5, 7},
			groupSize: 2,
			want:      false,
		},
		{
			name:      "multiple same cards",
			hand:      []int{1, 1, 2, 2, 3, 3, 4, 4},
			groupSize: 2,
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNStraightHand(tt.hand, tt.groupSize)
			assert.Equal(t, tt.want, got, "isNStraightHand(%v, %d) = %v, want %v", tt.hand, tt.groupSize, got, tt.want)
		})
	}
}
