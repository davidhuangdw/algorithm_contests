package math_geo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{
			name:     "leetcode example 1",
			a:        "2",
			b:        "3",
			expected: "6",
		},
		{
			name:     "leetcode example 2",
			a:        "123",
			b:        "456",
			expected: "56088",
		},
		{
			name:     "multiply by zero",
			a:        "123",
			b:        "0",
			expected: "0",
		},
		{
			name:     "zero multiplied by number",
			a:        "0",
			b:        "456",
			expected: "0",
		},
		{
			name:     "single digit multiplication",
			a:        "9",
			b:        "9",
			expected: "81",
		},
		{
			name:     "large numbers",
			a:        "999",
			b:        "999",
			expected: "998001",
		},
		{
			name:     "different length numbers",
			a:        "12",
			b:        "345",
			expected: "4140",
		},
		{
			name:     "carry handling",
			a:        "99",
			b:        "99",
			expected: "9801",
		},
		{
			name:     "one digit with multi-digit",
			a:        "5",
			b:        "1234",
			expected: "6170",
		},
		{
			name:     "large multiplication",
			a:        "123456789",
			b:        "987654321",
			expected: "121932631112635269",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := multiply(tt.a, tt.b)
			assert.Equal(t, tt.expected, result, "multiply(%q, %q) = %q, want %q", tt.a, tt.b, result, tt.expected)
		})
	}
}
