package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			name:        "empty intervals",
			intervals:   [][]int{},
			newInterval: []int{1, 3},
			want:        [][]int{{1, 3}},
		},
		{
			name:        "insert at beginning",
			intervals:   [][]int{{2, 5}, {6, 9}},
			newInterval: []int{0, 1},
			want:        [][]int{{0, 1}, {2, 5}, {6, 9}},
		},
		{
			name:        "insert at end",
			intervals:   [][]int{{1, 3}, {4, 6}},
			newInterval: []int{7, 9},
			want:        [][]int{{1, 3}, {4, 6}, {7, 9}},
		},
		{
			name:        "merge with one interval",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "merge with multiple intervals",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 9},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:        "leetcode example 1",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "leetcode example 2",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 9},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:        "new interval covers all",
			intervals:   [][]int{{1, 2}, {3, 4}, {5, 6}},
			newInterval: []int{0, 7},
			want:        [][]int{{0, 7}},
		},
		{
			name:        "new interval inside existing",
			intervals:   [][]int{{1, 10}},
			newInterval: []int{3, 5},
			want:        [][]int{{1, 10}},
		},
		{
			name:        "overlap with first interval",
			intervals:   [][]int{{1, 5}, {6, 8}},
			newInterval: []int{0, 3},
			want:        [][]int{{0, 5}, {6, 8}},
		},
		{
			name:        "overlap with last interval",
			intervals:   [][]int{{1, 3}, {4, 6}},
			newInterval: []int{5, 9},
			want:        [][]int{{1, 3}, {4, 9}},
		},
		{
			name:        "exact match with existing",
			intervals:   [][]int{{1, 3}, {4, 6}, {7, 9}},
			newInterval: []int{4, 6},
			want:        [][]int{{1, 3}, {4, 6}, {7, 9}},
		},
		{
			name:        "single interval merge",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 3},
			want:        [][]int{{1, 5}},
		},
		{
			name:        "adjacent intervals",
			intervals:   [][]int{{1, 2}, {3, 4}, {5, 6}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insert(tt.intervals, tt.newInterval)
			assert.Equal(t, tt.want, got, "insert(%v, %v) = %v, want %v", tt.intervals, tt.newInterval, got, tt.want)
		})
	}
}
