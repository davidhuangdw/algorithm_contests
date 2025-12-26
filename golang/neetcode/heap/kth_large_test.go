package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		// Basic cases
		{"single element", []int{5}, 1, 5},
		{"first largest", []int{3, 2, 1}, 1, 3},
		{"second largest", []int{3, 2, 1}, 2, 2},
		{"third largest", []int{3, 2, 1}, 3, 1},

		// Classic cases - LeetCode examples
		{"LeetCode example 1", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"LeetCode example 2", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},

		// Edge cases
		{"all same numbers", []int{7, 7, 7, 7}, 2, 7},
		{"negative numbers", []int{-1, -2, -3, -4}, 1, -1},
		{"mixed positive negative", []int{-1, 2, -3, 4}, 3, -1},
		{"k equals length", []int{1, 2, 3}, 3, 1},

		// Complex cases
		{"large numbers", []int{1000, 2000, 3000, 4000}, 2, 3000},
		{"unsorted array", []int{5, 3, 8, 1, 9, 2}, 3, 5},
		{"duplicates with k", []int{5, 5, 5, 2, 2, 3}, 3, 5},
		{"descending order", []int{9, 8, 7, 6, 5, 4}, 4, 6},
		
		// Additional edge cases from separate function
		{"k=1 largest element", []int{1, 9, 3, 7, 5}, 1, 9},
		{"k=len(nums) smallest element", []int{1, 9, 3, 7, 5}, 5, 1},
		{"single element test", []int{42}, 1, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findKthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result,
				"findKthLargest(%v, %d) = %d, expected %d",
				tt.nums, tt.k, result, tt.expected)
		})
	}
}