package main

import "testing"

func TestMinimumTimeRequired(t *testing.T) {
	tests := []struct {
		jobs     []int
		k        int
		expected int
	}{
		{
			jobs:     []int{3, 2, 3},
			k:        3,
			expected: 3,
		},
		{
			jobs:     []int{1, 2, 4, 7, 8},
			k:        2,
			expected: 11,
		},
		{
			jobs:     []int{5, 1, 6, 2, 8},
			k:        3,
			expected: 8,
		},
		{
			jobs:     []int{1, 2, 3, 4, 5},
			k:        1,
			expected: 15,
		},
		{
			jobs:     []int{5, 4, 3, 2, 1},
			k:        5,
			expected: 5,
		},
	}

	for _, tt := range tests {
		if got := minimumTimeRequired(tt.jobs, tt.k); got != tt.expected {
			t.Errorf("minimumTimeRequired(%v, %d) = %v, want %v", tt.jobs, tt.k, got, tt.expected)
		}
	}
}
