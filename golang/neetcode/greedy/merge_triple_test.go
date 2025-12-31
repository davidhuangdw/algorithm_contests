package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTriplets(t *testing.T) {
	tests := []struct {
		name     string
		triplets [][]int
		target   []int
		want     bool
	}{
		{
			name:     "leetcode example 1",
			triplets: [][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}},
			target:   []int{2, 7, 5},
			want:     true,
		},
		{
			name:     "leetcode example 2",
			triplets: [][]int{{3, 4, 5}, {4, 5, 6}},
			target:   []int{3, 2, 5},
			want:     false,
		},
		{
			name:     "exact match",
			triplets: [][]int{{1, 2, 3}},
			target:   []int{1, 2, 3},
			want:     true,
		},
		{
			name:     "multiple triplets needed",
			triplets: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			target:   []int{1, 1, 1},
			want:     true,
		},
		{
			name:     "one dimension too large",
			triplets: [][]int{{2, 5, 3}, {1, 8, 4}},
			target:   []int{2, 7, 5},
			want:     false,
		},
		{
			name:     "empty triplets",
			triplets: [][]int{},
			target:   []int{1, 2, 3},
			want:     false,
		},
		{
			name:     "single triplet with all dimensions matching",
			triplets: [][]int{{5, 5, 5}},
			target:   []int{5, 5, 5},
			want:     true,
		},
		{
			name:     "triplet with one dimension exceeding target",
			triplets: [][]int{{1, 2, 3}, {4, 5, 6}},
			target:   []int{3, 4, 5},
			want:     false,
		},
		{
			name:     "combination of triplets",
			triplets: [][]int{{2, 3, 4}, {1, 5, 3}, {3, 2, 5}},
			target:   []int{3, 5, 5},
			want:     true,
		},
		{
			name:     "target with zeros",
			triplets: [][]int{{0, 0, 0}, {1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			target:   []int{1, 1, 0},
			want:     true,
		},
		{
			name:     "all triplets have one dimension too large",
			triplets: [][]int{{2, 1, 1}, {1, 2, 1}, {1, 1, 2}},
			target:   []int{1, 1, 1},
			want:     false,
		},
		{
			name:     "large numbers",
			triplets: [][]int{{100, 200, 300}, {150, 250, 350}, {200, 300, 400}},
			target:   []int{200, 300, 400},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTriplets(tt.triplets, tt.target)
			assert.Equal(t, tt.want, got, "mergeTriplets(%v, %v) = %v, want %v", tt.triplets, tt.target, got, tt.want)
		})
	}
}