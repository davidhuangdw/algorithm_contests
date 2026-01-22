package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPreLen(t *testing.T) {
	tests := []struct {
		name        string
		s           string
		expected    []int
		description string
	}{
		{
			name:        "Empty string",
			s:           "",
			expected:    []int{},
			description: "Empty string should return empty slice",
		},
		{
			name:        "Single character",
			s:           "a",
			expected:    []int{0},
			description: "Single character has no proper prefix/suffix",
		},
		{
			name:        "All same characters",
			s:           "aaaaa",
			expected:    []int{0, 1, 2, 3, 4},
			description: "Each position increases prefix length by 1",
		},
		{
			name:        "No repeated patterns",
			s:           "abcd",
			expected:    []int{0, 0, 0, 0},
			description: "No overlapping prefixes and suffixes",
		},
		{
			name:        "Standard KMP example",
			s:           "abababc",
			expected:    []int{0, 0, 1, 2, 3, 4, 0},
			description: "Classic KMP prefix function example",
		},
		{
			name:        "Prefix equals suffix",
			s:           "abcabc",
			expected:    []int{0, 0, 0, 1, 2, 3},
			description: "Full string prefix equals suffix",
		},
		{
			name:        "Complex pattern",
			s:           "aabaaab",
			expected:    []int{0, 1, 0, 1, 2, 2, 3},
			description: "Complex pattern with multiple overlapping matches",
		},
		{
			name:        "Pattern with repeating pairs",
			s:           "ababab",
			expected:    []int{0, 0, 1, 2, 3, 4},
			description: "Repeating pairs create increasing prefix lengths",
		},
		{
			name:        "One mismatch breaks pattern",
			s:           "aaaaab",
			expected:    []int{0, 1, 2, 3, 4, 0},
			description: "Final character breaks the pattern",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getPreLen(tt.s)
			assert.Equal(t, tt.expected, result, "%s: Expected %v, got %v", tt.description, tt.expected, result)
		})
	}
}

func TestMatchLen(t *testing.T) {
	tests := []struct {
		name        string
		tar         string
		pat         string
		expected    []int
		description string
	}{
		{
			name:        "Empty target",
			tar:         "",
			pat:         "abc",
			expected:    []int{},
			description: "Empty target should return empty slice",
		},
		{
			name:        "Empty pattern",
			tar:         "abc",
			pat:         "",
			expected:    []int{0, 0, 0},
			description: "Empty pattern returns all zeros",
		},
		{
			name:        "Single character exact match",
			tar:         "a",
			pat:         "a",
			expected:    []int{1},
			description: "Single character pattern matches single character target",
		},
		{
			name:        "No match",
			tar:         "abcde",
			pat:         "xyz",
			expected:    []int{0, 0, 0, 0, 0},
			description: "No characters match",
		},
		{
			name:        "Exact match at end",
			tar:         "abcdabc",
			pat:         "abc",
			expected:    []int{1, 2, 3, 0, 1, 2, 3},
			description: "Pattern matches exactly at the end",
		},
		{
			name:        "Multiple matches",
			tar:         "abababa",
			pat:         "aba",
			expected:    []int{1, 2, 3, 2, 3, 2, 3},
			description: "Multiple non-overlapping matches",
		},
		{
			name:        "Overlapping matches",
			tar:         "aaaaa",
			pat:         "aaa",
			expected:    []int{1, 2, 3, 3, 3},
			description: "Overlapping matches of same characters",
		},
		{
			name:        "Match at beginning",
			tar:         "abcdef",
			pat:         "abc",
			expected:    []int{1, 2, 3, 0, 0, 0},
			description: "Pattern matches exactly at beginning",
		},
		{
			name:        "Partial matches",
			tar:         "abcabxabc",
			pat:         "abcabc",
			expected:    []int{1, 2, 3, 4, 5, 0, 1, 2, 3},
			description: "Partial matches followed by a mismatch and restart",
		},
		{
			name:        "Pattern longer than target",
			tar:         "ab",
			pat:         "abc",
			expected:    []int{1, 2},
			description: "Pattern longer than target can only have partial matches",
		},
		{
			name:        "Complex KMP example",
			tar:         "ABABDABACDABABCABAB",
			pat:         "ABABCABAB",
			expected:    []int{1, 2, 3, 4, 0, 1, 2, 3, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			description: "Classic KMP matching example",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preLen := getPreLen(tt.pat)
			result := matchLen(tt.tar, tt.pat, preLen)
			assert.Equal(t, tt.expected, result, "%s: Expected %v, got %v", tt.description, tt.expected, result)
		})
	}
}

// Test the complete KMP functionality (integration test)
func TestKMPComplete(t *testing.T) {
	tests := []struct {
		name              string
		target            string
		pattern           string
		expectedCounts    int
		expectedPositions []int // 1-based positions where pattern ends
	}{
		{
			name:              "Single occurrence",
			target:            "hello world",
			pattern:           "world",
			expectedCounts:    1,
			expectedPositions: []int{11},
		},
		{
			name:              "Multiple occurrences",
			target:            "abababa",
			pattern:           "aba",
			expectedCounts:    3,
			expectedPositions: []int{3, 5, 7},
		},
		{
			name:              "No occurrences",
			target:            "abcdef",
			pattern:           "xyz",
			expectedCounts:    0,
			expectedPositions: []int(nil),
		},
		{
			name:              "Pattern equals target",
			target:            "exact",
			pattern:           "exact",
			expectedCounts:    1,
			expectedPositions: []int{5},
		},
		{
			name:              "Overlapping patterns",
			target:            "aaaaa",
			pattern:           "aaa",
			expectedCounts:    3,
			expectedPositions: []int{3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preLen := getPreLen(tt.pattern)
			matchLengths := matchLen(tt.target, tt.pattern, preLen)

			// Find all positions where the match length equals pattern length
			var foundPositions []int
			patternLen := len(tt.pattern)
			for i, ml := range matchLengths {
				if ml == patternLen {
					// Convert to 1-based position where pattern ends
					foundPositions = append(foundPositions, i+1)
				}
			}

			assert.Equal(t, tt.expectedCounts, len(foundPositions), "Expected %d matches, got %d", tt.expectedCounts, len(foundPositions))
			assert.Equal(t, tt.expectedPositions, foundPositions, "Expected positions %v, got %v", tt.expectedPositions, foundPositions)
		})
	}
}
