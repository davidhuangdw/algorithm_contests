package main

import (
	"testing"
)

func TestMinimumScore(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		edges    [][]int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 5, 5, 4, 11},
			edges:    [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}},
			expected: 9,
		},
		{
			name:     "Example 2",
			nums:     []int{5, 5, 2, 4, 4, 2},
			edges:    [][]int{{0, 1}, {1, 2}, {5, 2}, {4, 3}, {1, 3}},
			expected: 0,
		},
		{
			name:     "Small tree",
			nums:     []int{1, 2, 3},
			edges:    [][]int{{0, 1}, {1, 2}},
			expected: 2, // components: {1}, {2}, {3}. max-min = 3-1 = 2.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumScore(tt.nums, tt.edges); got != tt.expected {
				t.Errorf("minimumScore() = %v, want %v", got, tt.expected)
			}
			if got1 := minimumScore1(tt.nums, tt.edges); got1 != tt.expected {
				t.Errorf("minimumScore1() = %v, want %v", got1, tt.expected)
			}
		})
	}
}
