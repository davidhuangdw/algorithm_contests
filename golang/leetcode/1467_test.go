package main

import (
	"math"
	"testing"
)

func TestGetProbability(t *testing.T) {
	tests := []struct {
		name     string
		balls    []int
		expected float64
	}{
		{
			name:     "Example 1",
			balls:    []int{1, 1},
			expected: 1.00000,
		},
		{
			name:     "Example 2",
			balls:    []int{2, 1, 1},
			expected: 0.66667,
		},
		{
			name:     "Example 3",
			balls:    []int{1, 2, 1, 2},
			expected: 0.60000,
		},
		{
			name:     "Uniform colors",
			balls:    []int{2, 2},
			expected: 1.00000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getProbability(tt.balls)
			if math.Abs(got-tt.expected) > 1e-5 {
				t.Errorf("getProbability(%v) = %v, want %v", tt.balls, got, tt.expected)
			}
		})
	}
}
