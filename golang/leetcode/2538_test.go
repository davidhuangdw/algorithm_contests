package main

import (
	"testing"
)

func TestMaxOutput(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		price    []int
		expected int64
	}{
		{
			name:     "Example 1",
			n:        6,
			edges:    [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}},
			price:    []int{9, 8, 7, 6, 10, 5},
			expected: 24,
		},
		{
			name:     "Example 2",
			n:        3,
			edges:    [][]int{{0, 1}, {1, 2}},
			price:    []int{1, 1, 1},
			expected: 2,
		},
		{
			name:     "Single node",
			n:        1,
			edges:    [][]int{},
			price:    []int{5},
			expected: 0,
		},
		{
			name:     "Star graph",
			n:        4,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}},
			price:    []int{1, 2, 3, 4},
			expected: 5, // Root at 2, path [2,0,3]=2+1+4=7, min=price[2]=2, 7-2=5. Or root at 3, path [3,0,2]=4+1+3=8, min=price[3]=4, 8-4=4. Wait.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxOutput(tt.n, tt.edges, tt.price)
			if got != tt.expected {
				t.Errorf("maxOutput() = %v, want %v", got, tt.expected)
			}
			got1 := maxOutput1(tt.n, tt.edges, tt.price)
			if got1 != tt.expected {
				t.Errorf("maxOutput1() = %v, want %v", got1, tt.expected)
			}
		})
	}
}
