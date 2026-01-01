package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected int
	}{
		{
			name:     "leetcode example 1",
			s:        "rabbbit",
			t:        "rabbit",
			expected: 3,
		},
		{
			name:     "leetcode example 2",
			s:        "babgbag",
			t:        "bag",
			expected: 5,
		},
		{
			name:     "empty t",
			s:        "abc",
			t:        "",
			expected: 1,
		},
		{
			name:     "empty s",
			s:        "",
			t:        "abc",
			expected: 0,
		},
		{
			name:     "both empty",
			s:        "",
			t:        "",
			expected: 1,
		},
		{
			name:     "t longer than s",
			s:        "ab",
			t:        "abc",
			expected: 0,
		},
		{
			name:     "exact match",
			s:        "abc",
			t:        "abc",
			expected: 1,
		},
		{
			name:     "no match",
			s:        "abc",
			t:        "def",
			expected: 0,
		},
		{
			name:     "multiple occurrences",
			s:        "aaaa",
			t:        "aa",
			expected: 6,
		},
		{
			name:     "complex pattern",
			s:        "aabbcc",
			t:        "abc",
			expected: 8,
		},
		{
			name:     "single character t",
			s:        "abcabc",
			t:        "a",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numDistinct(tt.s, tt.t)
			assert.Equal(t, tt.expected, result, "numDistinct(%q, %q) = %d, want %d", tt.s, tt.t, result, tt.expected)
		})
	}
}