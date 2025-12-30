package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "empty intervals",
			intervals: [][]int{},
			want:      [][]int{},
		},
		{
			name:      "single interval",
			intervals: [][]int{{1, 3}},
			want:      [][]int{{1, 3}},
		},
		{
			name:      "non-overlapping intervals",
			intervals: [][]int{{1, 3}, {4, 6}, {7, 9}},
			want:      [][]int{{1, 3}, {4, 6}, {7, 9}},
		},
		{
			name:      "overlapping intervals",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "leetcode example 1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "leetcode example 2",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "adjacent intervals",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}},
			want:      [][]int{{1, 4}},
		},
		{
			name:      "nested intervals",
			intervals: [][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}},
			want:      [][]int{{1, 10}},
		},
		{
			name:      "unsorted intervals",
			intervals: [][]int{{8, 10}, {1, 3}, {15, 18}, {2, 6}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "multiple merges needed",
			intervals: [][]int{{1, 4}, {0, 2}, {3, 5}},
			want:      [][]int{{0, 5}},
		},
		{
			name:      "large intervals",
			intervals: [][]int{{1, 100}, {10, 20}, {30, 40}},
			want:      [][]int{{1, 100}},
		},
		{
			name:      "negative numbers",
			intervals: [][]int{{-5, -3}, {-4, -2}, {-1, 1}},
			want:      [][]int{{-5, -2}, {-1, 1}},
		},
		{
			name:      "mixed positive and negative",
			intervals: [][]int{{-2, 0}, {1, 3}, {0, 2}},
			want:      [][]int{{-2, 3}},
		},
		{
			name:      "complex overlapping pattern",
			intervals: [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			want:      [][]int{{1, 10}},
		},
		{
			name:      "single point intervals",
			intervals: [][]int{{1, 1}, {2, 2}, {3, 3}},
			want:      [][]int{{1, 1}, {2, 2}, {3, 3}},
		},
		{
			name:      "sparse intervals with gaps",
			intervals: [][]int{{1, 2}, {5, 6}, {10, 11}, {15, 16}},
			want:      [][]int{{1, 2}, {5, 6}, {10, 11}, {15, 16}},
		},
		{
			name:      "duplicate intervals",
			intervals: [][]int{{1, 3}, {1, 3}, {1, 3}},
			want:      [][]int{{1, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.intervals)
			assert.Equal(t, tt.want, got, "merge(%v) = %v, want %v", tt.intervals, got, tt.want)
		})
	}
}
