package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		// Basic cases
		{"both empty", []int{}, []int{}, []int{}},
		{"first empty", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"second empty", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"single element each", []int{1}, []int{2}, []int{1, 2}},

		// Classic cases
		{"interleaved lists", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"all from first", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"all from second", []int{4, 5, 6}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}},

		// Edge cases
		{"duplicate values", []int{1, 3, 3}, []int{2, 3, 4}, []int{1, 2, 3, 3, 3, 4}},
		{"negative numbers", []int{-3, -1}, []int{-4, -2}, []int{-4, -3, -2, -1}},
		{"mixed positive negative", []int{-1, 0}, []int{-2, 1}, []int{-2, -1, 0, 1}},

		// Complex cases
		{"different sizes", []int{1, 2}, []int{3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"large lists", []int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"one much larger", []int{1}, []int{2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := createList(tt.list1)
			l2 := createList(tt.list2)
			merged := mergeTwoLists(l1, l2)
			result := listToSlice(merged)
			assert.Equal(t, tt.expected, result, "Merged list should be sorted")
		})
	}
}
