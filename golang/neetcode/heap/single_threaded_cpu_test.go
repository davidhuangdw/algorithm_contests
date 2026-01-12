package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrder(t *testing.T) {
	tests := []struct {
		name     string
		tasks    [][]int
		expected []int
	}{
		{
			name:     "LeetCode Example 1",
			tasks:    [][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}},
			expected: []int{0, 2, 3, 1},
		},
		{
			name:     "LeetCode Example 2",
			tasks:    [][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}},
			expected: []int{4, 3, 2, 0, 1},
		},
		{
			name:     "Single task",
			tasks:    [][]int{{1, 5}},
			expected: []int{0},
		},
		{
			name:     "Tasks with same start time, different processing times",
			tasks:    [][]int{{0, 3}, {0, 1}, {0, 2}},
			expected: []int{1, 2, 0},
		},
		{
			name:     "Tasks with different start times",
			tasks:    [][]int{{5, 2}, {7, 2}, {1, 3}, {2, 4}, {3, 1}},
			expected: []int{2, 4, 0, 1, 3},
		},
		{
			name:     "Tasks with same processing time, different start times",
			tasks:    [][]int{{3, 2}, {1, 2}, {5, 2}, {2, 2}},
			expected: []int{1, 0, 2, 3},
		},
		{
			name:     "Mixed case with ties in processing time",
			tasks:    [][]int{{1, 5}, {2, 5}, {3, 5}, {4, 5}},
			expected: []int{0, 1, 2, 3},
		},
		{
			name:     "Gap between tasks",
			tasks:    [][]int{{1, 2}, {5, 1}, {10, 3}},
			expected: []int{0, 1, 2},
		},
		{
			name:     "Complex overlapping tasks",
			tasks:    [][]int{{1, 4}, {2, 3}, {3, 2}, {4, 1}, {5, 5}},
			expected: []int{0, 3, 2, 1, 4},
		},
		{
			name:     "Tasks with zero processing time",
			tasks:    [][]int{{1, 0}, {2, 0}, {3, 0}},
			expected: []int{0, 1, 2},
		},
		{
			name:     "Large processing times",
			tasks:    [][]int{{1, 100}, {2, 50}, {3, 25}, {4, 10}},
			expected: []int{0, 3, 2, 1},
		},
		{
			name:     "Empty tasks",
			tasks:    [][]int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getOrder(tt.tasks)
			assert.Equal(t, tt.expected, result)
		})
	}
}
