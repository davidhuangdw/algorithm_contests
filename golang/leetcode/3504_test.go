package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "abcde",
			t:        "ecdba",
			expected: 5,
		},
		{
			name:     "Full overlap",
			s:        "abc",
			t:        "cba",
			expected: 6,
		},
		{
			name:     "No overlap, use internal palindrome",
			s:        "abacaba",
			t:        "z",
			expected: 7,
		},
		{
			name:     "Overlap with internal palindrome",
			s:        "ab",
			t:        "ba",
			expected: 4,
		},
		{
			name:     "All same characters",
			s:        "aaaaa",
			t:        "aaaaa",
			expected: 10,
		},
		{
			name:     "No common characters",
			s:        "abc",
			t:        "def",
			expected: 1,
		},
		{
			name:     "One empty string",
			s:        "",
			t:        "a",
			expected: 1,
		},
		{
			name:     "Both empty strings",
			s:        "",
			t:        "",
			expected: 0,
		},
		{
			name:     "Single characters",
			s:        "a",
			t:        "b",
			expected: 1,
		},
		{
			name:     "Split and merge",
			s:        "ghabcdef",
			t:        "fedcbaiz",
			expected: 12, // "abcdef" from s, "fedcba" from t. Total 12.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.s, tt.t)
			if got != tt.expected {
				t.Errorf("longestPalindrome(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.expected)
			}
		})
	}
}
