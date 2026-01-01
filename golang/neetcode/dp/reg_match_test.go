package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s        string
		p        string
		expected bool
	}{
		// LeetCode examples
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},

		// Edge cases
		{"", "", true},
		{"", "a*", true},
		{"a", "", false},
		{"", ".*", true},

		// Simple cases
		{"a", "a", true},
		{"a", ".", true},
		{"ab", "a.", true},
		{"abc", "a.c", true},

		// Complex cases with *
		{"aaa", "a*a", true},
		{"aaa", "ab*a*c*a", true},
		{"a", "ab*", true},
		{"bbbba", ".*a*a", true},
	}

	for _, tt := range tests {
		t.Run(tt.s+"_"+tt.p, func(t *testing.T) {
			result := isMatch(tt.s, tt.p)
			assert.Equal(t, tt.expected, result)
		})
	}
}
