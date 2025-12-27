package backtrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected [][]string
	}{
		{
			name:     "empty string",
			s:        "",
			expected: [][]string{{}},
		},
		{
			name:     "single character",
			s:        "a",
			expected: [][]string{{"a"}},
		},
		{
			name: "two characters - palindrome",
			s:    "aa",
			expected: [][]string{
				{"a", "a"},
				{"aa"},
			},
		},
		{
			name: "two characters - not palindrome",
			s:    "ab",
			expected: [][]string{
				{"a", "b"},
			},
		},
		{
			name: "three characters - all same",
			s:    "aaa",
			expected: [][]string{
				{"a", "a", "a"},
				{"a", "aa"},
				{"aa", "a"},
				{"aaa"},
			},
		},
		{
			name: "classic LeetCode example",
			s:    "aab",
			expected: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name: "longer palindrome",
			s:    "aabb",
			expected: [][]string{
				{"a", "a", "b", "b"},
				{"a", "a", "bb"},
				{"aa", "b", "b"},
				{"aa", "bb"},
			},
		},
		{
			name: "no palindromes except single chars",
			s:    "abc",
			expected: [][]string{
				{"a", "b", "c"},
			},
		},
		{
			name: "complex palindrome patterns",
			s:    "abba",
			expected: [][]string{
				{"a", "b", "b", "a"},
				{"a", "bb", "a"},
				{"abba"},
			},
		},
		{
			name: "mixed palindrome lengths",
			s:    "noon",
			expected: [][]string{
				{"n", "o", "o", "n"},
				{"n", "oo", "n"},
				{"noon"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := partition(tt.s)

			assert.Equal(t, tt.expected, result,
				"partition(%s) should equal %v", tt.s, tt.expected)

		})
	}
}
