package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{
			name: "empty strings",
			a:    "",
			b:    "",
			want: 0,
		},
		{
			name: "first string empty",
			a:    "",
			b:    "abc",
			want: 0,
		},
		{
			name: "second string empty",
			a:    "abc",
			b:    "",
			want: 0,
		},
		{
			name: "identical strings",
			a:    "abcde",
			b:    "abcde",
			want: 5,
		},
		{
			name: "no common subsequence",
			a:    "abc",
			b:    "def",
			want: 0,
		},
		{
			name: "classic example 1",
			a:    "abcde",
			b:    "ace",
			want: 3,
		},
		{
			name: "classic example 2",
			a:    "abc",
			b:    "abc",
			want: 3,
		},
		{
			name: "leetcode example 1",
			a:    "abcde",
			b:    "ace",
			want: 3,
		},
		{
			name: "leetcode example 2",
			a:    "abc",
			b:    "abc",
			want: 3,
		},
		{
			name: "leetcode example 3",
			a:    "abc",
			b:    "def",
			want: 0,
		},
		{
			name: "partial match",
			a:    "oxcpqrsvwf",
			b:    "shmtulqrypy",
			want: 2,
		},
		{
			name: "longer strings",
			a:    "abcdefghij",
			b:    "acfgi",
			want: 5,
		},
		{
			name: "repeated characters",
			a:    "aaa",
			b:    "aaa",
			want: 3,
		},
		{
			name: "different lengths",
			a:    "abcd",
			b:    "acd",
			want: 3,
		},
		{
			name: "complex pattern",
			a:    "pmjghexybyrgzczy",
			b:    "hafcdqbgncrcbihkd",
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonSubsequence(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "longestCommonSubsequence(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.want)
		})
	}
}
