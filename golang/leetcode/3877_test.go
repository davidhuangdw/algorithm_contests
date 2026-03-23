package main

import "testing"

func TestMinRemovals(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3},
			target:   0,
			expected: 0,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3, 4, 5},
			target:   7,
			expected: 2, // Keep {3, 4} -> XOR sum is 7. Remove 3 elements.
		},
		{
			name:     "No solution",
			nums:     []int{1, 2, 4, 8},
			target:   100,
			expected: -1,
		},
		{
			name:     "Keep all elements",
			nums:     []int{2, 4, 6},
			target:   0,
			expected: 0,
		},
		{
			name:     "Empty nums",
			nums:     []int{},
			target:   0,
			expected: 0,
		},
		{
			name:     "Empty nums, non-zero target",
			nums:     []int{},
			target:   5,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemovals(tt.nums, tt.target); got != tt.expected {
				t.Errorf("minRemovals() = %v, want %v", got, tt.expected)
			}
		})
	}
}
