package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		// Basic cases
		{"empty list", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{1, 2}},
		{"three elements", []int{1, 2, 3}, []int{1, 3, 2}},
		
		// Classic cases
		{"four elements", []int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{"five elements", []int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
		{"six elements", []int{1, 2, 3, 4, 5, 6}, []int{1, 6, 2, 5, 3, 4}},
		
		// Edge cases
		{"all same values", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{"negative numbers", []int{-1, -2, -3, -4}, []int{-1, -4, -2, -3}},
		{"mixed positive negative", []int{-1, 2, -3, 4}, []int{-1, 4, 2, -3}},
		
		// Complex cases
		{"large list odd", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 7, 2, 6, 3, 5, 4}},
		{"large list even", []int{1, 2, 3, 4, 5, 6, 7, 8}, []int{1, 8, 2, 7, 3, 6, 4, 5}},
		{"duplicate values mixed", []int{1, 2, 1, 3, 2}, []int{1, 2, 2, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputList := createList(tt.input)
			reorderList(inputList)
			result := listToSlice(inputList)
			assert.Equal(t, tt.expected, result, "Reordered list should match expected pattern")
		})
	}
}