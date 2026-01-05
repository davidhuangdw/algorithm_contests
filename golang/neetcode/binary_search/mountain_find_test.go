package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockMountainArray implements the MountainArray interface for testing
type MockMountainArray struct {
	arr []int
}

func (m *MockMountainArray) get(index int) int {
	return m.arr[index]
}

func (m *MockMountainArray) length() int {
	return len(m.arr)
}

func TestFindInMountainArray(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "LeetCode Example 1",
			arr:      []int{1, 2, 3, 4, 5, 3, 1},
			target:   3,
			expected: 2,
		},
		{
			name:     "LeetCode Example 2",
			arr:      []int{0, 1, 2, 4, 2, 1},
			target:   3,
			expected: -1,
		},
		{
			name:     "Target on left side",
			arr:      []int{1, 2, 3, 4, 3, 2, 1},
			target:   2,
			expected: 1,
		},
		{
			name:     "Target on right side",
			arr:      []int{1, 2, 3, 4, 3, 2, 1},
			target:   3,
			expected: 2,
		},
		{
			name:     "Target at peak",
			arr:      []int{1, 3, 5, 7, 5, 3, 1},
			target:   7,
			expected: 3,
		},
		{
			name:     "Target not found",
			arr:      []int{1, 2, 3, 4, 3, 2, 1},
			target:   5,
			expected: -1,
		},
		{
			name:     "Small mountain array",
			arr:      []int{1, 2, 1},
			target:   2,
			expected: 1,
		},
		{
			name:     "Single element mountain",
			arr:      []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "Single element not found",
			arr:      []int{5},
			target:   3,
			expected: -1,
		},
		{
			name:     "Symmetric mountain",
			arr:      []int{1, 2, 3, 4, 3, 2, 1},
			target:   4,
			expected: 3,
		},
		{
			name:     "Target appears on both sides, return leftmost",
			arr:      []int{1, 2, 3, 2, 1},
			target:   2,
			expected: 1,
		},
		{
			name:     "Large mountain array",
			arr:      []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 18, 16, 14, 12, 10, 8, 6, 4, 2},
			target:   10,
			expected: 14,
		},
		{
			name:     "Target at beginning",
			arr:      []int{5, 4, 3, 2, 1},
			target:   5,
			expected: 0,
		},
		{
			name:     "Target at end",
			arr:      []int{1, 2, 3, 4, 3, 2, 0},
			target:   0,
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockMountainArray{arr: tt.arr}
			result := findInMountainArray(tt.target, mock)
			assert.Equal(t, tt.expected, result)
		})
	}
}