package main

import "testing"

func TestNumberOfRoutes(t *testing.T) {
	tests := []struct {
		name     string
		grid     []string
		d        int
		expected int
	}{
		{
			name:     "Example from main",
			grid:     []string{"..", "#."},
			d:        1,
			expected: 2,
		},
		{
			name:     "Simple grid, no obstacles",
			grid:     []string{"..", ".."},
			d:        1,
			expected: 8,
		},
		{
			name:     "Start blocked",
			grid:     []string{"#.", ".."},
			d:        1,
			expected: 2,
		},
		{
			name:     "End blocked",
			grid:     []string{"..", ".#"},
			d:        1,
			expected: 2,
		},
		{
			name:     "Larger d",
			grid:     []string{"...", "...", "..."},
			d:        2,
			expected: 441,
		},
		{
			name:     "Single cell, open",
			grid:     []string{"."},
			d:        1,
			expected: 1,
		},
		{
			name:     "Single cell, blocked",
			grid:     []string{"#"},
			d:        1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfRoutes(tt.grid, tt.d); got != tt.expected {
				t.Errorf("numberOfRoutes() = %v, want %v", got, tt.expected)
			}
		})
	}
}
