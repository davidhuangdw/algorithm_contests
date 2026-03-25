package main

import (
	"reflect"
	"testing"
)

func TestMinDeletions(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		queries  [][]int
		expected []int
	}{
		{
			name: "Simple case with updates",
			s:    "01001",
			queries: [][]int{
				{2, 1, 3}, // Substring "100". One non-alternating pair ("00"). Expected: 1.
				{1, 2},    // Flip s[2]. s becomes "01101".
				{2, 1, 3}, // Substring "110". One non-alternating pair ("11"). Expected: 1.
			},
			expected: []int{1, 1},
		},
		{
			name: "All same characters",
			s:    "00000",
			queries: [][]int{
				{2, 0, 4}, // Substring "00000". Four non-alternating pairs. Expected: 4.
			},
			expected: []int{4},
		},
		{
			name: "Perfectly alternating",
			s:    "010101",
			queries: [][]int{
				{2, 0, 5}, // Substring "010101". Zero non-alternating pairs. Expected: 0.
			},
			expected: []int{0},
		},
		{
			name: "Update creates a non-alternating pair",
			s:    "01010",
			queries: [][]int{
				{2, 0, 4}, // Substring "01010". 0 non-alternating pairs.
				{1, 1},    // Flip s[1]. s becomes "00010".
				{2, 0, 4}, // Substring "00010". 2 non-alternating pairs ("00" at index 0-1 and 1-2).
			},
			expected: []int{0, 2},
		},
		{
			name: "Update breaks a non-alternating pair",
			s:    "00010",
			queries: [][]int{
				{2, 0, 4}, // Substring "00010". 2 non-alternating pairs.
				{1, 1},    // Flip s[1]. s becomes "01010".
				{2, 0, 4}, // Substring "01010". 0 non-alternating pairs.
			},
			expected: []int{2, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletions(tt.s, tt.queries); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("minDeletions() = %v, want %v", got, tt.expected)
			}
		})
	}
}
