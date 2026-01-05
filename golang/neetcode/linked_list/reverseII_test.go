package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		left     int
		right    int
		expected []int
	}{
		{
			name:     "LeetCode Example 1",
			input:    []int{1, 2, 3, 4, 5},
			left:     2,
			right:    4,
			expected: []int{1, 4, 3, 2, 5},
		},
		{
			name:     "LeetCode Example 2",
			input:    []int{5},
			left:     1,
			right:    1,
			expected: []int{5},
		},
		{
			name:     "Reverse entire list",
			input:    []int{1, 2, 3, 4, 5},
			left:     1,
			right:    5,
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Reverse first two elements",
			input:    []int{1, 2, 3, 4, 5},
			left:     1,
			right:    2,
			expected: []int{2, 1, 3, 4, 5},
		},
		{
			name:     "Reverse last two elements",
			input:    []int{1, 2, 3, 4, 5},
			left:     4,
			right:    5,
			expected: []int{1, 2, 3, 5, 4},
		},
		{
			name:     "Reverse single element in middle",
			input:    []int{1, 2, 3, 4, 5},
			left:     3,
			right:    3,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Empty list",
			input:    []int{},
			left:     1,
			right:    1,
			expected: []int{},
		},
		{
			name:     "Two elements reverse both",
			input:    []int{1, 2},
			left:     1,
			right:    2,
			expected: []int{2, 1},
		},
		{
			name:     "Three elements reverse middle",
			input:    []int{1, 2, 3},
			left:     2,
			right:    2,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Large list reverse middle section",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			left:     3,
			right:    7,
			expected: []int{1, 2, 7, 6, 5, 4, 3, 8, 9, 10},
		},
		{
			name:     "All same values",
			input:    []int{5, 5, 5, 5, 5},
			left:     2,
			right:    4,
			expected: []int{5, 5, 5, 5, 5},
		},
		{
			name:     "Negative numbers",
			input:    []int{-3, -2, -1, 0, 1},
			left:     2,
			right:    4,
			expected: []int{-3, 0, -1, -2, 1},
		},
		{
			name:     "Reverse from second to second last",
			input:    []int{1, 2, 3, 4, 5},
			left:     2,
			right:    4,
			expected: []int{1, 4, 3, 2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputList := createList(tt.input)
			result := reverseBetween(inputList, tt.left, tt.right)
			resultSlice := listToSlice(result)
			assert.Equal(t, tt.expected, resultSlice, 
				"reverseBetween(%v, %d, %d) should return %v, got %v", 
				tt.input, tt.left, tt.right, tt.expected, resultSlice)
		})
	}
}