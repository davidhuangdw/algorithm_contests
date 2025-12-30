package math_geo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name:   "1x1 matrix",
			matrix: [][]int{{1}},
			want:   []int{1},
		},
		{
			name:   "1x3 matrix",
			matrix: [][]int{{1, 2, 3}},
			want:   []int{1, 2, 3},
		},
		{
			name:   "3x1 matrix",
			matrix: [][]int{{1}, {2}, {3}},
			want:   []int{1, 2, 3},
		},
		{
			name:   "2x2 matrix",
			matrix: [][]int{{1, 2}, {3, 4}},
			want:   []int{1, 2, 4, 3},
		},
		{
			name:   "3x3 matrix",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name:   "3x4 matrix",
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name:   "4x3 matrix",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
			want:   []int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8},
		},
		{
			name:   "leetcode example 1",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name:   "leetcode example 2",
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := spiralOrder(tt.matrix)
			assert.Equal(t, tt.want, got, "spiralOrder(%v) = %v, want %v", tt.matrix, got, tt.want)
		})
	}
}