package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextLarge(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Simple sequence",
			input:    []int{1, 4, 2, 3, 5},
			expected: []int{2, 5, 4, 5, 0},
		},
		{
			name:     "Decreasing sequence",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			name:     "Increasing sequence",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{2, 3, 4, 5, 0},
		},
		{
			name:     "All same values",
			input:    []int{3, 3, 3, 3},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{0},
		},
		{
			name:     "Two elements ascending",
			input:    []int{1, 2},
			expected: []int{2, 0},
		},
		{
			name:     "Two elements descending",
			input:    []int{2, 1},
			expected: []int{0, 0},
		},
		{
			name:     "Mountain pattern",
			input:    []int{1, 5, 3, 4, 2},
			expected: []int{2, 0, 4, 0, 0},
		},
		{
			name:     "Valley pattern",
			input:    []int{5, 1, 4, 2, 6},
			expected: []int{5, 3, 5, 5, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := nextLarge(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
