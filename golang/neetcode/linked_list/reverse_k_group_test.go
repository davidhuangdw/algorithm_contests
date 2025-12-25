package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		// Basic cases
		{"empty list", []int{}, 2, []int{}},
		{"single element k=1", []int{1}, 1, []int{1}},
		{"single element k=2", []int{1}, 2, []int{1}},
		
		// Classic cases - k=2
		{"two elements k=2", []int{1, 2}, 2, []int{2, 1}},
		{"three elements k=2", []int{1, 2, 3}, 2, []int{2, 1, 3}},
		{"four elements k=2", []int{1, 2, 3, 4}, 2, []int{2, 1, 4, 3}},
		{"five elements k=2", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		
		// Classic cases - k=3
		{"three elements k=3", []int{1, 2, 3}, 3, []int{3, 2, 1}},
		{"four elements k=3", []int{1, 2, 3, 4}, 3, []int{3, 2, 1, 4}},
		{"five elements k=3", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{"six elements k=3", []int{1, 2, 3, 4, 5, 6}, 3, []int{3, 2, 1, 6, 5, 4}},
		
		// Edge cases
		{"k=1 (no reversal)", []int{1, 2, 3, 4}, 1, []int{1, 2, 3, 4}},
		{"k equals list length", []int{1, 2, 3, 4}, 4, []int{4, 3, 2, 1}},
		{"k larger than list length", []int{1, 2, 3}, 5, []int{1, 2, 3}},
		{"all same values k=2", []int{1, 1, 1, 1}, 2, []int{1, 1, 1, 1}},
		
		// Complex cases
		{"large list k=4", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, []int{4, 3, 2, 1, 8, 7, 6, 5, 9, 10}},
		{"negative numbers k=3", []int{-3, -2, -1, 0, 1, 2}, 3, []int{-1, -2, -3, 2, 1, 0}},
		{"mixed values k=2", []int{5, -3, 8, 0, -1, 4}, 2, []int{-3, 5, 0, 8, 4, -1}},
		{"LeetCode example 1", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{"LeetCode example 2", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create input linked list
			inputList := createList(tt.input)
			
			// Reverse in k-groups
			result := reverseKGroup(inputList, tt.k)
			
			// Convert result back to slice for comparison
			resultSlice := listToSlice(result)
			
			assert.Equal(t, tt.expected, resultSlice, 
				"reverseKGroup(%v, %d) should return %v, got %v", 
				tt.input, tt.k, tt.expected, resultSlice)
		})
	}
}