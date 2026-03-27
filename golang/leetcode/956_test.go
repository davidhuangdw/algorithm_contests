package main

import (
	"testing"
)

func TestTallestBillboard(t *testing.T) {
	tests := []struct {
		name     string
		rods     []int
		expected int
	}{
		{
			name:     "Example 1",
			rods:     []int{1, 2, 3, 6},
			expected: 6,
		},
		{
			name:     "Example 2",
			rods:     []int{1, 2, 3, 4, 5, 6},
			expected: 10,
		},
		{
			name:     "Example 3",
			rods:     []int{1, 2},
			expected: 0,
		},
		{
			name:     "Single rod",
			rods:     []int{5},
			expected: 0,
		},
		{
			name:     "Empty rods",
			rods:     []int{},
			expected: 0,
		},
		{
			name:     "Two equal rods",
			rods:     []int{3, 3},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tallestBillboard(tt.rods); got != tt.expected {
				t.Errorf("tallestBillboard(%v) = %v, want %v", tt.rods, got, tt.expected)
			}
		})
	}
}
