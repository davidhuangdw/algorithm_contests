package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaximizedCapital(t *testing.T) {
	tests := []struct {
		name     string
		k        int
		w        int
		profits  []int
		capital  []int
		expected int
	}{
		{
			name:     "LeetCode Example 1",
			k:        2,
			w:        0,
			profits:  []int{1, 2, 3},
			capital:  []int{0, 1, 1},
			expected: 4,
		},
		{
			name:     "LeetCode Example 2",
			k:        3,
			w:        0,
			profits:  []int{1, 2, 3},
			capital:  []int{0, 1, 2},
			expected: 6,
		},
		{
			name:     "Single project",
			k:        1,
			w:        10,
			profits:  []int{5},
			capital:  []int{5},
			expected: 15,
		},
		{
			name:     "No projects available",
			k:        5,
			w:        10,
			profits:  []int{},
			capital:  []int{},
			expected: 10,
		},
		{
			name:     "All projects too expensive",
			k:        3,
			w:        5,
			profits:  []int{10, 20, 30},
			capital:  []int{10, 15, 20},
			expected: 5,
		},
		{
			name:     "Can only afford some projects",
			k:        2,
			w:        10,
			profits:  []int{5, 10, 15},
			capital:  []int{35, 32, 8},
			expected: 25,
		},
		{
			name:     "Multiple projects with same capital requirement",
			k:        3,
			w:        5,
			profits:  []int{3, 5, 7, 1},
			capital:  []int{5, 5, 5, 3},
			expected: 20,
		},
		{
			name:     "Large profits with small capital requirements",
			k:        2,
			w:        1,
			profits:  []int{100, 200, 300},
			capital:  []int{1, 1, 1},
			expected: 501,
		},
		{
			name:     "k larger than available projects",
			k:        5,
			w:        10,
			profits:  []int{1, 2, 3},
			capital:  []int{5, 6, 7},
			expected: 16,
		},
		{
			name:     "Zero initial capital",
			k:        2,
			w:        0,
			profits:  []int{2, 3, 1},
			capital:  []int{0, 0, 1},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaximizedCapital(tt.k, tt.w, tt.profits, tt.capital)
			assert.Equal(t, tt.expected, result)
		})
	}
}
