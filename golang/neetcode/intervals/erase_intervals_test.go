package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name:      "empty intervals",
			intervals: [][]int{},
			want:      0,
		},
		{
			name:      "single interval",
			intervals: [][]int{{1, 2}},
			want:      0,
		},
		{
			name:      "non-overlapping intervals",
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			want:      0,
		},
		{
			name:      "overlapping intervals - remove one",
			intervals: [][]int{{1, 3}, {2, 4}, {3, 5}},
			want:      1,
		},
		{
			name:      "leetcode example 1",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			want:      1,
		},
		{
			name:      "leetcode example 2",
			intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
			want:      2,
		},
		{
			name:      "nested intervals",
			intervals: [][]int{{1, 5}, {2, 3}, {3, 4}},
			want:      1,
		},
		{
			name:      "adjacent intervals",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}},
			want:      0,
		},
		{
			name:      "complex overlapping pattern",
			intervals: [][]int{{0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}},
			want:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := eraseOverlapIntervals(tt.intervals)
			assert.Equal(t, tt.want, got, "eraseOverlapIntervals(%v) = %v, want %v", tt.intervals, got, tt.want)
		})
	}
}
