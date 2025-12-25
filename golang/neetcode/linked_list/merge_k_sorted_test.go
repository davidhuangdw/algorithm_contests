package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name     string
		lists    [][]int
		expected []int
	}{
		// Basic cases
		{"empty list of lists", [][]int{}, []int{}},
		{"single empty list", [][]int{{}}, []int{}},
		{"multiple empty lists", [][]int{{}, {}, {}}, []int{}},
		
		// Classic cases
		{"single list", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"two lists same size", [][]int{{1, 3, 5}, {2, 4, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{"three lists", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{"LeetCode example", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		
		// Edge cases
		{"lists with one element each", [][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
		{"one list much longer", [][]int{{1, 2, 3, 4, 5}, {6}}, []int{1, 2, 3, 4, 5, 6}},
		{"duplicate values across lists", [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, []int{1, 1, 1, 1, 1, 1, 1, 1, 1}},
		
		// Complex cases
		{"four lists different sizes", [][]int{{1, 5, 9}, {2, 6, 10}, {3, 7}, {4, 8}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"negative numbers", [][]int{{-5, -3, -1}, {-4, -2, 0}}, []int{-5, -4, -3, -2, -1, 0}},
		{"mixed positive and negative", [][]int{{-1, 3, 5}, {-2, 0, 4}, {-3, 1, 6}}, []int{-3, -2, -1, 0, 1, 3, 4, 5, 6}},
		{"large lists", [][]int{{1, 3, 5, 7, 9, 11, 13, 15}, {2, 4, 6, 8, 10, 12, 14, 16}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert slice inputs to linked lists
			lists := make([]*ListNode, len(tt.lists))
			for i, values := range tt.lists {
				lists[i] = createList(values)
			}
			
			// Merge the lists
			result := mergeKLists(lists)
			
			// Convert result back to slice for comparison
			resultSlice := listToSlice(result)
			
			assert.Equal(t, tt.expected, resultSlice, "Merged list should match expected order")
		})
	}
}