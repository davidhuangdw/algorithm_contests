package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		// Basic cases
		{"empty list", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"three elements", []int{1, 2, 3}, []int{3, 2, 1}},

		// Classic cases
		{"five elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"all same values", []int{1, 1, 1}, []int{1, 1, 1}},

		// Edge cases
		{"negative numbers", []int{-1, -2, -3}, []int{-3, -2, -1}},
		{"mixed numbers", []int{-1, 0, 1}, []int{1, 0, -1}},

		// Complex cases
		{"large list", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{"single element repeated", []int{5, 5, 5, 5}, []int{5, 5, 5, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputList := createList(tt.input)
			reversed := reverseList(inputList)
			result := listToSlice(reversed)
			assert.Equal(t, tt.expected, result, "Reversed list should match expected")
		})
	}
}
