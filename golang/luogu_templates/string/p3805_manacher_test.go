package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManacher(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []int
		description string
	}{
		{
			name:        "Empty string",
			input:       "",
			expected:    []int{},
			description: "Empty input should return empty slice",
		},
		{
			name:        "Single character",
			input:       "a",
			expected:    []int{0, 1},
			description: "Single character has even center (0) and odd center (1)",
		},
		{
			name:        "Two different characters",
			input:       "ab",
			expected:    []int{0, 1, 0, 1},
			description: "Two different characters, each has even center (0) and odd center (1)",
		},
		{
			name:        "Two same characters",
			input:       "aa",
			expected:    []int{0, 1, 2, 1},
			description: "Even center between two a's has length 2, odd centers have length 1",
		},
		{
			name:        "Odd length palindrome",
			input:       "aba",
			expected:    []int{0, 1, 0, 3, 0, 1},
			description: "Odd center at b has length 3, others have length 1 or 0",
		},
		{
			name:        "Even length palindrome",
			input:       "abba",
			expected:    []int{0, 1, 0, 1, 4, 1, 0, 1},
			description: "Even center between two b's has length 4, odd centers have length 1 or 2",
		},
		{
			name:        "Entire string is palindrome (odd)",
			input:       "abcba",
			expected:    []int{0, 1, 0, 1, 0, 5, 0, 1, 0, 1},
			description: "Odd center at c has length 5, others have length 1 or 0",
		},
		{
			name:        "Entire string is palindrome (even)",
			input:       "abccba",
			expected:    []int{0, 1, 0, 1, 0, 1, 6, 1, 0, 1, 0, 1},
			description: "Even center between two c's has length 6, others have length 1 or 0",
		},
		{
			name:        "Multiple palindromes",
			input:       "aabaa",
			expected:    []int{0, 1, 2, 1, 0, 5, 0, 1, 2, 1},
			description: "Multiple palindromes including full string (length 5) and 'aa' (length 2)",
		},
		{
			name:        "Palindrome at beginning",
			input:       "aaaaab",
			expected:    []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0, 1},
			description: "Long palindrome at beginning followed by a different character",
		},
		{
			name:        "Palindrome at end",
			input:       "abaaaaa",
			expected:    []int{0, 1, 0, 3, 0, 1, 2, 3, 4, 5, 4, 3, 2, 1},
			description: "Different character followed by long palindrome at end",
		},
		{
			name:        "No long palindromes",
			input:       "abcdef",
			expected:    []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			description: "All characters unique, only single-character palindromes",
		},
		{
			name:        "Complex palindromes",
			input:       "abacabacabb",
			expected:    []int{0, 1, 0, 3, 0, 1, 0, 7, 0, 1, 0, 9, 0, 1, 0, 5, 0, 1, 0, 1, 2, 1},
			description: "Complex string with multiple overlapping palindromes",
		},
		{
			name:        "Repeating characters",
			input:       "aaaaa",
			expected:    []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1},
			description: "All characters same, forming symmetric palindrome lengths",
		},
		{
			name:        "Mixed case (non-palindrome due to case)",
			input:       "Aa",
			expected:    []int{0, 1, 0, 1},
			description: "Case-sensitive, so 'Aa' has no long palindromes",
		},
		{
			name:        "Long palindrome with unique center",
			input:       "abcddcba",
			expected:    []int{0, 1, 0, 1, 0, 1, 0, 1, 8, 1, 0, 1, 0, 1, 0, 1},
			description: "Even center between two d's has length 8, others have length 1 or 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := manacher(tt.input)
			assert.Equal(t, tt.expected, result, "%s: Expected %v, got %v", tt.description, tt.expected, result)
		})
	}
}

// Test the main functionality of finding the longest palindrome
func TestLongestPalindromeLength(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedLength int
		description    string
	}{
		{
			name:           "Empty string",
			input:          "",
			expectedLength: 0,
			description:    "Empty string has no palindromes",
		},
		{
			name:           "Single character",
			input:          "a",
			expectedLength: 1,
			description:    "Single character is a palindrome of length 1",
		},
		{
			name:           "Even palindrome",
			input:          "abba",
			expectedLength: 4,
			description:    "Even length palindrome",
		},
		{
			name:           "Odd palindrome",
			input:          "aba",
			expectedLength: 3,
			description:    "Odd length palindrome",
		},
		{
			name:           "Longest at beginning",
			input:          "aaabbb",
			expectedLength: 3,
			description:    "Longest palindrome at the beginning",
		},
		{
			name:           "Longest in middle",
			input:          "abacaba",
			expectedLength: 7,
			description:    "Longest palindrome in the middle",
		},
		{
			name:           "Longest at end",
			input:          "abbbbb",
			expectedLength: 5,
			description:    "Longest palindrome at the end",
		},
		{
			name:           "All unique characters",
			input:          "abcdefg",
			expectedLength: 1,
			description:    "All unique characters have palindrome length 1",
		},
		{
			name:           "Repeating characters",
			input:          "aaaaa",
			expectedLength: 5,
			description:    "All repeating characters form a single palindrome",
		},
		{
			name:           "Complex string",
			input:          "babadada",
			expectedLength: 5,
			description:    "Complex string with multiple palindromes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replicate the logic from P3805_manacher function
			dis := manacher(tt.input)
			maxLen := 0
			for _, d := range dis {
				if d > maxLen {
					maxLen = d
				}
			}
			assert.Equal(t, tt.expectedLength, maxLen, "%s: Expected longest palindrome length %d, got %d", tt.description, tt.expectedLength, maxLen)
		})
	}
}
