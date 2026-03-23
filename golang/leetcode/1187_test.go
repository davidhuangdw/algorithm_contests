package main

import "testing"

func TestMakeArrayIncreasing(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected int
	}{
		{
			name:     "Example 1",
			a:        []int{1, 5, 3, 6, 7},
			b:        []int{1, 3, 2, 4},
			expected: 1,
		},
		{
			name:     "Example 2",
			a:        []int{1, 5, 3, 6, 7},
			b:        []int{4, 3, 1},
			expected: 2,
		},
		{
			name:     "Example 3: Impossible",
			a:        []int{1, 5, 3, 6, 7},
			b:        []int{1, 6, 3, 3},
			expected: -1,
		},
		{
			name:     "Already increasing",
			a:        []int{1, 2, 3, 4, 5},
			b:        []int{6, 7, 8},
			expected: 0,
		},
		{
			name:     "Empty b, impossible",
			a:        []int{5, 4, 3, 2, 1},
			b:        []int{},
			expected: -1,
		},
		{
			name:     "Needs one change at the end",
			a:        []int{1, 2, 5, 4},
			b:        []int{3, 6},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeArrayIncreasing(tt.a, tt.b); got != tt.expected {
				t.Errorf("makeArrayIncreasing() = %v, want %v", got, tt.expected)
			}
		})
	}
}
