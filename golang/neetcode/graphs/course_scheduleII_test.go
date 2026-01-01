package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOrder(t *testing.T) {
	tests := []struct {
		name          string
		n             int
		prerequisites [][]int
		expected      []int
	}{
		{
			name:          "LeetCode example 1 - valid ordering",
			n:             2,
			prerequisites: [][]int{{1, 0}},
			expected:      []int{0, 1},
		},
		{
			name:          "LeetCode example 2 - multiple prerequisites",
			n:             4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expected:      []int{0, 1, 2, 3},
		},
		{
			name:          "Cycle detection - no valid ordering",
			n:             2,
			prerequisites: [][]int{{0, 1}, {1, 0}},
			expected:      nil,
		},
		{
			name:          "No prerequisites",
			n:             3,
			prerequisites: [][]int{},
			expected:      []int{0, 1, 2},
		},
		{
			name:          "Single course",
			n:             1,
			prerequisites: [][]int{},
			expected:      []int{0},
		},
		{
			name:          "Complex cycle",
			n:             3,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 0}},
			expected:      nil,
		},
		{
			name:          "Multiple valid orderings - check consistency",
			n:             3,
			prerequisites: [][]int{{2, 0}, {2, 1}},
			expected:      []int{0, 1, 2},
		},
		{
			name:          "Disconnected graph",
			n:             4,
			prerequisites: [][]int{{1, 0}, {3, 2}},
			expected:      []int{0, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findOrder(tt.n, tt.prerequisites)
			assert.Equal(t, tt.expected, result)
		})
	}
}