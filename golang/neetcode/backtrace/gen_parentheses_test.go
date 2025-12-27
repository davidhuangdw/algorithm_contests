package backtrace

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []string
	}{
		{
			name:     "n=0 - empty result",
			n:        0,
			expected: []string{""},
		},
		{
			name:     "n=1 - single pair",
			n:        1,
			expected: []string{"()"},
		},
		{
			name:     "n=2 - two pairs",
			n:        2,
			expected: []string{"(())", "()()"},
		},
		{
			name: "n=3 - three pairs",
			n:    3,
			expected: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
		{
			name: "n=4 - four pairs",
			n:    4,
			expected: []string{
				"(((())))",
				"((()()))",
				"((())())",
				"((()))()",
				"(()(()))",
				"(()()())",
				"(()())()",
				"(())(())",
				"(())()()",
				"()((()))",
				"()(()())",
				"()(())()",
				"()()(())",
				"()()()()",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateParenthesis(tt.n)

			// Sort both result and expected for consistent comparison
			sort.Strings(result)
			sort.Strings(tt.expected)

			assert.Equal(t, tt.expected, result,
				"generateParenthesis(%d) should equal %v", tt.n, tt.expected)

			// Additional validation: all generated strings should be valid parentheses
			for _, s := range result {
				assert.True(t, isValidParentheses(s),
					"Generated string %s should be valid parentheses", s)
			}
		})
	}
}

// Helper function to validate parentheses string
func isValidParentheses(s string) bool {
	balance := 0
	for _, ch := range s {
		if ch == '(' {
			balance++
		} else if ch == ')' {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}
