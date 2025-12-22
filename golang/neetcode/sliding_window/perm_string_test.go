package sliding_window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		// Basic cases
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"abc", "bbbca", true},
		{"hello", "ooolleoooleh", false},
		{"adc", "dcda", true},

		// Edge cases
		{"", "", true},
		{"", "abc", true},
		{"a", "", false},
		{"a", "a", true},
		{"a", "b", false},
		{"a", "aa", true},
		{"ab", "a", false},
		{"ab", "ba", true},
		{"abc", "cba", true},
		{"abc", "bca", true},

		// Complex cases
		{"prosperity", "properties", false},
		{"abc", "ccccbbbbaaaa", false},
		{"abc", "abacaba", true},
		{"abcd", "dabc", true},
		{"abcd", "dcba", true},
		{"abcde", "edcba", true},
		{"abc", "abcbac", true},
		{"xyz", "abcdefghijklmnopqrstuvwxyz", true},
		{"xyz", "abcdefghijklmnopqrstuvwxy", false},

		// Same length strings
		{"abc", "abc", true},
		{"abc", "acb", true},
		{"abc", "bac", true},
		{"abc", "bca", true},
		{"abc", "cab", true},
		{"abc", "cba", true},
		{"abc", "abd", false},
		{"abc", "abcd", true},

		// Repeated characters
		{"aab", "baa", true},
		{"aab", "aba", true},
		{"aab", "abab", true},
		{"aab", "abac", true},
		{"aab", "abca", false},
		{"aaa", "aaa", true},
		{"aaa", "aa", false},
		{"aaa", "aaaa", true},
		{"aaa", "aaba", false},

		// Long strings
		{"abcdefghij", "hijabcdefg", true},
		{"abcdefghij", "abcdefghijk", true},
		{"abcdefghij", "abcdeffghij", false},
		{"abcdefghij", "abcdffghij", false},
	}

	for _, tt := range tests {
		t.Run(tt.s1+"_in_"+tt.s2, func(t *testing.T) {
			result := checkInclusion(tt.s1, tt.s2)
			assert.Equal(t, tt.want, result,
				"checkInclusion(%q, %q) = %v, want %v", tt.s1, tt.s2, result, tt.want)
		})
	}
}
