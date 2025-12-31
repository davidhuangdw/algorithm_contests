package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValidString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "leetcode example 1",
			s:        "()",
			expected: true,
		},
		{
			name:     "leetcode example 2",
			s:        "(*)",
			expected: true,
		},
		{
			name:     "leetcode example 3",
			s:        "(*))",
			expected: true,
		},
		{
			name:     "all wildcards",
			s:        "****",
			expected: true,
		},
		{
			name:     "empty string",
			s:        "",
			expected: true,
		},
		{
			name:     "unbalanced right",
			s:        ")(",
			expected: false,
		},
		{
			name:     "too many right",
			s:        "())",
			expected: false,
		},
		{
			name:     "too many left",
			s:        "(()",
			expected: false,
		},
		{
			name:     "complex valid with wildcards",
			s:        "(*())",
			expected: true,
		},
		{
			name:     "complex invalid",
			s:        "((*)",
			expected: true,
		},
		{
			name:     "leetcode edge case",
			s:        "((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkValidString(tt.s)
			assert.Equal(t, tt.expected, result, "checkValidString(%q) = %v, want %v", tt.s, result, tt.expected)
		})
	}
}
