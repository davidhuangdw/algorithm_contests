package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected float64
	}{
		{"test", []int{1, 3}, []int{2, 4}, 2.5},
		// Basic cases
		{"single element arrays", []int{1}, []int{2}, 1.5},
		{"empty first array", []int{}, []int{1, 2, 3}, 2.0},
		{"empty second array", []int{1, 2, 3}, []int{}, 2.0},
		{"both empty arrays", []int{}, []int{}, 0.0},

		// Classic cases
		{"odd total length", []int{1, 3}, []int{2}, 2.0},
		{"even total length", []int{1, 2}, []int{3, 4}, 2.5},
		{"interleaved arrays", []int{1, 3, 5}, []int{2, 4, 6}, 3.5},

		// Edge cases
		{"all same elements", []int{1, 1, 1}, []int{1, 1}, 1.0},
		{"negative numbers", []int{-5, -3, -1}, []int{-4, -2}, -3.0},
		{"mixed positive negative", []int{-1, 3}, []int{1, 2}, 1.5},

		// Complex cases
		{"different sizes", []int{1, 2, 3, 4, 5}, []int{6, 7, 8}, 4.5},
		{"large arrays", []int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}, 5.5},
		{"one array much larger", []int{1}, []int{2, 3, 4, 5, 6, 7, 8}, 4.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMedianSortedArrays(tt.nums1, tt.nums2)
			assert.Equal(t, tt.expected, result, "Median of %v and %v should be %v", tt.nums1, tt.nums2, tt.expected)
		})
	}
}
