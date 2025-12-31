package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinInterval(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		queries   []int
		want      []int
	}{
		{
			name:      "leetcode example 1",
			intervals: [][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}},
			queries:   []int{2, 3, 4, 5},
			want:      []int{3, 3, 1, 4},
		},
		{
			name:      "leetcode example 2",
			intervals: [][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}},
			queries:   []int{2, 19, 5, 22},
			want:      []int{2, -1, 4, 6},
		},
		{
			name:      "single interval",
			intervals: [][]int{{1, 5}},
			queries:   []int{1, 3, 5, 6},
			want:      []int{5, 5, 5, -1},
		},
		{
			name:      "no intervals",
			intervals: [][]int{},
			queries:   []int{1, 2, 3},
			want:      []int{-1, -1, -1},
		},
		{
			name:      "overlapping intervals",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}},
			queries:   []int{2, 3, 4},
			want:      []int{2, 2, 2},
		},
		{
			name:      "exact match intervals",
			intervals: [][]int{{1, 1}, {2, 2}, {3, 3}},
			queries:   []int{1, 2, 3, 4},
			want:      []int{1, 1, 1, -1},
		},
		{
			name:      "large intervals",
			intervals: [][]int{{2, 10}, {5, 9}, {10, 10}, {10, 15}},
			queries:   []int{5, 10, 15},
			want:      []int{5, 1, 6},
		},
		{
			name:      "query outside all intervals",
			intervals: [][]int{{1, 5}, {10, 15}},
			queries:   []int{0, 6, 16},
			want:      []int{-1, -1, -1},
		},
		{
			name:      "multiple intervals same size",
			intervals: [][]int{{1, 3}, {4, 6}, {7, 9}},
			queries:   []int{2, 5, 8},
			want:      []int{3, 3, 3},
		},
		{
			name:      "nested intervals",
			intervals: [][]int{{1, 10}, {2, 9}, {3, 8}, {4, 7}},
			queries:   []int{5, 6},
			want:      []int{4, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minInterval(tt.intervals, tt.queries)
			assert.Equal(t, tt.want, got, "minInterval(%v, %v) = %v, want %v", tt.intervals, tt.queries, got, tt.want)
		})
	}
}
