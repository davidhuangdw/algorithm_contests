package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		// Basic cases
		{"single duplicate in small array", []int{1, 3, 4, 2, 2}, 2},
		{"duplicate at beginning", []int{3, 1, 3, 4, 2}, 3},
		{"duplicate at end", []int{1, 4, 4, 2, 3}, 4},
		
		// Classic cases
		{"LeetCode example 1", []int{1, 3, 4, 2, 2}, 2},
		{"LeetCode example 2", []int{3, 1, 3, 4, 2}, 3},
		{"LeetCode example 3", []int{2, 2, 2, 2, 2}, 2},
		
		// Edge cases
		{"minimum size array", []int{1, 1}, 1},
		{"three elements with duplicate", []int{2, 1, 2}, 2},
		{"all same numbers", []int{5, 5, 5, 5, 5}, 5},
		
		// Complex cases
		{"large array with duplicate in middle", []int{8, 7, 1, 10, 17, 15, 18, 11, 16, 9, 19, 12, 5, 14, 3, 4, 2, 13, 6, 4, 20}, 4},
		{"duplicate appears multiple times", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 5, 5, 5}, 5},
		{"numbers not in order", []int{10, 5, 3, 7, 1, 9, 8, 2, 6, 4, 7}, 7},
		{"consecutive duplicates", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 42, 42}, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the input array since the function modifies it
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			
			result := findDuplicate(numsCopy)
			assert.Equal(t, tt.expected, result, "Duplicate should be found correctly")
		})
	}
}