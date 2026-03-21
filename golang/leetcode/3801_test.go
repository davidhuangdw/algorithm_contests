package main

import (
	"testing"
)

func TestMinMergeCost(t *testing.T) {
	tests := []struct {
		name     string
		lists    [][]int
		expected int64
	}{
		{
			name:     "single list",
			lists:    [][]int{{1, 3, 5}},
			expected: 0,
		},
		{
			name:     "two lists",
			lists:    [][]int{{1, 5}, {3, 4}},
			expected: 6,
		},
		{
			name:     "three simple lists",
			lists:    [][]int{{1}, {2}, {3}},
			expected: 7,
		},
		{
			name:     "three lists",
			lists:    [][]int{{1, 10}, {2, 9}, {3, 8}},
			expected: 12,
		},
		{
			name:     "empty input",
			lists:    [][]int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMergeCost(tt.lists); got != tt.expected {
				t.Errorf("minMergeCost() = %v, want %v", got, tt.expected)
			}
		})
	}
}
