package main

import "testing"

func TestIsEscapePossible(t *testing.T) {
	tests := []struct {
		name     string
		blocked  [][]int
		source   []int
		target   []int
		expected bool
	}{
		{
			name:     "Example 1: Trapped by blocks",
			blocked:  [][]int{{0, 1}, {1, 0}},
			source:   []int{0, 0},
			target:   []int{0, 2},
			expected: false,
		},
		{
			name:     "Example 2: No blocks",
			blocked:  [][]int{},
			source:   []int{0, 0},
			target:   []int{999999, 999999},
			expected: true,
		},
		{
			name:     "Source and target are far apart, no blocks",
			blocked:  [][]int{},
			source:   []int{0, 0},
			target:   []int{100, 100},
			expected: true,
		},
		{
			name:     "Source is completely enclosed",
			blocked:  [][]int{{0, 1}, {1, 0}, {1, 2}, {2, 1}},
			source:   []int{1, 1},
			target:   []int{3, 3},
			expected: false,
		},
		{
			name:     "Target is completely enclosed",
			blocked:  [][]int{{3, 4}, {4, 3}, {4, 5}, {5, 4}},
			source:   []int{0, 0},
			target:   []int{4, 4},
			expected: false,
		},
		{
			name:     "Source and target in same small enclosed area",
			blocked:  [][]int{{0, 1}, {1, 2}, {2, 1}, {1, 0}},
			source:   []int{1, 1},
			target:   []int{0, 0},
			expected: false,
		},
		{
			name:     "Large number of blocks not enclosing source or target",
			blocked:  [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}},
			source:   []int{0, 0},
			target:   []int{10, 10},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEscapePossible(tt.blocked, tt.source, tt.target); got != tt.expected {
				t.Errorf("isEscapePossible() = %v, want %v", got, tt.expected)
			}
		})
	}
}
