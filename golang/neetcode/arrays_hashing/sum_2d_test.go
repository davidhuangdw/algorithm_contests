package arrays_hashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumMatrix(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		queries  [][]int // [i1, j1, i2, j2]
		expected []int
	}{
		{
			name: "basic 3x3 matrix",
			matrix: [][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			},
			queries: [][]int{
				{2, 1, 4, 3}, // 8
				{1, 1, 2, 2}, // 11
				{1, 2, 2, 4}, // 12
			},
			expected: []int{8, 11, 12},
		},
		{
			name: "single element matrix",
			matrix: [][]int{
				{5},
			},
			queries: [][]int{
				{0, 0, 0, 0}, // 5
			},
			expected: []int{5},
		},
		{
			name: "1D row matrix",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
			},
			queries: [][]int{
				{0, 0, 0, 0}, // 1
				{0, 1, 0, 3}, // 9
				{0, 0, 0, 4}, // 15
			},
			expected: []int{1, 9, 15},
		},
		{
			name: "1D column matrix",
			matrix: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
			queries: [][]int{
				{0, 0, 0, 0}, // 1
				{1, 0, 3, 0}, // 9
				{0, 0, 4, 0}, // 15
			},
			expected: []int{1, 9, 15},
		},
		{
			name: "all zeros matrix",
			matrix: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			queries: [][]int{
				{0, 0, 2, 2}, // 0
				{1, 1, 1, 1}, // 0
			},
			expected: []int{0, 0},
		},
		{
			name: "negative numbers",
			matrix: [][]int{
				{-1, -2, -3},
				{-4, -5, -6},
			},
			queries: [][]int{
				{0, 0, 0, 0}, // -1
				{0, 0, 1, 2}, // -21
				{1, 1, 1, 2}, // -11
			},
			expected: []int{-1, -21, -11},
		},
		{
			name: "large values",
			matrix: [][]int{
				{1000, 2000},
				{3000, 4000},
			},
			queries: [][]int{
				{0, 0, 1, 1}, // 10000
				{0, 0, 0, 0}, // 1000
			},
			expected: []int{10000, 1000},
		},
		{
			name: "single row query",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			queries: [][]int{
				{1, 0, 1, 2}, // 15
				{2, 1, 2, 2}, // 17
			},
			expected: []int{15, 17},
		},
		{
			name: "single column query",
			matrix: [][]int{
				{1, 1, 1},
				{2, 2, 2},
				{3, 3, 3},
			},
			queries: [][]int{
				{0, 1, 2, 1}, // 6
				{1, 0, 2, 0}, // 5
			},
			expected: []int{6, 5},
		},
		{
			name: "same point query",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			queries: [][]int{
				{0, 0, 0, 0}, // 1
				{1, 1, 1, 1}, // 4
				{0, 1, 0, 1}, // 2
				{1, 0, 1, 0}, // 3
			},
			expected: []int{1, 4, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := Constructor(tt.matrix)
			for i, query := range tt.queries {
				result := nm.SumRegion(query[0], query[1], query[2], query[3])
				assert.Equal(t, tt.expected[i], result, "query %d failed: %v", i, query)
			}
		})
	}
}