package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindClosestElements(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		x        int
		expected []int
	}{
		{
			name:     "leetcode example 1",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        3,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "leetcode example 2",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        -1,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "x is exactly in middle",
			arr:      []int{1, 2, 3, 4, 5},
			k:        3,
			x:        3,
			expected: []int{2, 3, 4},
		},
		{
			name:     "x at beginning",
			arr:      []int{1, 2, 3, 4, 5},
			k:        3,
			x:        1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "x at end",
			arr:      []int{1, 2, 3, 4, 5},
			k:        3,
			x:        5,
			expected: []int{3, 4, 5},
		},
		{
			name:     "k equals array length",
			arr:      []int{1, 2, 3},
			k:        3,
			x:        2,
			expected: []int{1, 2, 3},
		},
		{
			name:     "k larger than array length",
			arr:      []int{1, 2, 3},
			k:        5,
			x:        2,
			expected: []int{1, 2, 3},
		},
		{
			name:     "empty array",
			arr:      []int{},
			k:        3,
			x:        5,
			expected: []int{},
		},
		{
			name:     "negative numbers",
			arr:      []int{-5, -3, -1, 1, 3, 5},
			k:        4,
			x:        0,
			expected: []int{-3, -1, 1, 3},
		},
		{
			name:     "duplicate values",
			arr:      []int{1, 1, 2, 2, 3, 3},
			k:        4,
			x:        2,
			expected: []int{1, 1, 2, 2},
		},
		{
			name:     "large numbers",
			arr:      []int{100, 200, 300, 400, 500},
			k:        3,
			x:        350,
			expected: []int{200, 300, 400},
		},
		{
			name:     "tie in distance",
			arr:      []int{1, 3, 5, 7, 9},
			k:        2,
			x:        4,
			expected: []int{3, 5}, // both are distance 1 from 4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findClosestElements(tt.arr, tt.k, tt.x)
			assert.Equal(t, tt.expected, result, "arr=%v, k=%d, x=%d", tt.arr, tt.k, tt.x)
		})
	}
}
