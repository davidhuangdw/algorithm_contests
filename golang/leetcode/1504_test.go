package main

import "testing"

func TestNumSubmat(t *testing.T) {
	tests := []struct {
		mat      [][]int
		expected int
	}{
		{
			mat:      [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}},
			expected: 13,
		},
		{
			mat:      [][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}},
			expected: 24,
		},
		{
			mat:      [][]int{{1, 1, 1, 1, 1, 1}},
			expected: 21,
		},
		{
			mat:      [][]int{{1, 0}, {1, 0}},
			expected: 3,
		},
		{
			mat:      [][]int{{1}},
			expected: 1,
		},
		{
			mat:      [][]int{{0}},
			expected: 0,
		},
	}

	for _, tt := range tests {
		if got := numSubmat(tt.mat); got != tt.expected {
			t.Errorf("numSubmat(%v) = %v, want %v", tt.mat, got, tt.expected)
		}
	}
}
