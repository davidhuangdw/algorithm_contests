package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		s3       string
		expected bool
	}{
		{
			name:     "leetcode example 1",
			s1:       "aabcc",
			s2:       "dbbca",
			s3:       "aadbbcbcac",
			expected: true,
		},
		{
			name:     "leetcode example 2",
			s1:       "aabcc",
			s2:       "dbbca",
			s3:       "aadbbbaccc",
			expected: false,
		},
		{
			name:     "empty strings",
			s1:       "",
			s2:       "",
			s3:       "",
			expected: true,
		},
		{
			name:     "s1 empty",
			s1:       "",
			s2:       "abc",
			s3:       "abc",
			expected: true,
		},
		{
			name:     "s2 empty",
			s1:       "abc",
			s2:       "",
			s3:       "abc",
			expected: true,
		},
		{
			name:     "different length",
			s1:       "abc",
			s2:       "def",
			s3:       "abcdefg",
			expected: false,
		},
		{
			name:     "simple interleave",
			s1:       "ab",
			s2:       "cd",
			s3:       "acbd",
			expected: true,
		},
		{
			name:     "invalid interleave",
			s1:       "ab",
			s2:       "cd",
			s3:       "abcdd",
			expected: false,
		},
		{
			name:     "same characters",
			s1:       "aaa",
			s2:       "aaa",
			s3:       "aaaaaa",
			expected: true,
		},
		{
			name:     "complex valid",
			s1:       "abc",
			s2:       "def",
			s3:       "adbecf",
			expected: true,
		},
		{
			name:     "complex invalid",
			s1:       "abc",
			s2:       "def",
			s3:       "abdecf",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isInterleave(tt.s1, tt.s2, tt.s3)
			assert.Equal(t, tt.expected, result, "isInterleave(%q, %q, %q) = %v, want %v", tt.s1, tt.s2, tt.s3, result, tt.expected)
		})
	}
}
