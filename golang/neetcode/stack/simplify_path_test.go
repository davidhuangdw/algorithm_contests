package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "LeetCode Example 1",
			path:     "/home/",
			expected: "/home",
		},
		{
			name:     "LeetCode Example 2",
			path:     "/../",
			expected: "/",
		},
		{
			name:     "LeetCode Example 3",
			path:     "/home//foo/",
			expected: "/home/foo",
		},
		{
			name:     "Complex path with parent directory",
			path:     "/a/./b/../../c/",
			expected: "/c",
		},
		{
			name:     "Multiple parent directories",
			path:     "/a/../../b/../c//.//",
			expected: "/c",
		},
		{
			name:     "Root path only",
			path:     "/",
			expected: "/",
		},
		{
			name:     "Empty path",
			path:     "",
			expected: "/",
		},
		{
			name:     "Multiple current directories",
			path:     "/a/././b/./c/",
			expected: "/a/b/c",
		},
		{
			name:     "Parent directory at root",
			path:     "/../../..",
			expected: "/",
		},
		{
			name:     "Mixed operations",
			path:     "/a/b/c/.././d/../../e/../f/",
			expected: "/a/f",
		},
		{
			name:     "Complex nested structure",
			path:     "/a/../b/./c/d/../../e/f/../g/",
			expected: "/b/e/g",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := simplifyPath(tt.path)
			assert.Equal(t, tt.expected, result)
		})
	}
}