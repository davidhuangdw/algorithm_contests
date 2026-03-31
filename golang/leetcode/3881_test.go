package main

import (
	"testing"
)

func Test_countVisiblePeople(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		pos      int
		k        int
		expected int
	}{
		{"Example 1", 3, 2, 1, 4},
		{"Example 2", 1, 0, 0, 2},
		{"Example 3", 3, 1, 0, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVisiblePeople(tt.n, tt.pos, tt.k); got != tt.expected {
				t.Errorf("countVisiblePeople(%v, %v, %v) = %v, want %v", tt.n, tt.pos, tt.k, got, tt.expected)
			}
		})
	}
}
