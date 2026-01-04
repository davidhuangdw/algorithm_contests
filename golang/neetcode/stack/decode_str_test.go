package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "LeetCode Example 1",
			input:    "3[a]2[bc]",
			expected: "aaabcbc",
		},
		{
			name:     "LeetCode Example 2",
			input:    "3[a2[c]]",
			expected: "accaccacc",
		},
		{
			name:     "LeetCode Example 3",
			input:    "2[abc]3[cd]ef",
			expected: "abcabccdcdcdef",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Nested encoding",
			input:    "2[3[a]b]",
			expected: "aaabaaab",
		},
		{
			name:     "Multiple digits",
			input:    "10[a]",
			expected: "aaaaaaaaaa",
		},
		{
			name:     "Complex nested",
			input:    "2[abc3[de]f]",
			expected: "abcdededefabcdededef",
		},
		{
			name:     "No encoding",
			input:    "abcdef",
			expected: "abcdef",
		},
		{
			name:     "Single bracket",
			input:    "3[]",
			expected: "",
		},
		{
			name:     "Deep nesting",
			input:    "2[3[4[a]]]",
			expected: "aaaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := decodeString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
