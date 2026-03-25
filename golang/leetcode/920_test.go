package main

import (
	"testing"
)

func TestNumMusicPlaylists(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		goal     int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			n:        3,
			goal:     3,
			k:        1,
			expected: 6,
		},
		{
			name:     "Example 2",
			n:        2,
			goal:     3,
			k:        0,
			expected: 6,
		},
		{
			name:     "Example 3",
			n:        2,
			goal:     3,
			k:        1,
			expected: 2,
		},
		{
			name:     "Custom Test 1: n=1, goal=1, k=0",
			n:        1,
			goal:     1,
			k:        0,
			expected: 1,
		},
		{
			name:     "Custom Test 2: n=1, goal=2, k=0",
			n:        1,
			goal:     2,
			k:        0,
			expected: 1,
		},
		{
			name:     "Custom Test 3: n=1, goal=2, k=1",
			n:        1,
			goal:     2,
			k:        1,
			expected: 0, // Cannot play 2 distinct songs with k=1 if only 1 song available.
		},
		{
			name:     "Custom Test 4: n=3, goal=4, k=1",
			n:        3,
			goal:     4,
			k:        1,
			expected: 18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMusicPlaylists(tt.n, tt.goal, tt.k); got != tt.expected {
				t.Errorf("numMusicPlaylists(%d, %d, %d) = %d, want %d", tt.n, tt.goal, tt.k, got, tt.expected)
			}
		})
	}
}
