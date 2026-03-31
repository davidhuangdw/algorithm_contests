package main

import (
	"testing"
)

func TestMaxStudents(t *testing.T) {
	tests := []struct {
		name     string
		seats    [][]byte
		expected int
	}{
		{
			name: "Example 1",
			seats: [][]byte{
				{'#', '.', '#', '#', '.', '#'},
				{'.', '#', '#', '#', '#', '.'},
				{'#', '.', '#', '#', '.', '#'},
			},
			expected: 4,
		},
		{
			name: "Example 2",
			seats: [][]byte{
				{'.', '#'},
				{'#', '#'},
				{'#', '.'},
				{'#', '#'},
				{'.', '#'},
			},
			expected: 3,
		},
		{
			name: "Example 3",
			seats: [][]byte{
				{'#', '.', '.', '.', '#'},
				{'.', '#', '.', '#', '.'},
				{'.', '.', '#', '.', '.'},
				{'.', '#', '.', '#', '.'},
				{'#', '.', '.', '.', '#'},
			},
			expected: 10,
		},
		{
			name: "Small grid all broken",
			seats: [][]byte{
				{'#', '#'},
				{'#', '#'},
			},
			expected: 0,
		},
		{
			name: "Small grid all good",
			seats: [][]byte{
				{'.', '.'},
				{'.', '.'},
			},
			expected: 2, // (0,0) and (1,1) or (0,1) and (1,0)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxStudents(tt.seats); got != tt.expected {
				t.Errorf("maxStudents() = %v, want %v", got, tt.expected)
			}
		})
	}
}
