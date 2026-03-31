package main

import (
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int64
	}{
		{
			name:     "Example 1",
			s:        "12936",
			expected: 11,
		},
		{
			name:     "Example 2",
			s:        "5701283",
			expected: 18,
		},
		{
			name:     "Example 3",
			s:        "1010101010",
			expected: 25,
		},
		{
			name:     "Single non-zero digit",
			s:        "5",
			expected: 1,
		},
		{
			name:     "Single zero digit",
			s:        "0",
			expected: 0,
		},
		{
			name:     "All zeros",
			s:        "000",
			expected: 0,
		},
		{
			name:     "Zeros and non-zeros",
			s:        "0102",
			expected: 6, // Index 1 ('1'): "01", "1" (2); Index 3 ('2'): "0102" (102), "102", "02", "2" (4). Total 6.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings(tt.s); got != tt.expected {
				t.Errorf("countSubstrings(%q) = %v, want %v", tt.s, got, tt.expected)
			}
		})
	}
}
