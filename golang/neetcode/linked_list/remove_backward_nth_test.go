package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		// Basic cases
		{"remove first from end (single element)", []int{1}, 1, []int{}},
		{"remove first from end (two elements)", []int{1, 2}, 1, []int{1}},
		{"remove second from end (two elements)", []int{1, 2}, 2, []int{2}},
		{"remove middle from three elements", []int{1, 2, 3}, 2, []int{1, 3}},

		// Classic cases
		{"remove first from end (classic)", []int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4}},
		{"remove second from end (classic)", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{"remove third from end (classic)", []int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}},
		{"remove fourth from end (classic)", []int{1, 2, 3, 4, 5}, 4, []int{1, 3, 4, 5}},
		{"remove fifth from end (classic)", []int{1, 2, 3, 4, 5}, 5, []int{2, 3, 4, 5}},

		// Edge cases
		{"remove from empty list", []int{}, 1, []int{}},
		{"remove beyond list length", []int{1, 2}, 3, []int{1, 2}},
		{"all same values", []int{1, 1, 1, 1}, 2, []int{1, 1, 1}},
		{"negative numbers", []int{-1, -2, -3}, 1, []int{-1, -2}},

		// Complex cases
		{"large list remove first from end", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"large list remove middle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, []int{1, 2, 3, 4, 5, 7, 8, 9, 10}},
		{"large list remove last from end", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, []int{2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputList := createList(tt.input)
			resultList := removeNthFromEnd(inputList, tt.n)
			result := listToSlice(resultList)
			assert.Equal(t, tt.expected, result, "List after removing %dth from end should match expected", tt.n)
		})
	}
}
