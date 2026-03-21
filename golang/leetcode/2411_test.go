package main

import (
	"reflect"
	"testing"
)

func TestSmallestSubarrays(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 0, 2, 1, 3},
			expected: []int{3, 3, 2, 2, 1},
		},
		{
			nums:     []int{1, 2},
			expected: []int{2, 1},
		},
		{
			nums:     []int{0},
			expected: []int{1},
		},
		{
			nums:     []int{1, 1, 1, 1, 1},
			expected: []int{1, 1, 1, 1, 1},
		},
		{
			nums:     []int{8, 4, 2, 1},
			expected: []int{4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		if got := smallestSubarrays(tt.nums); !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("smallestSubarrays(%v) = %v, want %v", tt.nums, got, tt.expected)
		}
	}
}
