package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{
			name:     "LeetCode example 1",
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			name:     "LeetCode example 2",
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
		{
			name:     "both zeros",
			a:        "0",
			b:        "0",
			expected: "0",
		},
		{
			name:     "zero plus one",
			a:        "0",
			b:        "1",
			expected: "1",
		},
		{
			name:     "one plus zero",
			a:        "1",
			b:        "0",
			expected: "1",
		},
		{
			name:     "one plus one",
			a:        "1",
			b:        "1",
			expected: "10",
		},
		{
			name:     "carry propagation",
			a:        "111",
			b:        "1",
			expected: "1000",
		},
		{
			name:     "different lengths 1",
			a:        "101",
			b:        "11",
			expected: "1000",
		},
		{
			name:     "different lengths 2",
			a:        "1",
			b:        "111",
			expected: "1000",
		},
		{
			name:     "large numbers 1",
			a:        "111111",
			b:        "1",
			expected: "1000000",
		},
		{
			name:     "large numbers 2",
			a:        "111111",
			b:        "111111",
			expected: "1111110",
		},
		{
			name:     "all ones",
			a:        "1111",
			b:        "1111",
			expected: "11110",
		},
		{
			name:     "with multiple carries",
			a:        "10101",
			b:        "11011",
			expected: "110000",
		},
		{
			name:     "one empty string",
			a:        "101",
			b:        "",
			expected: "101",
		},
		{
			name:     "another empty string",
			a:        "",
			b:        "110",
			expected: "110",
		},
		{
			name:     "very large numbers",
			a:        "11111111111111111111",
			b:        "1",
			expected: "100000000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addBinary(tt.a, tt.b)
			assert.Equal(t, tt.expected, result, "addBinary(%q, %q)", tt.a, tt.b)
		})
	}
}
